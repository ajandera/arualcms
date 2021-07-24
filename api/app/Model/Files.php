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
        $STORAGE_FRONT = Config::get('STORAGE_FRONT', '');
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

        $STORAGE_FRONT = Config::get('STORAGE_FRONT', '');
        return ['link' => 'http://' . $_SERVER["HTTP_HOST"] . $STORAGE_FRONT . '/storage/' . $file->name, 'name' => $file->name];
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
        $STORAGE_PATH = Config::get('STORAGE_PATH', __DIR__ . '/../../storage/');
        unlink($STORAGE_PATH . $file[0]['name']);
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
