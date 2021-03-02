<?php
require __DIR__ . '/vendor/autoload.php';

use ArualCms\App;
use ArualCms\Controller\AuthController;
use ArualCms\Controller\PostsController;
use ArualCms\Controller\SettingsController;
use ArualCms\Controller\TextsController;
use ArualCms\Controller\UsersController;
use ArualCms\Lib\Request;
use ArualCms\Lib\Response;
use ArualCms\Lib\Router;

//posts
Router::get('/post', function (Request $req, Response $res) {
    (new PostsController())->getPosts($res);
});

Router::get('/post/([0-9]*)', function (Request $req, Response $res) {
    (new PostsController())->getPost($req->params[0], $res);
});

Router::put('/post/([0-9]*)', function (Request $req, Response $res) {
    (new PostsController())->edit($req->getJSON(), $res);
});

Router::post('/post', function (Request $req, Response $res) {
    (new PostsController())->add($req->getJSON(), $res);
});

Router::delete('/post/([0-9]*)', function (Request $req, Response $res) {
    (new PostsController())->remove($req->params[0], $res);
});

//settings
Router::get('/setting', function (Request $req, Response $res) {
    (new SettingsController())->getSettings($res);
});

Router::get('/setting/([a-z]*)', function (Request $req, Response $res) {
    (new SettingsController())->getSetting($req->params[0], $res);
});

Router::put('/setting', function (Request $req, Response $res) {
    (new SettingsController())->save($req->getJSON(), $res);
});

//texts
Router::get('/text', function (Request $req, Response $res) {
    (new TextsController())->getTexts($res);
});

Router::get('/text/([a-z]*)', function (Request $req, Response $res) {
    (new TextsController())->getText($req->params[0], $res);
});

Router::put('/text', function (Request $req, Response $res) {
    (new TextsController())->save($req->getJSON(), $res);
});

//users
Router::get('/user', function (Request $req, Response $res) {
    (new UsersController())->getUsers($res);
});

Router::get('/users/id/([0-9]*)', function (Request $req, Response $res) {
    (new UsersController())->getUser($req->params[0], $res);
});

Router::get('/users/username/([a-z]*)', function (Request $req, Response $res) {
    (new UsersController())->getUserByUsername($req->params[0], $res);
});

Router::put('/user/([0-9]*)', function (Request $req, Response $res) {
    (new UsersController())->edit($req->getJSON(), $res);
});

Router::post('/user', function (Request $req, Response $res) {
    (new UsersController())->add($req->getJSON(), $res);
});

Router::delete('/user/([0-9]*)', function (Request $req, Response $res) {
    (new UsersController())->remove($req->params[0], $res);
});

// authorization
Router::post('/auth', function (Request $req, Response $res) {
    (new AuthController())->auth($req->getJSON(), $res);
});

App::run();
