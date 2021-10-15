<?php

return [
    'MONGO_DSN' => 'mongodb://root:3FSFSDFDFrew@mongo:27017',
    'LOG_PATH' => __DIR__ . '/logs',
    'STORAGE_PATH' => __DIR__ . '/www/storage/',
    'STORAGE_FRONT' => '',
    'KEY' => 'example_key',
    'EXPIRATION_TIME' => time() + (60 * 60), // valid for 1 hour
    'ISSUER' => 'http://localhost:8080'
];
