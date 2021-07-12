<?php
declare(strict_types = 1);

namespace ArualCms;

use ArualCms\Controller\AuthController;
use ArualCms\Controller\FileController;
use ArualCms\Controller\LanguagesController;
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
        Router::get('/', function (Request $req, Response $res) {
            $data['message'] = password_hash('Password123', PASSWORD_BCRYPT);//'arualCMS is running';
            $data['success'] = true;
            $res->toJSON($data);
        });

        //posts
        Router::get('/post', function (Request $req, Response $res) {
            (new PostsController())->getPosts($res);
        });

        Router::get('/post/([0-9]*)', function (Request $req, Response $res) {
            (new PostsController())->getPost((int) $req->params[0], $res);
        });

        Router::put('/post', function (Request $req, Response $res) {
            $bearer = explode(" ", getallheaders()['Authorization']);
            Guard::check($bearer[1], $res);
            (new PostsController())->edit($req->getJSON(), $res);
        });

        Router::post('/post', function (Request $req, Response $res) {
            $bearer = explode(" ", getallheaders()['Authorization']);
            Guard::check($bearer[1], $res);
            (new PostsController())->add($req->getJSON(), $res);
        });

        Router::delete('/post/([0-9]*)', function (Request $req, Response $res) {
            $bearer = explode(" ", getallheaders()['Authorization']);
            Guard::check($bearer[1], $res);
            (new PostsController())->remove((int) $req->params[0], $res);
        });

        //settings
        Router::get('/setting', function (Request $req, Response $res) {
            (new SettingsController())->getSettings($res);
        });

        Router::get('/setting/([a-z]*)', function (Request $req, Response $res) {
            (new SettingsController())->getSetting($req->params[0], $res);
        });

        Router::put('/setting', function (Request $req, Response $res) {
            $bearer = explode(" ", getallheaders()['Authorization']);
            Guard::check($bearer[1], $res);
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
            $bearer = explode(" ", getallheaders()['Authorization']);
            Guard::check($bearer[1], $res);
            (new TextsController())->save($req->getJSON(), $res);
        });

        //users
        Router::get('/user', function (Request $req, Response $res) {
            (new UsersController())->getUsers($res);
        });

        Router::get('/users/id/([0-9]*)', function (Request $req, Response $res) {
            (new UsersController())->getUser((int) $req->params[0], $res);
        });

        Router::get('/users/username/([a-zA-Z]*)', function (Request $req, Response $res) {
            (new UsersController())->getUserByUsername($req->params[0], $res);
        });

        Router::put('/user', function (Request $req, Response $res) {
            $bearer = explode(" ", getallheaders()['Authorization']);
            Guard::check($bearer[1], $res);
            (new UsersController())->edit($req->getJSON(), $res);
        });

        Router::post('/user', function (Request $req, Response $res) {
            $bearer = explode(" ", getallheaders()['Authorization']);
            Guard::check($bearer[1], $res);
            (new UsersController())->add($req->getJSON(), $res);
        });

        Router::delete('/user/([0-9]*)', function (Request $req, Response $res) {
            $bearer = explode(" ", getallheaders()['Authorization']);
            Guard::check($bearer[1], $res);
            (new UsersController())->remove((int) $req->params[0], $res);
        });

        // files
        Router::get('/files', function (Request $req, Response $res) {
            (new FileController())->getFiles($res);
        });

        Router::get('/files/([0-9]*)', function (Request $req, Response $res) {
            (new FileController())->getFile((int) $req->params[0], $res);
        });

        Router::post('/files/upload', function (Request $req, Response $res) {
            $bearer = explode(" ", getallheaders()['Authorization']);
            Guard::check($bearer[1], $res);
            (new FileController())->save($res);
        });

        Router::put('/files/gallery/([0-9]*)', function (Request $req, Response $res) {
            $bearer = explode(" ", getallheaders()['Authorization']);
            Guard::check($bearer[1], $res);
            (new FileController())->addGallery((int) $req->params[0], $req->getJSON(), $res);
        });

        Router::delete('/files/([0-9]*)', function (Request $req, Response $res) {
            $bearer = explode(" ", getallheaders()['Authorization']);
            Guard::check($bearer[1], $res);
            (new FileController())->remove((int) $req->params[0], $res);
        });

        // authorization
        Router::post('/auth', function (Request $req, Response $res) {
            (new AuthController())->auth($req->getJSON(), $res);
        });

        // recovery
        Router::post('/recovery', function (Request $req, Response $res) {
            (new AuthController())->recovery($req->getJSON(), $res);
        });

        // email
        Router::post('/mail', function (Request $req, Response $res) {
            (new MailController())->send($req->getJSON(), $res);
        });

        //languages
        Router::get('/languages', function (Request $req, Response $res) {
            (new LanguagesController())->getLanguages($res);
        });

        Router::get('/languages/([a-z]*)', function (Request $req, Response $res) {
            (new LanguagesController())->getLanguage($req->params[0], $res);
        });

        Router::put('/languages', function (Request $req, Response $res) {
            $bearer = explode(" ", getallheaders()['Authorization']);
            Guard::check($bearer[1], $res);
            (new LanguagesController())->save($req->getJSON(), $res);
        });

        Logger::enableSystemLogs();
    }
}
