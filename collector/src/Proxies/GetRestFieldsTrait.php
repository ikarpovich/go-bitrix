<?php

namespace GoBitrix\Proxies;

trait GetRestFieldsTrait {
    public function getRestFields() {
        if(method_exists($this, 'getFieldsInfo')) {
            return $this->getFieldsInfo();
        }
        elseif(method_exists($this, 'getFields')) {
            return $this->getFields();
        }
    }
}