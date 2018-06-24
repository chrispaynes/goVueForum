#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-SQL
  CREATE TABLE IF NOT EXISTS forum_category (
      id        SERIAL PRIMARY KEY,
      title     VARCHAR(50) NOT NULL UNIQUE
  );

  CREATE TABLE IF NOT EXISTS forum (
      id          SERIAL PRIMARY KEY,
      category    INTEGER REFERENCES forum_category(id) ON DELETE CASCADE,
      title       VARCHAR(50) NOT NULL UNIQUE,
      description VARCHAR(250) NOT NULL
  );
SQL