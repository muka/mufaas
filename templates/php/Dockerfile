FROM composer

COPY . /usr/src/myapp
WORKDIR /usr/src/myapp

RUN composer install

CMD [ "php", "./index.php" ]
