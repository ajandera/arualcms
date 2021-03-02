<?php

namespace ArualCms\Model;

use ArualCms\Lib\Config;

class Settings
{
    private static $DATA = [];

    public static function all()
    {
        return self::$DATA;
    }

    public static function update($settings)
    {
        self::$DATA = $settings;
        self::save();
    }

    public static function findByKey(string $key)
    {
        foreach (self::$DATA as $setting) {
            if ($setting->key === $key) {
                return $setting;
            }
        }
        return [];
    }

    public static function load()
    {
        $DB_PATH = Config::get('DB_PATH', __DIR__ . '/../database/');
        self::$DATA = json_decode(file_get_contents($DB_PATH . 'settings.json'));
    }

    public static function save()
    {
        $DB_PATH = Config::get('DB_PATH', __DIR__ . '/../database/');
        file_put_contents($DB_PATH . 'settings.json', json_encode(self::$DATA, JSON_PRETTY_PRINT));
    }
}
