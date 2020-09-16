<?php

if (false === extension_loaded('rdkafka')) {
    die('Required extension "rdkafka".');
} else {
    echo 'Extension installed.', PHP_EOL;
}

[$cmd, $host, $topic, $limit, $msize] = $argv;

if (null === $host) {
    $host = '192.168.99.101:9092,192.168.99.101:9093,192.168.99.101:9094';
}

$limit = $limit ?? 50000;
$topic = $topic ?? 'topic_bench_php';
$msize = $msize ?? 1024;

$conf = new RdKafka\Conf();
//$conf->set('log_level', LOG_DEBUG);
//$conf->set('debug', 'all');
$rk = new RdKafka\Producer($conf);
$rk->addBrokers($host);

$topic = $rk->newTopic($topic);

$messageString = generateRandomStringBytes($msize);
echo sprintf('Message size: %d bytes.', $msize), PHP_EOL, '---------', PHP_EOL;

$timerStart = microtime(true);

for ($i = 0; $i < $limit; $i++) {
    $topic->produce(RD_KAFKA_PARTITION_UA, 0, $messageString);
    $rk->poll(0);

    if ($i > 0 && 0 === $i % 1000) {
        echo "Processed {$i} of {$limit}" . PHP_EOL;
    }
}

while ($rk->getOutQLen() > 0) {
    $rk->poll(50);
}
$timerEnd = microtime(true);
$timeDiff = $timerEnd - $timerStart;

echo sprintf('Elapsed: %f seconds', $timeDiff), PHP_EOL;
echo sprintf('Throuphut: %f / second', $limit / $timeDiff), PHP_EOL;

function generateRandomStringBytes(int $n)
{
    $letters = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ';

    $response = '';
    for ($i = 0; $i < $n; $i++) {
        $response .= $letters[random_int(0, strlen($letters))];
    }

    return $response;
}
