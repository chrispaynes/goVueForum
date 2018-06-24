FROM postgres:10.4
ADD ./docker/initDB.sh /docker-entrypoint-initdb.d/