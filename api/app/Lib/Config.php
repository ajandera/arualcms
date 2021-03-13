<?php
declare(strict_types = 1);

namespace ArualCms\Lib;

/**
 * Class Config
 * @package ArualCms\Lib
 */
class Config
{
    /** @var array */
    private static $config;

    /**
     * @param string $key
     * @param string|null $default
     * @return mixed|null
     */
    public static function get(string $key, string $default = null)
    {
        if (is_null(self::$config)) {
            self::$config = require_once(__DIR__.'/../../config.php');
        }

        return !empty(self::$config[$key])?self::$config[$key]:$default;
    }
}
