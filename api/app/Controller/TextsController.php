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
    use Texts;

    /**
     * @param Response $res
     */
    public function getTexts(Response $res): void
    {
        $data['texts'] = $this->all();
        $data['success'] = true;
        $res->toJSON($data);
    }

    /**
     * @param string $key
     * @param Response $res
     */
    public function getText(string $key, Response $res): void
    {
        $text = $this->findByKey($key);
        if ($text) {
            $res->toJSON($text);
        } else {
            $res->status(404)->toJSON(['error' => "Not Found"]);
        }
    }

    /**
     * @param \stdClass $text
     * @param Response $res
     */
    public function addText(\stdClass $text, Response $res): void
    {
        $this->add($text);
        $res->toJSON(['success' => true, 'message' => 'Record added successfully']);
    }

    /**
     * @param \stdClass $text
     * @param Response $res
     */
    public function editText(\stdClass $text, Response $res): void
    {
        $this->edit($text);
        $res->toJSON(['success' => true, 'message' => 'Record updated successfully']);
    }

    /**
     * @param string $id
     * @param Response $res
     */
    public function removeText(string $id, Response $res): void
    {
        $this->remove($id);
        $res->toJSON(['success' => true, 'message' => 'Record removed successfully']);
    }
}
