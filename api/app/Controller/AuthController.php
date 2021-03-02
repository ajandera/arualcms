<?php

namespace ArualCms\Controller;

use ArualCms\Lib\Response;

class AuthController
{
    public function __construct()
    {

    }

    public function auth($req, Response $res)
    {
        $data['username'] = $req->username;
        $data['userId'] = 1;
        $data['success'] = true;
        $res->toJSON($data);
    }
}
