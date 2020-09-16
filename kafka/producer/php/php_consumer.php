<?php

if (false === extension_loaded('rdkafka')) {
    die('Required extension "rdkafka".');
} else {
    echo 'Extension installed.', PHP_EOL;
}

[$cmd, $host, $topic, $limit, $msize] = $argv;

if (null === $host) {
    $host = '192.168.99.101:909,192.168.99.101:9093,192.168.99.101:9094';
}

$limit = $limit ?? 50000;
$topic = $topic ?? 'topic_bench_php';
$msize = $msize ?? 1024;

$conf = new RdKafka\Conf();
$rk = new RdKafka\Consumer($conf);
$rk->addBrokers($host);

$topic = $rk->newTopic($topic);
$topic->consumeStart(0, RD_KAFKA_OFFSET_BEGINNING);

$timerStart = microtime(true);

while (true) {
    // The first argument is the partition (again).
    // The second argument is the timeout.
    $msg = $topic->consume(0, 3000);
    if (null === $msg) {
        continue;
    } elseif ($msg->err) {
        echo $msg->errstr(), "\n";
        break;
    } else {
        echo $msg->payload, "\n";
    }
}

$timerEnd = microtime(true);
$timeDiff = $timerEnd - $timerStart;

echo sprintf('Elapsed: %f seconds', $timeDiff), PHP_EOL;
echo sprintf('Read Throuphut %f / second', $limit / $timeDiff), PHP_EOL;
