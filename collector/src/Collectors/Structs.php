<?php

namespace GoBitrix\Collectors;

use Bitrix\Rest\RestException;

class Structs {

    protected const dataTypeMap = [
        'integer' => 'int',
        'char' => 'string',
        'date' => 'time.Time',
        'double' => 'float64',
        'datetime' => 'time.Time',
        'file' => 'int'
    ];

    protected const dataTypeImports = [
        'time.Time' => 'time'
    ];

    public function collect(string $path) {
        $provider = new \CRestProvider();
        $description = $provider->getDescription();

        foreach($description as $module => $commands) {

            foreach ($commands as $commandName => $commandMethod) {
                if ($commandName == \CRestUtil::EVENTS) {
                    continue;
                }

                $command = explode(".", $commandName);
                if ($command[0] == $module) {
                    array_shift($command);
                }

                $entity = $this->getFields($module, $command);

                /*
                 * @todo handle types: ['location', 'file', 'currency_localization', 'product_file', 'object']
                 */

                $this->writeEntity($path, $entity);
            }
        }
    }

    protected function getFields($module, $command) {

        $proxyClassName = "\\GoBitrix\\Proxies\\" . ucfirst($module);

        foreach($command as $word) {
            $proxyClassName .= "\\" . ucfirst($word);
        }

        if (!class_exists($proxyClassName)) {
            if(count($command) > 1) {
                array_pop($command);
                return $this->getFields($module, $command);
            } else {
                return;
            }
        }

        try {
            $proxy = new $proxyClassName();

            return [
                'module'    => $module,
                'command'   => $command,
                'fields'    => $proxy->getRestFields(),
            ];
        } catch(RestException $e) {
            return;
        }
    }

    protected function writeEntity($path, $entity) {

        if(!$entity['module'] || !$entity['command']) {
            return;
        }

        @mkdir($path . '/'. $entity['module']);

        $templateField = file_get_contents("templates/struct_field.txt");
        $templateFile = file_get_contents("templates/struct_file.txt");
        $templateImportLine = file_get_contents("templates/struct_import_line.txt");
        $templateImportFile = file_get_contents("templates/struct_import_file.txt");

        $structName = "";
        foreach($entity['command'] as $word) {
            $structName .= ucfirst(strtolower($word));
        }

        $fieldString = "";

        /**
         * @todo handle attributes
         */

        // Handle fields

        $imports = [];
        foreach($entity['fields'] as $fieldName => $fieldParams) {

            $fieldNameArray = explode("_", $fieldName);
            $fieldNameGo = "";
            foreach($fieldNameArray as &$fieldNameElement) {
                $fieldNameGo .= ucfirst(strtolower($fieldNameElement));
            }

            $dataType = $this->convertDataType($fieldParams);
            if(in_array($dataType, array_keys(self::dataTypeImports))) {
                $imports[self::dataTypeImports[$dataType]] = $dataType;
            }

            $fieldString .= str_replace([
                '#GO_NAME#',
                '#BITRIX_NAME#',
                '#GO_TYPE#',
            ], [
                $fieldNameGo,
                $fieldName,
                $dataType,
            ], $templateField);
        }

        if(!$fieldString) {
            return;
        }

        // Handle imports

        $importsFileString = "";
        if(count($imports)) {
            $importsString = "";
            foreach ($imports as $import => $dataType) {
                $importsString .= str_replace([
                    '#IMPORT#',
                ], [
                    $import,
                ], $templateImportLine);
            }

            $importsFileString = str_replace([
                '#IMPORTS#',
            ], [
                $importsString,
            ], $templateImportFile);
        }

        // Write struct

        $filename = $path . '/' . $entity['module'] . '/' . implode('_', $entity['command']) . '.go';

        $fileString = str_replace([
            '#MODULE#',
            '#STRUCT_NAME#',
            '#FIELDS#',
            '#IMPORTS#',
        ], [
            $entity['module'],
            $structName,
            $fieldString,
            $importsFileString,
        ], $templateFile);

        file_put_contents($filename, $fileString);
    }

    protected function convertDataType($fieldParams) {
        /**
         * @todo handle complex types like crm_company
         */

        if(isset($fieldParams["TYPE"])) {
            $bitrixDataType = $fieldParams["TYPE"];
        } elseif(isset($fieldParams["type"])) {
            $bitrixDataType = $fieldParams["type"];
        } else {
            return;
        }

        if(in_array($bitrixDataType, array_keys(self::dataTypeMap))) {
            return self::dataTypeMap[$bitrixDataType];
        }

        return "string";
    }
}