<?php
require __DIR__ . '/vendor/autoload.php';

use ArualCms\App;
use ArualCms\Controller\PostsController;
use ArualCms\Lib\Request;
use ArualCms\Lib\Response;
use ArualCms\Lib\Router;

Router::get('/post', function (Request $req, Response $res) {
    (new PostsController())->getPosts($res);
});

Router::get('/post/([0-9]*)', function (Request $req, Response $res) {
    (new PostsController())->getPost($req->params[0], $res);
});

Router::put('/post/([0-9]*)', function (Request $req, Response $res) {
    (new PostsController())->save($req->getBody(), $res);
});

App::run();
