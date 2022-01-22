<?php
declare(strict_types = 1);

namespace ArualCms\Controller;

use ArualCms\Lib\Config;
use ArualCms\Lib\Response;
use ArualCms\Model\Users;
use ArualCms\Model\Setting;
use Firebase\JWT\JWT;

/**
 * Class AuthController
 * @package ArualCms\Controller
 */
class AuthController
{
    use Users;
    use Setting;

    /**
     * @param $req
     * @param Response $res
     * @throws \Exception
     */
    public function auth($req, Response $res): void
    {
        $user = $this->findByUsername($req->username)[0];
        if ($user && password_verify($req->password, $user['password'])) {
            $token = [
                "iat" => time(),
                "exp" => Config::get("EXPIRATION_TIME"),
                "iss" => Config::get("ISSUER", "http://localhost:8080"),
                "data" => [
                    "id" => $user['id'],
                    "username" => $user['username']
                ]
            ];
            $data["id"] = $user['id'];
            $data["username"] = $user['username'];
            $data['jwt'] = JWT::encode($token, Config::get("KEY", "fsdfsdfsdSDF34"));
            $data['success'] = true;
        } else {
            $data['success'] = false;
            $data['message'] = "Bad credentials.";
        }

        $res->toJSON($data);
    }

    /**
     * @param $req
     * @param Response $res
     * @throws \MongoCursorException
     * @throws \Exception
     */
    public function recovery($req, Response $res): void {
        $user = $this->findByUsername($req->username);
            if ($user["recovery"] === $req->token) {
            $user['password'] = password_hash($req->password, PASSWORD_BCRYPT);
            $this->edit($user);
            $data['success'] = true;
            $data['message'] = "New password was set.";
        } else {
            $data['success'] = false;
            $data['message'] = "Token is invalid";
        }
        $res->toJSON($data);
    }
    
    /**
     * @param $req
     * @param Response $res
     * @throws \MongoCursorException
     * @throws \Exception
     */
    public function sendRecovery($req, Response $res): void {
        $user = $this->findByUsername($req->username);
        $token = bin2hex(random_bytes(20));
        $user["recovery"] = $token;
        $this->edit($user);
        
        // send to user
        $smtp = $this->findByKey('smtp');
        $smtpUser = $this->findByKey('smtp_user');
        $smtpPassword = $this->findByKey('smtp_password');
        $smtpPort = $this->findByKey('smtp_port');

        $link = $req->domain + "/" + $req->username + "/" + $token;
      
        $body = "For recovery your password use this link: <a href='{$link}'>{$link}</a>";

        //Instantiation and passing 'true' enables exceptions
        $mail = new PHPMailer(true);

        try {
            //Server settings
            $mail->SMTPDebug = SMTP::DEBUG_OFF;
            $mail->isSMTP();
            $mail->Host = $smtp->value;
            $mail->SMTPAuth = true;
            $mail->Username = $smtpUser->value;
            $mail->Password = $smtpPassword->value;
            $mail->SMTPSecure = PHPMailer::ENCRYPTION_STARTTLS;
            $mail->Port = $smtpPort->value;

            //Recipients
            $mail->setFrom('noreply@arualcms.com', 'ArualCMS');
            $mail->addAddress($email->email);

            //Content
            $mail->isHTML(true);
            $mail->Subject = "Recovery Password";
            $mail->Body = $body;
            $mail->AltBody = strip_tags($body);

            $mail->send();
        } catch (Exception $e) {
            $to = $email->email;
            $subject = "Recovery Password";
            $message = $body;
            mail($to, $subject, $message);
        }
        
        $data['success'] = true;
        $data['message'] = "Recovery email was send.";
        $res->toJSON($data);
    }
}
