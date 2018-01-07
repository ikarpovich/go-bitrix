<?php

namespace GoBitrix\Commands;

use Notamedia\ConsoleJedi\Application\Command\BitrixCommand;
use Bitrix\Main\Loader;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Output\OutputInterface;
use Symfony\Component\Console\Input\InputArgument;
use GoBitrix\Collectors;

class CollectCommand extends BitrixCommand
{
    protected function configure()
    {
        $this
            ->setName('gobitrix:collect')
            ->setDescription('Collect REST methods and entities into Go code')
            ->addArgument(
                'methodsPath',
                InputArgument::OPTIONAL,
                'Output path to put generated method files',
                '../build/methods'
            )
            ->addArgument(
                'structsPath',
                InputArgument::OPTIONAL,
                'Output path to put generated struct files',
                '../build/structs'
            );
    }

    protected function initialize(InputInterface $input, OutputInterface $output)
    {
        Loader::includeModule('rest');
    }

    protected function execute(InputInterface $input, OutputInterface $output)
    {
        //$methodsCollector = new Collectors\Methods();
        //$methodsCollector->collect($input->getArgument("methodsPath"));

        $structsCollector = new Collectors\Structs();
        $structsCollector->collect($input->getArgument("structsPath"));
    }
}