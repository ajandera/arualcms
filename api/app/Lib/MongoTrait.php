<?php
declare(strict_types=1);

namespace ArualCms\Lib;

use MongoDB\BSON\ObjectId;
use MongoDB\Client;
use MongoDB\DeleteResult;

trait MongoTrait
{
    private function getClient()
    {
        return new Client("mongodb://root:3FSFSDFDFrew@mongo:27017");
    }

    public function findBy(string $collection, array $criteria = []): array
    {
        return $this->getClient()->selectCollection('arualcms', $collection)->find($criteria)->toArray();
    }

    public function insert(string $collection, \stdClass $data): \MongoDB\InsertOneResult
    {
        return $this->getClient()->selectCollection('arualcms',$collection)->insertOne($data);
    }

    public function update(string $collection, \stdClass $data, ObjectID $id): array|object
    {
        unset($data->_id);
        return $this->getClient()->selectCollection('arualcms', $collection)->findOneAndUpdate(
            ["_id" => $id],
            ['$set' => $data]
        );
    }

    public function delete(string $collection, array $criteria): DeleteResult
    {
        return $this->getClient()->selectCollection('arualcms',$collection)->deleteOne($criteria);
    }
}