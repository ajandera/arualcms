<?php
declare(strict_types = 1);

namespace ArualCms\Controller;

use ArualCms\Lib\Config;
use ArualCms\Lib\Response;
use ArualCms\Model\Files;

/**
 * Class FileController
 * @package ArualCms\Controller
 */
class FileController
{
    use Files;

    /**
     * @param Response $res
     */
    public function getFiles(Response $res): void
    {
        $data['files'] = $this->all();
        $data['success'] = true;
        $res->toJSON($data);
    }

    /**
     * @param string $id
     * @param Response $res
     */
    public function getFile(string $id, Response $res): void
    {
        $file = $this->findById($id);
        if ($file) {
            $res->toJSON($file);
        } else {
            $res->status(404)->toJSON(['error' => "Not Found"]);
        }
    }

    /**
     * @param Response $res
     */
    public function save(Response $res): void
    {
        $file = self::upload();
        $res->toJSON(['success' => true, 'message' => 'Record added successfully', 'file' => $file]);
    }

    /**
     * @param $id
     * @param Response $res
     */
    public function removeFile($id, Response $res): void
    {
        $this->remove($id);
        $res->toJSON(['success' => true, 'message' => 'Record removed successfully']);
    }

    /**
     * @return array
     */
    public function upload(): array
    {
        $targetDir = Config::get('STORAGE_PATH', __DIR__ . '../../storage/');
        $targetFile = $targetDir . basename($_FILES["file"]["name"]);
        $imageFileType = strtolower(pathinfo($targetFile,PATHINFO_EXTENSION)); // todo check filesize
        move_uploaded_file($_FILES["file"]["tmp_name"], $targetFile);
        $data = new \stdClass();
        $data->name = basename($_FILES["file"]["name"]);
        $data->gallery = "";
        return $this->add($data);
    }

    /**
     * @param \stdClass $data
     * @param Response $res
     */
    public function addGallery(string $id, \stdClass $data, Response $res): void
    {
        $this->edit($id, $data);
        $res->toJSON(['success' => true, 'message' => 'Record updated successfully']);
    }
}
