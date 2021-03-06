<?php

return [
    'LOG_PATH' => __DIR__ . '/logs',
    'DB_PATH' => __DIR__ . '/database/',
    'STORAGE_PATH' => __DIR__ . '/storage/',
    // variables used for jwt
    "KEY" => "example_key",
    "EXPIRATION_TIME" => time() + (60 * 60), // valid for 1 hour
    "ISSUER" => "http://localhost:8080"
];
