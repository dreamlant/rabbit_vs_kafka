FROM php:7.1.3

ENV LIBRDKAFKA_VERSION=0.9.4

# https://pecl.php.net/package/rdkafka --- look here the stable version
ENV PHPRDKAFKA_VERSION=3.1.2
ENV PHP_VERSION=7.1.3



RUN apt-get update && apt-get -y install \
	build-essential \
	libc6-dev \
	zlib1g-dev \
	python-dev \
	wget \
	sudo \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /tmp

RUN useradd -m docker \
 && echo "docker:docker" | chpasswd \
 && adduser docker sudo

RUN wget https://github.com/edenhill/librdkafka/archive/v${LIBRDKAFKA_VERSION}.tar.gz \
 && tar -xvf v${LIBRDKAFKA_VERSION}.tar.gz \
 && cd librdkafka-${LIBRDKAFKA_VERSION} \
 && ./configure \
 && make \
 && make install \
 && make clean \
 && ./configure --clean \
 && ldconfig \
 && cd /tmp \
 && rm -rf v${LIBRDKAFKA_VERSION}.tar.gz librdkafka-${LIBRDKAFKA_VERSION}

RUN wget https://github.com/arnaud-lb/php-rdkafka/archive/${PHPRDKAFKA_VERSION}.tar.gz \
 && tar -xvf ${PHPRDKAFKA_VERSION}.tar.gz \
 && cd php-rdkafka-${PHPRDKAFKA_VERSION} \
 && phpize \
 && ./configure \
 && make all -j 5 \
 && sudo make install \
 && echo "; RdKafka Extension\nextension=rdkafka.so" >/usr/local/etc/php/conf.d/rdkafka.ini \
 && cd /tmp \
 && rm -rf ${PHPRDKAFKA_VERSION}.tar.gz php-rdkafka-${PHPRDKAFKA_VERSION}

RUN wget https://github.com/php/php-src/archive/php-${PHP_VERSION}.tar.gz \
 && tar -xvf php-${PHP_VERSION}.tar.gz \
 && cd php-src-php-${PHP_VERSION}/ext/pcntl \
 && phpize && ./configure && make install \
 && echo "; Process Control support Extension\nextension=pcntl.so" >/usr/local/etc/php/conf.d/pcntl.ini \
 && cd /tmp \
 && rm -rf php-${PHP_VERSION}.tar.gz php-src-php-${PHP_VERSION}

CMD ["tail", "-f", "/dev/null"]
