<?php
declare(strict_types = 1);

namespace ArualCms\Lib;

/**
 * Class Request
 * @package ArualCms\Lib
 */
class Request
{
    /** @var array|mixed  */
    public mixed $params;

    /** @var string  */
    public string $reqMethod;

    /** @var string  */
    public string $contentType;

    public function __construct(array $params = [])
    {
        $this->params = $params;
        $this->reqMethod = trim($_SERVER['REQUEST_METHOD']);
        $this->contentType = !empty($_SERVER["CONTENT_TYPE"]) ? trim($_SERVER["CONTENT_TYPE"]) : '';
    }

    /**
     * @return array
     */
    public function getBody(): array
    {
        if ($this->reqMethod !== 'POST') {
            return [];
        }

        $body = [];
        foreach ($_POST as $key => $value) {
            $body[$key] = filter_input(INPUT_POST, $key, FILTER_SANITIZE_SPECIAL_CHARS);
        }

        return $body;
    }

    /**
     * @return array|mixed
     */
    public function getJSON()
    {
        header('Access-Control-Allow-Origin: *');
        header('Access-Control-Allow-Methods: GET, PUT, POST, DELETE, OPTIONS');
        header('Access-Control-Max-Age: 1000');
        header('Access-Control-Allow-Headers: Content-Type, X-Requested-With, Accept, Authorization');
        header('Access-Control-Allow-Credentials: true');
        header('Content-Type: application/json;charset=utf-8');

        if ($this->reqMethod !== 'POST' && $this->reqMethod !== 'PUT' ) {
            return '';
        }

        if (strcasecmp($this->contentType, 'application/json;charset=utf-8') !== 0) {
            return '';
        }

        // Receive the RAW post data.
        $content = trim(file_get_contents("php://input"));
        return json_decode($content);
    }
}
