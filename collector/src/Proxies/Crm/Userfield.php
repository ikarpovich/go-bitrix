<?php

namespace GoBitrix\Proxies\Crm;

use GoBitrix\Proxies;

class Userfield extends \CCrmUserFieldRestProxy {
    use Proxies\GetRestFieldsTrait;

    public function __construct() {
        parent::__construct(null);
    }
}