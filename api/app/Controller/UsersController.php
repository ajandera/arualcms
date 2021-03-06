<?php

namespace ArualCms\Controller;

use ArualCms\Lib\Response;
use ArualCms\Model\Users;

class UsersController
{
    public function __construct()
    {
        Users::load();
    }

    public function getUsers(Response $res)
    {
        $data['users'] = Users::all();
        $data['success'] = true;
        $res->toJSON($data);
    }

    public function getUserByUsername(string $key, Response $res)
    {
        $user = Users::findByUsername($key);
        if ($user) {
            $res->toJSON($user);
        } else {
            $res->status(404)->toJSON(['error' => "Not Found"]);
        }
    }

    public function getUser(int $id, Response $res)
    {
        $user = Users::findById($id);
        if ($user) {
            $res->toJSON($user);
        } else {
            $res->status(404)->toJSON(['error' => "Not Found"]);
        }
    }

    public function add($user, Response $res)
    {
        $user->password = password_hash($user->password, PASSWORD_BCRYPT);
        Users::add($user);
        $res->toJSON(['success' => true, 'message' => 'Record added successfully']);
    }

    public function edit($user, Response $res)
    {
        $user->password = password_hash($user->password, PASSWORD_BCRYPT);
        Users::edit($user);
        $res->toJSON(['success' => true, 'message' => 'Record updated successfully']);
    }

    public function remove($id, Response $res)
    {
        Users::remove($id);
        $res->toJSON(['success' => true, 'message' => 'Record removed successfully']);
    }
}
