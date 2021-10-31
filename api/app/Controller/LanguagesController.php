<?php
declare(strict_types = 1);

namespace ArualCms\Controller;

use ArualCms\Lib\Response;
use ArualCms\Model\Languages;
use ArualCms\Model\Files;
use ArualCms\Model\Posts;
use ArualCms\Model\Setting;
use ArualCms\Model\Texts;

/**
 * Class LanguagesController
 * @package ArualCms\Controller
 */
class LanguagesController
{
    use Languages;

    /**
     * @param Response $res
     */
    public function getLanguages(Response $res): void
    {
        $data['languages'] = $this->all();
        $data['success'] = true;
        $res->toJSON($data);
    }

    /**
     * @param string $key
     * @param Response $res
     */
    public function getLanguage(string $key, Response $res): void
    {
        $text = $this->findByKey($key);
        if ($text) {
            $res->toJSON($text);
        } else {
            $res->status(404)->toJSON(['error' => "Not Found"]);
        }
    }

    /**
     * @param \stdClass $language
     * @param Response $res
     */
    public function addLanguage(\stdClass $language, Response $res): void
    {
        $this->add($language);
        $res->toJSON(['success' => true, 'message' => 'Record added successfully']);
    }

    /**
     * @param \stdClass $language
     * @param Response $res
     */
    public function editLanguage(\stdClass $language, Response $res): void
    {
        if ($language->default === 1) {
            // todo unset default from others
        }
        $this->edit($language);
        $res->toJSON(['success' => true, 'message' => 'Record updated successfully']);
    }

    /**
     * @param string $id
     * @param Response $res
     */
    public function removeLanguage(string $id, Response $res): void
    {   
        $this->remove($id);
        $res->toJSON(['success' => true, 'message' => 'Record removed successfully']);
    }
}
