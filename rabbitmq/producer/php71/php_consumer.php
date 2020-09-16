<?php
/**
 * USAGE:
 * $ docker exec -it rabbitmq_clustered_php71-cli_1 php /app/php_consumer.php localhost admin Passw0rd 10000
 */
if (false === extension_loaded('amqp')) {
    die('AMQP extension must be installed and enabled.');
}

class Metric
{
    public static $i;
    public static $limit;

    public static function getLimit()
    {
        return self::$limit;
    }
}

[$null, $host, $user, $password, $limit] = $argv;
$limit = $limit ?? 1000;

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

$channel = new AMQPChannel($connection);

$timerStart = microtime(true);

$queue = new AMQPQueue($channel);
$queue->setName('php71-amqp');
$queue->setFlags(AMQP_DURABLE);
$queue->declareQueue();

$metric = new Metric();
$metric::$i = 1;
$metric::$limit = $limit;

try {
    $queue->consume('handleEnvelope');
} catch (\Exception $e) {
    // null
} finally {
    $timerEnd = microtime(true);
    $timeDiff = $timerEnd - $timerStart;

    echo sprintf('Elapsed: %f seconds', $timeDiff), PHP_EOL;
    echo sprintf('Read Throughput: %f / second', $limit / $timeDiff), PHP_EOL;
}


$queue->cancel();
$channel->close();
$connection->disconnect();

function handleEnvelope(AMQPEnvelope $envelope, AMQPQueue $queue)
{
    $limit = Metric::getLimit();
    $i = Metric::$i++;

    if ($i < $limit) {
        $queue->nack($envelope->getDeliveryTag());

        if ($i > 0 && 0 === $i % 1000) {
            echo "Processed {$i} of {$limit}" . PHP_EOL;
        }
    } else {
        throw new \Exception('Out of limit');
    }
}
