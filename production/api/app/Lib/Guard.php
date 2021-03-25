<?php
declare(strict_types = 1);

namespace ArualCms\Lib;

use Firebase\JWT\JWT;

/**
 * Class Guard
 * @package ArualCms\Lib
 */
class Guard
{
    /**
     * @param string $jwt
     * @param Response $response
     */
    public static function check(string $jwt, Response $response): void
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
