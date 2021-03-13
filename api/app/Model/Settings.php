<?php
declare(strict_types = 1);

namespace ArualCms\Model;

use ArualCms\Lib\Config;

/**
 * Class Settings
 * @package ArualCms\Model
 */
class Settings
{
    /** @var array  */
    private static $DATA = [];

    /**
     * @return array
     */
    public static function all(): array
    {
        return self::$DATA;
    }

    /**
     * @param $settings
     */
    public static function update($settings): void
    {
        self::$DATA = $settings;
        self::save();
    }

    /**
     * @param string $key
     * @return array|mixed
     */
    public static function findByKey(string $key)
    {
        foreach (self::$DATA as $setting) {
            if ($setting->key === $key) {
                return $setting;
            }
        }
        return [];
    }

    /**
     * Load setting from storage
     */
    public static function load(): void
    {
        $DB_PATH = Config::get('DB_PATH', __DIR__ . '/../../database/');
        self::$DATA = json_decode(file_get_contents($DB_PATH . 'settings.json'));
    }

    /**
     * Save setting to storage
     */
    public static function save(): void
    {
        $DB_PATH = Config::get('DB_PATH', __DIR__ . '/../../database/');
        file_put_contents($DB_PATH . 'settings.json', json_encode(self::$DATA, JSON_PRETTY_PRINT));
    }
}
