<?php
declare(strict_types= 1);

namespace ArualCms\Model;

use ArualCms\Lib\MongoTrait;
use MongoDB\BSON\ObjectId;

/**
 * Class Users
 * @package ArualCms\Model
 */
trait Users
{
    use MongoTrait;

    /**
     * @return array
     * @throws \Exception
     */
    public function all(): array
    {
        return $this->findBy('users');
    }

    /**
     * @param string $username
     * @return array
     * @throws \Exception
     */
    public function findByUsername(string $username): array
    {
        return $this->findBy('users',["username" => $username]);
    }

    /**
     * @param int $id
     * @return array
     */
    public function findById(string $id): array
    {
        return $this->findBy('users',["_id" => new ObjectID($id)]);
    }

    /**
     * @param \stdClass $user
     */
    public function add(\stdClass $user): void
    {
        $this->insert('users', $user);
    }

    /**
     * @param array $user
     */
    public function edit(\stdClass $user): void
    {
        $oid = '$oid';
        $this->update('users', $user, new ObjectID($user->_id->$oid));
    }

    /**
     * @param string $id
     */
    public function remove(string $id): void
    {
        $this->delete('users', ["_id" => new ObjectID($id)]);
    }
}
