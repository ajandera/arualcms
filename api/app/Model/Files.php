<?php

namespace ArualCms\Model;

use ArualCms\Lib\Config;

class Files
{
    private static $DATA = [];

    public static function all()
    {
        return self::$DATA;
    }

    public static function add($file)
    {
        $file['id'] = self::getNextId();
        self::$DATA[] = $file;
        self::save();
    }

    private static function getNextId()
    {
        $ids = [];
        foreach (self::$DATA as $data) {
            $ids[] = $data->id;
        }
        return max($ids) + 1;
    }

    public static function findById(int $id)
    {
        foreach (self::$DATA as $post) {
            if ($post->id === $id) {
                return $post;
            }
        }
        return [];
    }

    public static function load()
    {
        $DB_PATH = Config::get('DB_PATH', __DIR__ . '/../../database/');
        $images = json_decode(file_get_contents($DB_PATH . 'storage.json'));
        $response = [];
        foreach ($images as $image){
            $image->src = "http://" . $_SERVER["HTTP_HOST"] . '/storage/' . $image->name;
            $response[] = $image;
        }
        self::$DATA = $response;
    }

    public static function save()
    {
        $DB_PATH = Config::get('DB_PATH', __DIR__ . '/../../database/');
        file_put_contents($DB_PATH . 'storage.json', json_encode(self::$DATA, JSON_PRETTY_PRINT));
    }

    public static function remove(int $id)
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
}
