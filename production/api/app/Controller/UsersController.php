<?php
declare(strict_types = 1);

namespace ArualCms\Controller;

use ArualCms\Lib\Response;
use ArualCms\Model\Users;

/**
 * Class UsersController
 * @package ArualCms\Controller
 */
class UsersController
{
    /**
     * UsersController constructor.
     */
    public function __construct()
    {
        Users::load();
    }

    /**
     * @param Response $res
     */
    public function getUsers(Response $res): void
    {
        $data['users'] = Users::all();
        $data['success'] = true;
        $res->toJSON($data);
    }

    /**
     * @param string $key
     * @param Response $res
     */
    public function getUserByUsername(string $key, Response $res): void
    {
        $user = Users::findByUsername($key);
        if ($user) {
            $res->toJSON($user);
        } else {
            $res->status(404)->toJSON(['error' => "Not Found"]);
        }
    }

    /**
     * @param int $id
     * @param Response $res
     */
    public function getUser(int $id, Response $res): void
    {
        $user = Users::findById($id);
        if ($user) {
            $res->toJSON($user);
        } else {
            $res->status(404)->toJSON(['error' => "Not Found"]);
        }
    }

    /**
     * @param $user
     * @param Response $res
     */
    public function add($user, Response $res): void
    {
        $user->password = password_hash($user->password, PASSWORD_BCRYPT);
        Users::add($user);
        $res->toJSON(['success' => true, 'message' => 'Record added successfully']);
    }

    /**
     * @param $user
     * @param Response $res
     */
    public function edit($user, Response $res): void
    {
        $user->password = password_hash($user->password, PASSWORD_BCRYPT);
        Users::edit($user);
        $res->toJSON(['success' => true, 'message' => 'Record updated successfully']);
    }

    /**
     * @param $id
     * @param Response $res
     */
    public function remove($id, Response $res): void
    {
        Users::remove($id);
        $res->toJSON(['success' => true, 'message' => 'Record removed successfully']);
    }
}
