<?php

namespace ArualCms\Controller;

use ArualCms\Lib\Response;
use ArualCms\Model\Settings;

class SettingsController
{
    public function __construct()
    {
        Settings::load();
    }

    public function getSettings(Response $res)
    {
        $data['settings'] = Settings::all();
        $data['success'] = true;
        $res->toJSON($data);
    }

    public function getSetting(string $key, Response $res)
    {
        $setting = Settings::findByKey($key);
        if ($setting) {
            $res->toJSON($setting);
        } else {
            $res->status(404)->toJSON(['error' => "Not Found"]);
        }
    }

    public function save($settings, Response $res)
    {
        Settings::update($settings);
        $res->toJSON(['success' => true, 'message' => 'Record updated successfully']);
    }
}
