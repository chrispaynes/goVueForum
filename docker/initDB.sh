#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-SQL
  SET TIME ZONE 'UTC';

  CREATE TABLE IF NOT EXISTS forum_category (
      forum_category_id       SERIAL PRIMARY KEY,
      title                   VARCHAR(50) NOT NULL UNIQUE
  );

  INSERT INTO forum_category(title)
  VALUES  ('Test Category 1'),
          ('Test Category 2'),
          ('Test Category 3');

  CREATE TABLE IF NOT EXISTS forum (
      forum_id               SERIAL PRIMARY KEY,
      forum_category_id      INTEGER REFERENCES forum_category(forum_category_id) ON DELETE CASCADE,
      title                  VARCHAR(50) NOT NULL UNIQUE,
      description            VARCHAR(250) NOT NULL
  );

  INSERT INTO forum(forum_category_id, title, description)
  VALUES  (1, 'Test Forum 1 - Category 1', 'Test Description'),
          (1, 'Test Forum 2 - Category 1', 'Test Description'),
          (2, 'Test Forum 1 - Category 2', 'Test Description'),
          (2, 'Test Forum 2 - Category 2', 'Test Description'),
          (3, 'Test Forum 1 - Category 3', 'Test Description');

  CREATE TABLE IF NOT EXISTS user_account (
      user_account_id     SERIAL PRIMARY KEY,
      first_name          VARCHAR(25) NOT NULL,
      last_name           VARCHAR(25) NOT NULL,
      post_count          SERIAL NOT NULL DEFAULT 0,
      email               VARCHAR(50) NOT NULL UNIQUE,
      avatar_url          VARCHAR(50),
      location            VARCHAR(50),
      username            VARCHAR(50) NOT NULL UNIQUE,
      last_login          TIMESTAMPTZ
  );

  INSERT INTO user_account (first_name, last_name, email, username, avatar_url, location)
  VALUES  ('test', 'author', 'test_author@author.com', 'TestAuthor1', 'foo.jpg', 'Budapest'),
          ('test', 'author2', 'test_author2@author.com', 'TestAuthor2', 'bar.jpg', 'Kiev');

  CREATE TABLE IF NOT EXISTS thread (
    thread_id           SERIAL PRIMARY KEY,
    forum_id            INTEGER REFERENCES forum(forum_id) ON DELETE CASCADE,
    title               VARCHAR(100) NOT NULL,
    author_id           INTEGER REFERENCES user_account(user_account_id) ON DELETE CASCADE,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_reply_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
  );

  INSERT INTO thread(forum_id, title, author_id)
  VALUES  (1, 'Test Thread - Forum 1 - Author 1', 1),
          (1, 'Test Thread - Forum 1 - Author 2', 2),
          (2, 'Test Thread - Forum 2 - Author 1', 1),
          (2, 'Test Thread - Forum 2 - Author 2', 2);


  CREATE TABLE IF NOT EXISTS post_body (
      post_body_id      SERIAL PRIMARY KEY,
      body              TEXT
  );

  INSERT INTO post_body(body)
  VALUES  ('<p>Post Body 1</p>'),
          ('<p>Post Body 2</p>'),
          ('<p>Post Body 3</p>'),
          ('<p>Post Body 4</p>'),
          ('<p>Post Body 5</p>');

  CREATE TABLE IF NOT EXISTS post (
      post_id             SERIAL PRIMARY KEY,
      author_id           INTEGER REFERENCES user_account(user_account_id) ON DELETE CASCADE,
      thread_id           INTEGER REFERENCES thread(thread_id) ON DELETE CASCADE,
      title               VARCHAR(75),
      post_body_id        INTEGER REFERENCES post_body(post_body_id) ON DELETE CASCADE,
      created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
      last_updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW()
  );

  INSERT INTO post(author_id, thread_id, title, post_body_id)
  VALUES  (1, 1, 'Test Article - Author 1 - Thread 1', 1),
          (2, 1, 'Test Article - Author 2 - Thread 1', 2),
          (1, 2, 'Test Article - Author 1 - Thread 2', 3),
          (2, 2, 'Test Article - Author 2 - Thread 2', 4);

  CREATE VIEW posts_v AS
    SELECT p.post_id, ua.username, p.author_id, p.title, p.created_at, p.last_updated_at, pb.body
    FROM post AS p
    LEFT JOIN post_body AS pb
    ON p.post_body_id = pb.post_body_id
    LEFT JOIN user_account AS ua
    ON p.author_id = ua.user_account_id;

  CREATE FUNCTION update_last_updated_column()
    RETURNS TRIGGER AS $$
    BEGIN
        NEW.last_updated_at = now();
        RETURN NEW;
    END;
    $$ language 'plpgsql';

  CREATE TRIGGER update_post_modtime
    BEFORE UPDATE ON post
    FOR EACH ROW EXECUTE PROCEDURE update_last_updated_column();

  CREATE or REPLACE FUNCTION get_user(user_id integer)
    RETURNS TABLE ( avatar_url varchar,
                    email varchar,
                    first_name varchar,
                    user_account_id integer,
                    last_login timestamptz,
                    last_name varchar,
                    location varchar,
                    post_count integer,
                    username varchar
                  ) AS $function$
    SELECT
        avatar_url,
        email,
        first_name,
        user_account_id,
        last_login,
        last_name,
        location,
        post_count,
        username
        FROM user_account
        WHERE user_account_id = $1;
    $function$ LANGUAGE SQL;

  CREATE VIEW users_v AS
    SELECT
        avatar_url,
        email,
        first_name,
        user_account_id,
        last_login,
        last_name,
        location,
        post_count,
        username
        FROM user_account;

  CREATE or REPLACE FUNCTION increment_post_count() RETURNS TRIGGER AS
    $$
    BEGIN
        UPDATE user_account
        SET post_count = post_count + 1
        WHERE user_account_id = new.author_id;
        RETURN new;
    end;
    $$ LANGUAGE plpgsql;

  CREATE TRIGGER increment_user_post_count
    AFTER INSERT ON post
    FOR EACH ROW EXECUTE PROCEDURE increment_post_count();
SQL