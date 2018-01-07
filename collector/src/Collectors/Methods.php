<?php

namespace GoBitrix\Collectors;

class Methods {

    protected function buildFileList() {
        $provider = new \CRestProvider();
        $description = $provider->getDescription();

        $files = [];

        foreach($description as $module => $commands) {
            if($module == \CRestUtil::GLOBAL_SCOPE) {
                $module = "";
            }

            foreach($commands as $commandName => $commandMethod) {
                if($commandName == \CRestUtil::EVENTS) {
                    continue;
                }

                $fullCommand = explode(".", $commandName);
                $command = $fullCommand;
                if($command[0] == $module) {
                    array_shift($command);
                }
                if($module && count($command) > 1) {
                    $filename = $module . "_" . array_shift($command);
                } elseif(count($command) > 1) {
                    $filename = array_shift($command);
                } elseif($module) {
                    $filename = $module;
                } else {
                    $filename = "global";
                }

                $files[$filename][] = $fullCommand;
            }
        }

        return $files;
    }

    protected function writeFunctions($files, $path) {

        $templateFunc = file_get_contents("templates/command_func.txt");
        $templateFile = file_get_contents("templates/command_file.txt");

        foreach($files as $filename => $commands) {
            $funcString = "";

            foreach($commands as $command) {

                $commandName = "";
                foreach($command as $word) {
                    $commandName .= ucfirst($word);
                }

                $funcString .= str_replace([
                    '#COMMAND#'
                ], [
                    $commandName
                ], $templateFunc);
            }

            $fileString = str_replace([
                '#FUNCTIONS#'
            ], [
                $funcString
            ], $templateFile);

            file_put_contents($path . "/" . $filename . ".go", $fileString);
        }
    }

    public function collect($path) {
        $this->writeFunctions($this->buildFileList(), $path);
    }
}