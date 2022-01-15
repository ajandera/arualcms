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
        $dsn = Config::get('MONGO_DSN', 'mongodb://root:3FSFSDFDFrew@mongo:27017');
        return new Client($dsn);
    }

    public function findBy(string $collection, array $criteria = []): array
    {
        return $this->getClient()->selectCollection(Config::get("MONGO_DBNAME", 'arualcms'), $collection)->find($criteria)->toArray();
    }

    public function insert(string $collection, \stdClass $data): \MongoDB\InsertOneResult
    {
        return $this->getClient()->selectCollection(Config::get("MONGO_DBNAME", 'arualcms'),$collection)->insertOne($data);
    }

    public function update(string $collection, \stdClass $data, ObjectID $id): array|object
    {
        unset($data->_id);
        return $this->getClient()->selectCollection(Config::get("MONGO_DBNAME", 'arualcms'), $collection)->findOneAndUpdate(
            ["_id" => $id],
            ['$set' => $data]
        );
    }

    public function delete(string $collection, array $criteria): DeleteResult
    {
        return $this->getClient()->selectCollection(Config::get("MONGO_DBNAME", 'arualcms'),$collection)->deleteOne($criteria);
    }
}
