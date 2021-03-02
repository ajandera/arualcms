<?php

namespace ArualCms;

use ArualCms\Lib\Logger;

class App
{
    public static function run()
    {
        Logger::enableSystemLogs();
    }
}
