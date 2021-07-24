<?php
declare(strict_types = 1);

namespace ArualCms\Model;

use ArualCms\Lib\Config;
use ArualCms\Lib\MongoTrait;
use MongoDB\BSON\ObjectId;

/**
 * Class Languages
 * @package ArualCms\Model
 */
trait Languages
{
    use MongoTrait;

    /**
     * @return array
     */
    public function all(): array
    {
        return $this->findBy('languages');
    }

    /**
     * @param \stdClass $language
     */
    public function add(\stdClass $language): void
    {
        $this->insert('languages', $language);
    }

    /**
     * @param \stdClass $language
     */
    public function edit(\stdClass $language): void
    {
        $oid = '$oid';
        $this->update('languages', $language, new ObjectID($language->_id->$oid));
    }

    /**
     * @param string $id
     */
    public function remove(string $id): void
    {
        $this->delete('languages', ["_id" => new ObjectID($id)]);
    }

    /**
     * @param string $key
     * @return array
     */
    public function findByKey(string $key): array
    {
        return $this->findBy('languages',["key" => $key]);
    }
}
