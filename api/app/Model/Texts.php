<?php

namespace ArualCms\Model;

use ArualCms\Lib\Config;

class Texts
{
    private static $DATA = [];

    public static function all()
    {
        return self::$DATA;
    }

    public static function update($texts)
    {
        self::$DATA = $texts;
        self::save();
    }

    public static function findByKey(string $key)
    {
        foreach (self::$DATA as $text) {
            if ($text->key === $key) {
                return $text;
            }
        }
        return [];
    }

    public static function load()
    {
        $DB_PATH = Config::get('DB_PATH', __DIR__ . '/../database/');
        self::$DATA = json_decode(file_get_contents($DB_PATH . 'texts.json'));
    }

    public static function save()
    {
        $DB_PATH = Config::get('DB_PATH', __DIR__ . '/../database/');
        file_put_contents($DB_PATH . 'texts.json', json_encode(self::$DATA, JSON_PRETTY_PRINT));
    }
}
