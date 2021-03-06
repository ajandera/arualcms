<?php

namespace ArualCms\Controller;

use ArualCms\Lib\Config;
use ArualCms\Lib\Response;
use ArualCms\Model\Files;

class FileController
{
    public function __construct()
    {
        Files::load();
    }

    public function getFiles(Response $res)
    {
        $data['files'] = Files::all();
        $data['success'] = true;
        $res->toJSON($data);
    }

    public function getFile(string $id, Response $res)
    {
        $file = Files::findById($id);
        if ($file) {
            $res->toJSON($file);
        } else {
            $res->status(404)->toJSON(['error' => "Not Found"]);
        }
    }

    public function save(Response $res)
    {
        self::upload();
        $res->toJSON(['success' => true, 'message' => 'Record added successfully']);
    }

    public function remove($id, Response $res)
    {
        Files::remove($id);
        $res->toJSON(['success' => true, 'message' => 'Record removed successfully']);
    }

    public function upload()
    {
        $targetDir = Config::get('STORAGE_PATH',__DIR__ . '../../storage/');
        $targetFile = $targetDir . basename($_FILES["file"]["name"]);
        $imageFileType = strtolower(pathinfo($targetFile,PATHINFO_EXTENSION)); // todo check filesize
        move_uploaded_file($_FILES["file"]["tmp_name"], $targetFile);
        Files::add([
            "name" => basename($_FILES["file"]["name"])
        ]);
    }
}
