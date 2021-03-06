<?php

namespace ArualCms\Lib;

use Firebase\JWT\JWT;

class Guard
{
    public static function check(string $jwt, Response $response)
    {
        // if jwt is not empty
        if ($jwt) {
            // if decode succeed, show user details
            try {
                // decode jwt
                $decoded = JWT::decode($jwt, Config::get("KEY", "fsdfsdfsdSDF34"), ['HS256']);
                if ($decoded->exp < time()) {
                    $response->status(401)->toJSON(['success' => false, 'message' => 'Token expired']);
                }
            } catch (\Exception $e) {
                $response->status(401)->toJSON(['success' => false, 'message' => $e->getMessage()]);
            }
        }
    }
}
