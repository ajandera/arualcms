<?php
declare(strict_types = 1);

namespace ArualCms\Controller;

use ArualCms\Lib\Config;
use ArualCms\Lib\Response;
use ArualCms\Model\Users;
use Firebase\JWT\JWT;

/**
 * Class AuthController
 * @package ArualCms\Controller
 */
class AuthController
{
    use Users;

    /**
     * @param $req
     * @param Response $res
     * @throws \Exception
     */
    public function auth($req, Response $res): void
    {
        $user = $this->findByUsername($req->username)[0];
        if ($user && password_verify($req->password, $user['password'])) {
            $token = [
                "iat" => time(),
                "exp" => Config::get("EXPIRATION_TIME"),
                "iss" => Config::get("ISSUER", "http://localhost:8080"),
                "data" => [
                    "id" => $user['id'],
                    "username" => $user['username']
                ]
            ];
            $data["id"] = $user['id'];
            $data["username"] = $user['username'];
            $data['jwt'] = JWT::encode($token, Config::get("KEY", "fsdfsdfsdSDF34"));
            $data['success'] = true;
        } else {
            $data['success'] = false;
            $data['message'] = "Bad credentials.";
        }

        $res->toJSON($data);
    }

    /**
     * @param $req
     * @param Response $res
     * @throws \MongoCursorException
     * @throws \Exception
     */
    public function recovery($req, Response $res): void {
        $user = $this->findByUsername($req->username);
        $user['password'] = password_hash($user['password'], PASSWORD_BCRYPT);
        $this->edit($user);
        $data['success'] = false;
        $data['message'] = "New password was set.";
        $res->toJSON($data);
    }
}
