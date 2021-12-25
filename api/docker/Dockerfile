FROM php:8.0-apache

MAINTAINER shopyCRM <jandera@shopycrm.com>

# Install libraries
RUN apt-get update && \
    apt-get install -y \
    openssl libssl-dev libzip-dev libfreetype6-dev libjpeg62-turbo-dev libpng-dev git

RUN docker-php-ext-install bcmath zip

RUN docker-php-ext-configure gd && docker-php-ext-install gd

RUN docker-php-ext-configure zip && docker-php-ext-install zip

RUN pecl install mongodb && docker-php-ext-enable mongodb

COPY ./php.ini /usr/local/etc/php/

RUN curl -sS https://getcomposer.org/installer | php -- \
--install-dir=/usr/bin --filename=composer && chmod +x /usr/bin/composer

ADD ./apache.conf /etc/apache2/sites-available/
RUN ln -s /etc/apache2/sites-available/apache.conf /etc/apache2/sites-enabled/
RUN rm -rf /etc/apache2/sites-enabled/000-default.conf
RUN a2enmod rewrite

