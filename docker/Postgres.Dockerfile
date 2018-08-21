FROM postgres:10.4
ADD ./docker/init.sql /docker-entrypoint-initdb.d/