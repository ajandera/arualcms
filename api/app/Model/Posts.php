<?php
declare(strict_types = 1);

namespace ArualCms\Model;

use ArualCms\Lib\MongoTrait;
use MongoDB\BSON\ObjectId;

/**
 * Class Posts
 * @package ArualCms\Model
 */
trait Posts
{
    use MongoTrait;

    /**
     * @return array
     */
    public function all(): array
    {
        return $this->findBy('posts');
    }

    /**
     * @param \stdClass $post
     */
    public function add(\stdClass $post): void
    {
        $this->insert('posts', $post);
    }

    /**
     * @param string $id
     * @return array
     */
    public function findById(string $id): array
    {
        return $this->findBy('posts',["_id" => new ObjectID($id)]);
    }

    /**
     * @param \stdClass $post
     */
    public function edit(\stdClass $post): void
    {
        $oid = '$oid';
        $this->update('posts', $post, new ObjectID($post->_id->$oid));
    }

    /**
     * @param string $id
     */
    public function remove(string $id): void
    {
        $this->delete('posts', ["_id" => new ObjectID($id)]);
    }
}
