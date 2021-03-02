<?php

namespace ArualCms\Controller;

use ArualCms\Lib\Response;
use ArualCms\Model\Posts;

class PostsController
{
    public function __construct()
    {
        Posts::load();
    }

    public function getPosts(Response $res)
    {
        $data['posts'] = Posts::all();
        $data['success'] = true;
        $res->toJSON($data);
    }

    public function getPost(int $id, Response $res)
    {
        $post = Posts::findById($id);
        if ($post) {
            $res->toJSON($post);
        } else {
            $res->status(404)->toJSON(['error' => "Not Found"]);
        }
    }

    public function add($post, Response $res)
    {
        Posts::add($post);
        $res->toJSON(['success' => true, 'message' => 'Record added successfully']);
    }

    public function edit($post, Response $res)
    {
        Posts::edit($post);
        $res->toJSON(['success' => true, 'message' => 'Record updated successfully']);
    }

    public function remove($id, Response $res)
    {
        Posts::remove($id);
        $res->toJSON(['success' => true, 'message' => 'Record removed successfully']);
    }
}
