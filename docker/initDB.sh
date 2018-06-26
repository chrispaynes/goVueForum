#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-SQL
  SET TIME ZONE 'UTC';

  CREATE TABLE IF NOT EXISTS forum_category (
      id          SERIAL PRIMARY KEY,
      title       VARCHAR(50) NOT NULL UNIQUE
  );

  CREATE TABLE IF NOT EXISTS forum (
      id          SERIAL PRIMARY KEY,
      category    INTEGER REFERENCES forum_category(id) ON DELETE CASCADE,
      title       VARCHAR(50) NOT NULL UNIQUE,
      description VARCHAR(250) NOT NULL
  );

  CREATE TABLE IF NOT EXISTS user_account (
      id          SERIAL PRIMARY KEY,
      first_name  VARCHAR(25) NOT NULL,
      last_name   VARCHAR(25) NOT NULL,
      post_count  INTEGER NOT NULL DEFAULT 0,
      email       VARCHAR(50) NOT NULL UNIQUE,
      avatar_url  VARCHAR(50),
      location    VARCHAR(50),
      username    VARCHAR(50) NOT NULL UNIQUE
  );

  INSERT INTO user_account (first_name,last_name, email, username, avatar_url, location)
  VALUES  ('test', 'author', 'test_author@author.com', 'TestAuthor1', 'foo.jpg', 'Budapest'),
          ('test', 'author2', 'test_author2@author.com', 'TestAuthor2', 'bar.jpg', 'Kiev');

  CREATE TABLE IF NOT EXISTS post_body (
      id          SERIAL PRIMARY KEY,
      body        TEXT
  );

  INSERT INTO post_body(body)
  VALUES  ('<p>Hello World 1</p>'),
          ('<p>Hello World 2</p>');

  CREATE TABLE IF NOT EXISTS post (
      id            SERIAL PRIMARY KEY,
      author_id     INTEGER REFERENCES user_account(id) ON DELETE CASCADE,
      title         VARCHAR(25) NOT NULL,
      body          INTEGER REFERENCES post_body(id) ON DELETE CASCADE,
      created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
      last_updated  TIMESTAMPTZ NOT NULL DEFAULT NOW()
  );

  INSERT INTO post(author_id, title, body)
  VALUES  (1, 'Test Article 1', 1),
          (2, 'Test Article 2', 2);

  CREATE VIEW posts_v AS
    SELECT p.id, ua.username AS author, p.title, p.created_at, p.last_updated, pb.body
    FROM post AS p
    LEFT JOIN post_body AS pb
    ON p.body = pb.id
    LEFT JOIN user_account AS ua
    ON p.author_id = ua.id;

SQL