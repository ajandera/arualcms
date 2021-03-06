<?php

namespace ArualCms\Lib;

class Response
{
    private $status = 200;

    public function status(int $code)
    {
        $this->status = $code;
        return $this;
    }

    public function toJSON($data = [])
    {
        http_response_code($this->status);
        header('Access-Control-Allow-Origin: *');
        header('Access-Control-Allow-Headers: Content-Type, X-Requested-With, Authorization');
        header('Access-Control-Max-Age: 1000');
        header('Content-Type: application/json');
        echo json_encode($data);
        exit();
    }
}
