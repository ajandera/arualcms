<?php
declare(strict_types = 1);

namespace ArualCms\Lib;

/**
 * Class Router
 * @package ArualCms\Lib
 */
class Router
{
    /**
     * @param $route
     * @param $callback
     */
    public static function get(string $route, callable $callback)
    {
        self::isOption();
        if (strcasecmp($_SERVER['REQUEST_METHOD'], 'GET') !== 0) {
            return;
        }

        self::on($route, $callback);
    }

    /**
     * @param $route
     * @param $callback
     */
    public static function post(string $route, callable $callback)
    {
        self::isOption();
        if (strcasecmp($_SERVER['REQUEST_METHOD'], 'POST') !== 0) {
            return;
        }

        self::on($route, $callback);
    }

    /**
     * @param $route
     * @param $callback
     */
    public static function put(string $route, callable $callback)
    {
        self::isOption();
        if (strcasecmp($_SERVER['REQUEST_METHOD'], 'PUT') !== 0) {
            return;
        }

        self::on($route, $callback);
    }

    /**
     * @param $route
     * @param $callback
     */
    public static function delete(string $route, callable $callback)
    {
        self::isOption();
        if (strcasecmp($_SERVER['REQUEST_METHOD'], 'DELETE') !== 0) {
            return;
        }

        self::on($route, $callback);
    }

    /**
     * @param $regex
     * @param $cb
     */
    public static function on(string $regex, callable $cb): void
    {
        $params = $_SERVER['REQUEST_URI'];
        $params = (stripos($params, "/") !== 0) ? "/" . $params : $params;
        $regex = str_replace('/', '\/', $regex);
        $is_match = preg_match('/^' . ($regex) . '$/', $params, $matches, PREG_OFFSET_CAPTURE);

        if ($is_match) {
            // first value is normally the route, lets remove it
            array_shift($matches);
            // Get the matches as parameters
            $params = array_map(function ($param) {
                return $param[0];
            }, $matches);
            $cb(new Request($params), new Response());
        }
    }

    /**
     * Handle option pre-request.
     */
    public static function isOption(): void
    {
        if ($_SERVER['REQUEST_METHOD'] === 'OPTIONS') {
            header('Access-Control-Allow-Origin: *');
            header('Access-Control-Allow-Methods: POST, GET, DELETE, PUT, PATCH, OPTIONS');
            header('Access-Control-Allow-Headers: token, Content-Type, Authorization');
            header('Access-Control-Max-Age: 1728000');
            header('Content-Length: 0');
            header('Content-Type: text/plain');
            die();
        }
    }
}
