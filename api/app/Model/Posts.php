<?php

namespace ArualCms\Model;

use ArualCms\Lib\Config;

class Posts
{
    private static $DATA = [];

    public static function all()
    {
        return self::$DATA;
    }

    public static function add($post)
    {
        $post->id = self::getNextId();
        self::$DATA[] = $post;
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
        $DB_PATH = Config::get('DB_PATH', __DIR__ . '/../database/');
        self::$DATA = json_decode(file_get_contents($DB_PATH . 'posts.json'));
    }

    public static function save()
    {
        $DB_PATH = Config::get('DB_PATH', __DIR__ . '/../database/');
        file_put_contents($DB_PATH . 'posts.json', json_encode(self::$DATA, JSON_PRETTY_PRINT));
    }

    public static function edit($post)
    {
        if (!isset($post->id)) {
            self::add($post);
            return;
        }
        $data = [];
        foreach (self::$DATA as $item) {
            if ($item->id === $post->id) {
                $data[] = $post;
            } else {
                $data[] = $item;
            }
        }
        self::$DATA = $data;
        self::save();
    }

    public static function remove(int $id)
    {
        $new = [];
        foreach (self::$DATA as $post) {
            if ($post->id !== $id) {
                $new[] = $post;
            }
        }

        self::$DATA = $new;
        self::save();
    }
}
