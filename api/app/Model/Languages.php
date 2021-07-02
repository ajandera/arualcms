<?php
declare(strict_types = 1);

namespace ArualCms\Model;

use ArualCms\Lib\Config;

/**
 * Class Languages
 * @package ArualCms\Model
 */
class Languages
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
     * @param $texts
     */
    public static function update($texts): void
    {
        self::$DATA = $texts;
        self::save();
    }

    /**
     * @param string $key
     * @return array|mixed
     */
    public static function findByKey(string $key): array
    {
        foreach (self::$DATA as $text) {
            if ($text->key === $key) {
                return $text;
            }
        }
        return [];
    }

    /**
     * Load texts from storage
     */
    public static function load(): void
    {
        $DB_PATH = Config::get('DB_PATH', __DIR__ . '/../../database/');
        self::$DATA = json_decode(file_get_contents($DB_PATH . 'languages.json'));
    }

    /**
     * Save data to storage
     */
    public static function save(): void
    {
        $DB_PATH = Config::get('DB_PATH', __DIR__ . '/../../database/');
        file_put_contents($DB_PATH . 'languages.json', json_encode(self::$DATA, JSON_PRETTY_PRINT));
    }
}
