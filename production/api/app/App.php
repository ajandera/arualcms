<?php
declare(strict_types = 1);

namespace ArualCms;

use ArualCms\Controller\AuthController;
use ArualCms\Controller\FileController;
use ArualCms\Controller\MailController;
use ArualCms\Controller\PostsController;
use ArualCms\Controller\SettingsController;
use ArualCms\Controller\TextsController;
use ArualCms\Controller\UsersController;
use ArualCms\Lib\Guard;
use ArualCms\Lib\Logger;
use ArualCms\Lib\Request;
use ArualCms\Lib\Response;
use ArualCms\Lib\Router;

/**
 * Class App
 * @package ArualCms
 */
class App
{
    /**
     * Start application
     */
    public static function run(): void
    {
        Router::get('/api/', function (Request $req, Response $res) {
            $data['message'] = 'arualCMS is running';
            $data['success'] = true;
            $res->toJSON($data);
        });

        //posts
        Router::get('/api/post', function (Request $req, Response $res) {
            (new PostsController())->getPosts($res);
        });

        Router::get('/api/post/([0-9]*)', function (Request $req, Response $res) {
            (new PostsController())->getPost((int) $req->params[0], $res);
        });

        Router::put('/api/post/([0-9]*)', function (Request $req, Response $res) {
            $bearer = explode(" ", getallheaders()['Authorization']);
            Guard::check($bearer[1], $res);
            (new PostsController())->edit($req->getJSON(), $res);
        });

        Router::post('/api/post', function (Request $req, Response $res) {
            $bearer = explode(" ", getallheaders()['Authorization']);
            Guard::check($bearer[1], $res);
            (new PostsController())->add($req->getJSON(), $res);
        });

        Router::delete('/api/post/([0-9]*)', function (Request $req, Response $res) {
            $bearer = explode(" ", getallheaders()['Authorization']);
            Guard::check($bearer[1], $res);
            (new PostsController())->remove((int) $req->params[0], $res);
        });

        //settings
        Router::get('/api/setting', function (Request $req, Response $res) {
            (new SettingsController())->getSettings($res);
        });

        Router::get('/api/setting/([a-z]*)', function (Request $req, Response $res) {
            (new SettingsController())->getSetting($req->params[0], $res);
        });

        Router::put('/api/setting', function (Request $req, Response $res) {
            $bearer = explode(" ", getallheaders()['Authorization']);
            Guard::check($bearer[1], $res);
            (new SettingsController())->save($req->getJSON(), $res);
        });

        //texts
        Router::get('/api/text', function (Request $req, Response $res) {
            (new TextsController())->getTexts($res);
        });

        Router::get('/api/text/([a-z]*)', function (Request $req, Response $res) {
            (new TextsController())->getText($req->params[0], $res);
        });

        Router::put('/api/text', function (Request $req, Response $res) {
            $bearer = explode(" ", getallheaders()['Authorization']);
            Guard::check($bearer[1], $res);
            (new TextsController())->save($req->getJSON(), $res);
        });

        //users
        Router::get('/api/user', function (Request $req, Response $res) {
            (new UsersController())->getUsers($res);
        });

        Router::get('/api/users/id/([0-9]*)', function (Request $req, Response $res) {
            (new UsersController())->getUser((int) $req->params[0], $res);
        });

        Router::get('/api/users/username/([a-zA-Z]*)', function (Request $req, Response $res) {
            (new UsersController())->getUserByUsername($req->params[0], $res);
        });

        Router::put('/api/user/([0-9]*)', function (Request $req, Response $res) {
            $bearer = explode(" ", getallheaders()['Authorization']);
            Guard::check($bearer[1], $res);
            (new UsersController())->edit($req->getJSON(), $res);
        });

        Router::post('/api/user', function (Request $req, Response $res) {
            $bearer = explode(" ", getallheaders()['Authorization']);
            Guard::check($bearer[1], $res);
            (new UsersController())->add($req->getJSON(), $res);
        });

        Router::delete('/api/user/([0-9]*)', function (Request $req, Response $res) {
            $bearer = explode(" ", getallheaders()['Authorization']);
            Guard::check($bearer[1], $res);
            (new UsersController())->remove((int) $req->params[0], $res);
        });

        // files
        Router::get('/api/files', function (Request $req, Response $res) {
            (new FileController())->getFiles($res);
        });

        Router::get('/api/files/([0-9]*)', function (Request $req, Response $res) {
            (new FileController())->getFile((int) $req->params[0], $res);
        });

        Router::post('/api/files/upload', function (Request $req, Response $res) {
            $bearer = explode(" ", getallheaders()['Authorization']);
            Guard::check($bearer[1], $res);
            (new FileController())->save($res);
        });

        Router::put('/api/files/gallery/([0-9]*)', function (Request $req, Response $res) {
            $bearer = explode(" ", getallheaders()['Authorization']);
            Guard::check($bearer[1], $res);
            (new FileController())->addGallery((int) $req->params[0], $req->getJSON(), $res);
        });

        Router::delete('/api/files/([0-9]*)', function (Request $req, Response $res) {
            $bearer = explode(" ", getallheaders()['Authorization']);
            Guard::check($bearer[1], $res);
            (new FileController())->remove((int) $req->params[0], $res);
        });

        // authorization
        Router::post('/api/auth', function (Request $req, Response $res) {
            (new AuthController())->auth($req->getJSON(), $res);
        });

        // recovery
        Router::post('/api/recovery', function (Request $req, Response $res) {
            (new AuthController())->recovery($req->getJSON(), $res);
        });

        // email
        Router::post('/api/mail', function (Request $req, Response $res) {
            (new MailController())->send($req->getJSON(), $res);
        });

        Logger::enableSystemLogs();
    }
}
