<?php
if (false === extension_loaded('amqp')) {
    die('AMQP extension must be installed and enabled.');
}

[$null, $host, $user, $password, $limit, $msize] = $argv;
$limit = $limit ?? 1000;
$msize = $msize ?? 5;

$messageString = generateRandomStringBytes($msize);

$connection = new AMQPConnection([
    'host' => $host,
    'login' => $user,
    'password' => $password,
]);

if ($connection->connect()) {
    echo 'RabbitMQ connected', PHP_EOL;
} else {
    die('Failed to connect to RabbitMQ');
}

echo sprintf('Message size: %d bytes.', $msize), PHP_EOL, '---------', PHP_EOL;

$channel = new AMQPChannel($connection);
$channel->qos(0, 100);

$timerStart = microtime(true);

$exchange = new AMQPExchange($channel);
$exchange->setName('php71-amqp');
$exchange->setType(AMQP_EX_TYPE_DIRECT);
$exchange->setFlags(AMQP_NOPARAM);
$exchange->setFlags(AMQP_DURABLE);


$exchange->declareExchange();

$queue = new AMQPQueue($channel);
$queue->setName('php71-amqp');
$queue->setFlags(AMQP_NOPARAM);
$queue->setFlags(AMQP_DURABLE);

$queue->declareQueue();
$queue->bind($exchange->getName(), $queue->getName());

for ($i = 0; $i < $limit; $i++) {
    $exchange->publish($messageString, $queue->getName());

    if ($i > 0 && 0 === $i % 1000) {
        echo "Processed {$i} of {$limit}" . PHP_EOL;
    }
}

$timerEnd = microtime(true);
$timeDiff = $timerEnd - $timerStart;

echo sprintf('Elapsed: %f seconds', $timeDiff), PHP_EOL;
echo sprintf('Throughput: %f / second', $limit / $timeDiff), PHP_EOL;

$channel->close();
$connection->disconnect();

function generateRandomStringBytes(int $n) {
    $letters = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ';

    $response = '';

    for ($i = 0; $i < $n; $i++) {
        $response .= $letters[random_int(0, strlen($letters))];
    }

    return $response;
}
