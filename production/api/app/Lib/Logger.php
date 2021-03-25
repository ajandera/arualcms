<?php
declare(strict_types = 1);

namespace ArualCms\Lib;

use Monolog\ErrorHandler;
use Monolog\Handler\StreamHandler;

/**
 * Class Logger
 * @package ArualCms\Lib
 */
class Logger extends \Monolog\Logger
{
    /** @var array  */
    private static $loggers = [];

    /**
     * Logger constructor.
     * @param string $key
     * @param null $config
     */
    public function __construct($key = "app", $config = null)
    {
        parent::__construct($key);

        if (empty($config)) {
            $LOG_PATH = Config::get('LOG_PATH', __DIR__ . '/../../logs');
            $config = [
                'logFile' => "{$LOG_PATH}/{$key}.log",
                'logLevel' => \Monolog\Logger::DEBUG
            ];
        }

        try {
            $this->pushHandler(new StreamHandler($config['logFile'], $config['logLevel']));
        } catch (\Exception $e) {
            var_dump($e->getMessage());
        }
    }

    /**
     * @param string $key
     * @param null $config
     * @return Logger|mixed
     */
    public static function getInstance(string $key = "app", $config = null): array
    {
        if (empty(self::$loggers[$key])) {
            self::$loggers[$key] = new Logger($key, $config);
        }

        return self::$loggers[$key];
    }

    /**
     * Enable logged system.
     */
    public static function enableSystemLogs(): void
    {

        $LOG_PATH = Config::get('LOG_PATH', __DIR__ . '/../../logs');

        // Error Log
        self::$loggers['error'] = new Logger('errors');
        self::$loggers['error']->pushHandler(new StreamHandler("{$LOG_PATH}/errors.log"));
        ErrorHandler::register(self::$loggers['error']);

        // Request Log
        $data = [
            $_SERVER,
            $_REQUEST,
            trim(file_get_contents("php://input"))
        ];
        self::$loggers['request'] = new Logger('request');
        self::$loggers['request']->pushHandler(new StreamHandler("{$LOG_PATH}/request.log"));
        self::$loggers['request']->info("REQUEST", $data);
    }
}
