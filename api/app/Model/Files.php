<?php
declare(strict_types = 1);

namespace ArualCms\Model;

use ArualCms\Lib\Config;
use ArualCms\Lib\MongoTrait;
use MongoDB\BSON\ObjectId;

/**
 * Class Files
 * @package ArualCms\Model
 */
trait Files
{
    use MongoTrait;

    /**
     * @return array
     */
    public function all(): array
    {
        $files = $this->findBy('files');
        $response = [];
        foreach ($files as $image){
            $image->src = "https://" . $_SERVER["HTTP_HOST"] . $STORAGE_FRONT . '/storage/' . $image->name;
            $response[] = $image;
        }
        return $response;
    }

    /**
     * @param $file
     * @return array
     */
    public function add(\stdClass $file): array
    {
        $this->insert('files', $file);
        return $file->name;
    }

    /**
     * @param string $id
     * @return array
     */
    public function findById(string $id): array
    {
        return $this->findBy('files',["_id" =>  new ObjectID($id)]);
    }

    /**
     * @param string $id
     */
    public function remove(string $id): void
    {
        $file = $this->findById($id);
        $this->delete('files', ["_id" => new ObjectID($id)]);
    }

    /**
     * @param string $id
     * @param \stdClass $file
     */
    public function edit(string $id, \stdClass $file): void
    {
        $this->update('files', $file, new ObjectID($id));
    }
}
