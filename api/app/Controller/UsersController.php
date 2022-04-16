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
    use Users;

    /**
     * @param Response $res
     * @throws \Exception
     */
    public function getUsers(Response $res): void
    {
        $data['users'] = $this->all();
        $data['success'] = true;
        $res->toJSON($data);
    }

    /**
     * @param string $key
     * @param Response $res
     * @throws \Exception
     */
    public function getUserByUsername(string $key, Response $res): void
    {
        $user = $this->findByUsername($key);
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
        $user = $this->findById((string) $id);
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
    public function addUser($user, Response $res): void
    {
        $user->password = password_hash($user->password, PASSWORD_BCRYPT);
        $this->add($user);
        $res->toJSON(['success' => true, 'message' => 'Record added successfully']);
    }

    /**
     * @param $user
     * @param Response $res
     */
    public function editUser($user, Response $res): void
    {
        $user->password = password_hash($user->password, PASSWORD_BCRYPT);
        $this->edit($user);
        $res->toJSON(['success' => true, 'message' => 'Record updated successfully']);
    }

    /**
     * @param $id
     * @param Response $res
     */
    public function removeUser($id, Response $res): void
    {
        $this->remove($id);
        $res->toJSON(['success' => true, 'message' => 'Record removed successfully']);
    }
}
