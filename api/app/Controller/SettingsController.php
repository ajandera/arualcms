<?php
declare(strict_types = 1);

namespace ArualCms\Controller;

use ArualCms\Lib\Response;
use ArualCms\Model\Settings;

/**
 * Class SettingsController
 * @package ArualCms\Controller
 */
class SettingsController
{
    use Settings;

    /**
     * @param Response $res
     */
    public function getSettings(Response $res): void
    {
        $data['settings'] = $this->allSetting();
        $data['success'] = true;
        $res->toJSON($data);
    }

    /**
     * @param string $key
     * @param Response $res
     */
    public function getSetting(string $key, Response $res): void
    {
        $setting = $this->findByKey($key);
        if ($setting) {
            $res->toJSON($setting);
        } else {
            $res->status(404)->toJSON(['error' => "Not Found"]);
        }
    }

    /**
     * @param $settings
     * @param Response $res
     */
    public function save($settings, Response $res): void
    {
        $this->editSetting($settings);
        $res->toJSON(['success' => true, 'message' => 'Record updated successfully']);
    }
}
