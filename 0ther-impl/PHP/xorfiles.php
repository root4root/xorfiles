#!/usr/bin/env php
<?php

const BUF = 524288;

/// Start processing CLI params ---->
if (count($argv) < 2 || count($argv) > 4) {
    echo "Usage:\n\$ xorfiles.php {first-file | stdin} second-file [output-file]\n\n";
    exit(22);
}

switch (count($argv)) {
    case 2:
        $firstFile = fopen('php://stdin', 'r');
        $secondFile = fopen($argv[1], 'r');
        $outputFile = fopen('php://stdout', 'w');
        break;
    case 3:
        if (file_exists($argv[2]) === true) {
            $firstFile = fopen($argv[1], 'r');
            $secondFile = fopen($argv[2], 'r');
            $outputFile = fopen('php://stdout', 'w');
        } else {
            $firstFile = fopen('php://stdin', 'r');
            $secondFile = fopen($argv[1], 'r');
            $outputFile = fopen($argv[2], 'w');
        }
        break;
    default:
            $firstFile = fopen($argv[1], 'r');
            $secondFile = fopen($argv[2], 'r');
            $outputFile = fopen($argv[3], 'w');
}

if ($firstFile === false || $secondFile === false || $outputFile === false) {
    echo "\nERROR! Can't open some specified files. Please check.\n\n";
    exit(2);
} /// <---- End processing CLI params

$firstUnit = '';
$secondUnit = '';
$result = '';

while(!feof($firstFile) && !feof($secondFile)) {

    $firstUnit = fread($firstFile, BUF);

    if (strlen($firstUnit) > 0) {
        $secondUnit = fread($secondFile, strlen($firstUnit));
        $result = $firstUnit ^ $secondUnit;
        if (fwrite($outputFile, $result, BUF) === false) {
            fwrite(STDERR, "\nERROR! Can't write data to the output-file (prehaps no space left). Exiting.\n");
            exit(5);
        }
    }
}

fclose($firstFile);
fclose($secondFile);
fclose($outputFile);
