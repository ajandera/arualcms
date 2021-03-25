<?php
declare(strict_types = 1);

namespace ArualCms\Model;

use ArualCms\Lib\Config;

/**
 * Class Posts
 * @package ArualCms\Model
 */
class Posts
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
     * @param $post
     */
    public static function add($post): void
    {
        $post->id = self::getNextId();
        self::$DATA[] = $post;
        self::save();
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
    public static function findById(int $id)
    {
        foreach (self::$DATA as $post) {
            if ($post->id === $id) {
                return $post;
            }
        }
        return [];
    }

    /**
     * Load posts from storage
     */
    public static function load(): void
    {
        $DB_PATH = Config::get('DB_PATH', __DIR__ . '/../../database/');
        self::$DATA = json_decode(file_get_contents($DB_PATH . 'posts.json'));
    }

    /**
     * Persist post to storage
     */
    public static function save(): void
    {
        $DB_PATH = Config::get('DB_PATH', __DIR__ . '/../../database/');
        file_put_contents($DB_PATH . 'posts.json', json_encode(self::$DATA, JSON_PRETTY_PRINT));
    }

    /**
     * @param $post
     */
    public static function edit($post): void
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

    /**
     * @param int $id
     */
    public static function remove(int $id): void
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
