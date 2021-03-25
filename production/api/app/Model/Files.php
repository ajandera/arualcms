<?php
declare(strict_types = 1);

namespace ArualCms\Model;

use ArualCms\Lib\Config;

/**
 * Class Files
 * @package ArualCms\Model
 */
class Files
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
     * @param $file
     * @return array
     */
    public static function add($file): array
    {
        $file['id'] = self::getNextId();
        self::$DATA[] = $file;
        self::save();

        $STORAGE_FRONT = Config::get('STORAGE_FRONT', '');
        return ['link' => 'http://' . $_SERVER["HTTP_HOST"] . $STORAGE_FRONT . '/storage/' . $file['name'], 'name' => $file['name']];
    }

    /**
     * @return int|mixed
     */
    private static function getNextId(): int
    {
        $ids = [];
        foreach (self::$DATA as $data) {
            $ids[] = $data->id;
        }
        return max($ids) + 1;
    }

    /**
     * @param int $id
     * @return array|mixed
     */
    public static function findById(int $id): array
    {
        foreach (self::$DATA as $post) {
            if ($post->id === $id) {
                return $post;
            }
        }
        return [];
    }

    /**
     * Load data from storage
     */
    public static function load(): void
    {
        $DB_PATH = Config::get('DB_PATH', __DIR__ . '/../../database/');
        $STORAGE_FRONT = Config::get('STORAGE_FRONT', '');
        $images = json_decode(file_get_contents($DB_PATH . 'storage.json'));
        $response = [];
        foreach ($images as $image){
            $image->src = "http://" . $_SERVER["HTTP_HOST"] . $STORAGE_FRONT . '/storage/' . $image->name;
            $response[] = $image;
        }
        self::$DATA = $response;
    }

    /**
     * Persist data to json storage
     */
    public static function save(): void
    {
        $DB_PATH = Config::get('DB_PATH', __DIR__ . '/../../database/');
        file_put_contents($DB_PATH . 'storage.json', json_encode(self::$DATA, JSON_PRETTY_PRINT));
    }

    /**
     * @param int $id
     */
    public static function remove(int $id): void
    {
        $new = [];
        foreach (self::$DATA as $file) {
            if ($file->id !== $id) {
                $new[] = $file;
            } else {
                $STORAGE_PATH = Config::get('STORAGE_PATH', __DIR__ . '/../../storage/');
                unlink($STORAGE_PATH . $file->name);
            }
        }

        self::$DATA = $new;
        self::save();
    }

    /**
     * @param $id
     * @param $gallery
     */
    public static function edit($id, $gallery): void
    {
        $data = [];
        foreach (self::$DATA as $item) {
            if ((int) $item->id === (int) $id) {
                $item->gallery = $gallery;
                $data[] = $item;
            } else {
                $data[] = $item;
            }
        }
        self::$DATA = $data;
        self::save();
    }
}
