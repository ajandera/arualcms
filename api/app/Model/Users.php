<?php

namespace ArualCms\Model;

use ArualCms\Lib\Config;

class Users
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

    public static function findByUsername(string $username)
    {
        foreach (self::$DATA as $user) {
            if ($user->key === $username) {
                return $user;
            }
        }
        return [];
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
        foreach (self::$DATA as $user) {
            if ($user->id === $id) {
                return $user;
            }
        }
        return [];
    }

    public static function add($user)
    {
        $user->id = self::getNextId();
        self::$DATA[] = $user;
        self::save();
    }

    public static function edit($user)
    {
        if (!isset($user->id)) {
            self::add($user);
            return;
        }
        $data = [];
        foreach (self::$DATA as $item) {
            if ($item->id === $user->id) {
                $data[] = $user;
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
        foreach (self::$DATA as $user) {
            if ($user->id !== $id) {
                $new[] = $user;
            }
        }

        self::$DATA = $new;
        self::save();
    }

    public static function load()
    {
        $DB_PATH = Config::get('DB_PATH', __DIR__ . '/../database/');
        self::$DATA = json_decode(file_get_contents($DB_PATH . 'users.json'));
    }

    public static function save()
    {
        $DB_PATH = Config::get('DB_PATH', __DIR__ . '/../database/');
        file_put_contents($DB_PATH . 'users.json', json_encode(self::$DATA, JSON_PRETTY_PRINT));
    }
}
