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
    use Posts;

    /**
     * @param Response $res
     */
    public function getPosts(Response $res): void
    {
        $data['posts'] = $this->all();
        $data['success'] = true;
        $res->toJSON($data);
    }

    /**
     * @param string $id
     * @param Response $res
     */
    public function getPost(string $id, Response $res): void
    {
        $post = $this->findById($id);
        if ($post) {
            $res->toJSON(['success' => true, 'post' => $post]);
        } else {
            $res->status(404)->toJSON(['error' => "Not Found"]);
        }
    }

    /**
     * @param $post
     * @param Response $res
     */
    public function addPost($post, Response $res): void
    {
        $this->add($post);
        $res->toJSON(['success' => true, 'message' => 'Record added successfully']);
    }

    /**
     * @param $post
     * @param Response $res
     */
    public function editPost($post, Response $res): void
    {
        $this->edit($post);
        $res->toJSON(['success' => true, 'message' => 'Record updated successfully']);
    }

    /**
     * @param string $id
     * @param Response $res
     */
    public function removePost(string $id, Response $res): void
    {
        $this->remove($id);
        $res->toJSON(['success' => true, 'message' => 'Record removed successfully']);
    }
}
