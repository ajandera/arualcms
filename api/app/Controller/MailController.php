<?php
declare(strict_types = 1);

namespace ArualCms\Controller;

use ArualCms\Lib\Response;
use ArualCms\Model\Settings;
use PHPMailer\PHPMailer\PHPMailer;
use PHPMailer\PHPMailer\SMTP;
use PHPMailer\PHPMailer\Exception;

/**
 * Class MailController
 * @package ArualCms\Controller
 */
class MailController
{
    public function __construct()
    {
        Settings::load();
    }

    /**
     * @param $email
     * @param Response $res
     */
    public function send($email, Response $res): void
    {
        $smtp = Settings::findByKey('smtp');
        $smtpUser = Settings::findByKey('smtp_user');
        $smtpPassword = Settings::findByKey('smtp_password');
        $smtpPort = Settings::findByKey('smtp_port');

        //Instantiation and passing `true` enables exceptions
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
            $mail->Subject = $email->subject;
            $mail->Body = $email->body;
            $mail->AltBody = strip_tags($email->body);

            $mail->send();
            $res->toJSON(['success' => true, 'message' => 'Message has been sent']);
        } catch (Exception $e) {
            $to = $email->email;
            $subject = $email->subject;
            $message = $email->body;
            $response = mail($to, $subject, $message);
            $res->status(200)->toJSON(['success' => true, 'message' => "Message send by mail function", 'res' => $response]);
        }
    }
}
