<?php
declare(strict_types = 1);

namespace ArualCms\Lib;

/**
 * Class Response
 * @package ArualCms\Lib
 */
class Response
{
    /** @var int  */
    private $status = 200;

    /**
     * @param int $code
     * @return $this
     */
    public function status(int $code): Response
    {
        $this->status = $code;
        return $this;
    }

    /**
     * @param array $data
     */
    public function toJSON(array $data = []): void
    {
        http_response_code($this->status);
        header('Access-Control-Allow-Origin: *');
        header('Access-Control-Allow-Headers: Content-Type, X-Requested-With, Authorization');
        header('Access-Control-Max-Age: 1000');
        header('Content-Type: application/json');
        echo json_encode($data);
        exit();
    }
}
