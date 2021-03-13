<?php
declare(strict_types = 1);

namespace ArualCms\Controller;

use ArualCms\Lib\Response;
use ArualCms\Model\Posts;

/**
 * Class PostsController
 * @package ArualCms\Controller
 */
class PostsController
{
    /**
     * PostsController constructor.
     */
    public function __construct()
    {
        Posts::load();
    }

    /**
     * @param Response $res
     */
    public function getPosts(Response $res): void
    {
        $data['posts'] = Posts::all();
        $data['success'] = true;
        $res->toJSON($data);
    }

    /**
     * @param int $id
     * @param Response $res
     */
    public function getPost(int $id, Response $res): void
    {
        $post = Posts::findById($id);
        if ($post) {
            $res->toJSON($post);
        } else {
            $res->status(404)->toJSON(['error' => "Not Found"]);
        }
    }

    /**
     * @param $post
     * @param Response $res
     */
    public function add($post, Response $res): void
    {
        Posts::add($post);
        $res->toJSON(['success' => true, 'message' => 'Record added successfully']);
    }

    /**
     * @param $post
     * @param Response $res
     */
    public function edit($post, Response $res): void
    {
        Posts::edit($post);
        $res->toJSON(['success' => true, 'message' => 'Record updated successfully']);
    }

    /**
     * @param $id
     * @param Response $res
     */
    public function remove($id, Response $res): void
    {
        Posts::remove($id);
        $res->toJSON(['success' => true, 'message' => 'Record removed successfully']);
    }
}
