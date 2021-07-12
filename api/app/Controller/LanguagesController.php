<?php
declare(strict_types = 1);

namespace ArualCms\Controller;

use ArualCms\Lib\Response;
use ArualCms\Model\Languages;

/**
 * Class LanguagesController
 * @package ArualCms\Controller
 */
class LanguagesController
{
    public function __construct()
    {
        Languages::load();
    }

    /**
     * @param Response $res
     */
    public function getLanguages(Response $res): void
    {
        $data['languages'] = Languages::all();
        $data['success'] = true;
        $res->toJSON($data);
    }

    /**
     * @param string $key
     * @param Response $res
     */
    public function getLanguage(string $key, Response $res): void
    {
        $text = Languages::findByKey($key);
        if ($text) {
            $res->toJSON($text);
        } else {
            $res->status(404)->toJSON(['error' => "Not Found"]);
        }
    }

    /**
     * @param $languages
     * @param Response $res
     */
    public function save($languages, Response $res): void
    {
        Languages::update($languages);
        $res->toJSON(['success' => true, 'message' => 'Record added successfully']);
    }
}
