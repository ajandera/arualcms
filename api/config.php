<?php

return [
    'MONGO_DSN' => getenv('mongo_dsn'),
    'MONGO_DBNAME' => getenv('mongo_dbname'),
    'LOG_PATH' => __DIR__ . '/logs',
    'KEY' => 'example_key',
    'EXPIRATION_TIME' => time() + (60 * 60), // valid for 1 hour
    'ISSUER' => 'http://localhost:8080'
];
