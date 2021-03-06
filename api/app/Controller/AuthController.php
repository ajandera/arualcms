<?php

namespace ArualCms\Controller;

use ArualCms\Lib\Config;
use ArualCms\Lib\Response;
use ArualCms\Model\Users;
use Firebase\JWT\JWT;

class AuthController
{
    public function __construct()
    {
        Users::load();
    }

    public function auth($req, Response $res)
    {
        $user = Users::findByUsername($req->username);
        if ($user && password_verify($req->password, $user->password)) {
            $token = [
                "iat" => time(),
                "exp" => Config::get("EXPIRATION_TIME", time()),
                "iss" => Config::get("ISSUER", "http://localhost:8080"),
                "data" => [
                    "id" => $user->id,
                    "username" => $user->username
                ]
            ];
            $data["id"] = $user->id;
            $data["username"] = $user->username;
            $data['jwt'] = $jwt = JWT::encode($token, Config::get("KEY", "fsdfsdfsdSDF34"));
            $data['success'] = true;
        } else {
            $data['success'] = false;
            $data['message'] = "Bad credential.";
        }

        $res->toJSON($data);
    }
}
