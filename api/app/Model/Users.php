<?php
declare(strict_types= 1);

namespace ArualCms\Model;

use ArualCms\Lib\Config;

/**
 * Class Users
 * @package ArualCms\Model
 */
class Users
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
     * @param string $username
     * @return array|mixed
     */
    public static function findByUsername(string $username)
    {
        foreach (self::$DATA as $user) {
            if ($user->username === $username) {
                return $user;
            }
        }
        return [];
    }

    /**
     * @return int|mixed
     */
    private static function getNextId()
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
    public static function findById(int $id)
    {
        foreach (self::$DATA as $user) {
            if ($user->id === $id) {
                return $user;
            }
        }
        return [];
    }

    /**
     * @param $user
     */
    public static function add($user): void
    {
        $user->id = self::getNextId();
        self::$DATA[] = $user;
        self::save();
    }

    /**
     * @param $user
     */
    public static function edit($user): void
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

    /**
     * @param int $id
     */
    public static function remove(int $id): void
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

    /**
     * Load users from storage
     */
    public static function load(): void
    {
        $DB_PATH = Config::get('DB_PATH', __DIR__ . '/../../database/');
        self::$DATA = json_decode(file_get_contents($DB_PATH . 'users.json'));
    }

    /**
     * Save users to storage
     */
    public static function save(): void
    {
        $DB_PATH = Config::get('DB_PATH', __DIR__ . '/../../database/');
        file_put_contents($DB_PATH . 'users.json', json_encode(self::$DATA, JSON_PRETTY_PRINT));
    }
}
