<?php
declare(strict_types = 1);

namespace ArualCms\Model;

use ArualCms\Lib\MongoTrait;
use MongoDB\BSON\ObjectId;

/**
 * Class Texts
 * @package ArualCms\Model
 */
trait Texts
{
    use MongoTrait;

    /**
     * @return array
     */
    public function all(): array
    {
        return $this->findBy('texts');
    }

    /**
     * @param \stdClass $text
     */
    public function add(\stdClass $text): void
    {
        $this->insert('texts', $text);
    }

    /**
     * @param \stdClass $text
     */
    public function edit(\stdClass $text): void
    {
        $oid = '$oid';
        $this->update('texts', $text, new ObjectID($text->_id->$oid));
    }

    /**
     * @param string $id
     */
    public function remove(string $id): void
    {
        $this->delete('texts', ["_id" => new ObjectID($id)]);
    }

    /**
     * @param string $key
     * @return array
     */
    public function findByKey(string $key): array
    {
        return $this->findBy('texts',["key" => $key]);
    }
}
