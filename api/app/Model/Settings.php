<?php
declare(strict_types = 1);

namespace ArualCms\Model;

use ArualCms\Lib\MongoTrait;
use MongoDB\BSON\ObjectId;

/**
 * Class Settings
 * @package ArualCms\Model
 */
trait Settings
{
    use MongoTrait;

    /**
     * @return array
     */
    public function allSetting(): array
    {
        return $this->findBy('settings');
    }

    /**
     * @param \stdClass $setting
     */
    public function editSetting(\stdClass $setting): void
    {
        $oid = '$oid';
        $this->update('settings', $setting, new ObjectID($setting->_id->$oid));
    }

    /**
     * @param \stdClass $setting
     */
    public function addSetting(\stdClass $setting): void
    {
        $this->insert('users', $setting);
    }

    /**
     * @param string $key
     * @return array
     */
    public function findByKey(string $key): array
    {
        return $this->findBy('settings',["key" => $key]);
    }
}
