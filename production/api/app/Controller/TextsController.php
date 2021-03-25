<?php
declare(strict_types = 1);

namespace ArualCms\Controller;

use ArualCms\Lib\Response;
use ArualCms\Model\Texts;

/**
 * Class TextsController
 * @package ArualCms\Controller
 */
class TextsController
{
    public function __construct()
    {
        Texts::load();
    }

    /**
     * @param Response $res
     */
    public function getTexts(Response $res): void
    {
        $data['texts'] = Texts::all();
        $data['success'] = true;
        $res->toJSON($data);
    }

    /**
     * @param string $key
     * @param Response $res
     */
    public function getText(string $key, Response $res): void
    {
        $text = Texts::findByKey($key);
        if ($text) {
            $res->toJSON($text);
        } else {
            $res->status(404)->toJSON(['error' => "Not Found"]);
        }
    }

    /**
     * @param $texts
     * @param Response $res
     */
    public function save($texts, Response $res): void
    {
        Texts::update($texts);
        $res->toJSON(['success' => true, 'message' => 'Record added successfully']);
    }
}
