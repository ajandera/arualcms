<?php

namespace ArualCms\Controller;

use ArualCms\Lib\Response;
use ArualCms\Model\Texts;

class TextsController
{
    public function __construct()
    {
        Texts::load();
    }

    public function getTexts(Response $res)
    {
        $data['texts'] = Texts::all();
        $data['success'] = true;
        $res->toJSON($data);
    }

    public function getText(string $key, Response $res)
    {
        $text = Texts::findByKey($key);
        if ($text) {
            $res->toJSON($text);
        } else {
            $res->status(404)->toJSON(['error' => "Not Found"]);
        }
    }

    public function save($texts, Response $res)
    {
        Texts::update($texts);
        $res->toJSON(['success' => true, 'message' => 'Record added successfully']);
    }
}
