SET TIME ZONE 'UTC';

CREATE TABLE IF NOT EXISTS forum_category (
    forum_category_id                     serial NOT NULL,
    title                                 varchar(50) NOT NULL,
    CONSTRAINT forum_category_pkey        PRIMARY KEY (forum_category_id),
    CONSTRAINT forum_category_title_key   UNIQUE (title)
);

INSERT INTO forum_category (title)
  VALUES  ('Test Category 1'),
          ('Test Category 2'),
          ('Test Category 3');

CREATE TABLE IF NOT EXISTS forum (
    forum_id                                  serial NOT NULL,
    forum_category_id                         int4 NULL,
    title                                     varchar(50) NOT NULL,
    description                               varchar(250) NOT NULL,
    CONSTRAINT forum_pkey                     PRIMARY KEY (forum_id),
    CONSTRAINT forum_title_key UNIQUE         (title),
    CONSTRAINT forum_forum_category_id_fkey   FOREIGN KEY (forum_category_id) REFERENCES forum_category (forum_category_id) ON DELETE CASCADE
);

INSERT INTO forum (forum_category_id, title, description)
  VALUES  (1, 'Test Forum 1 - Category 1', 'Test Description'),
          (1, 'Test Forum 2 - Category 1', 'Test Description'),
          (2, 'Test Forum 1 - Category 2', 'Test Description'),
          (2, 'Test Forum 2 - Category 2', 'Test Description'),
          (3, 'Test Forum 1 - Category 3', 'Test Description');

-- ON CONFLICT ON CONSTRAINT customers_name_key DO NOTHING;
CREATE TABLE IF NOT EXISTS user_account (
    user_account_id                       serial NOT NULL,
    first_name                            varchar(25) NOT NULL,
    last_name                             varchar(25) NOT NULL,
    post_count                            int4 DEFAULT 0,
    email                                 varchar(50) NOT NULL,
    avatar_url                            varchar(50) NULL,
    "location"                            varchar(50) NULL,
    username                              varchar(50) NOT NULL,
    last_login                            timestamptz NULL,
    CONSTRAINT user_account_email_key     UNIQUE (email),
    CONSTRAINT user_account_pkey          PRIMARY KEY (user_account_id),
    CONSTRAINT user_account_username_key  UNIQUE (username)
);

INSERT INTO user_account (first_name, last_name, email, username, avatar_url, LOCATION)
  VALUES  ('test', 'author', 'test_author@author.com', 'TestAuthor1', 'foo.jpg', 'Budapest'),
          ('test', 'author2', 'test_author2@author.com', 'TestAuthor2', 'bar.jpg', 'Kiev');

CREATE TABLE IF NOT EXISTS thread (
    thread_id                         serial NOT NULL,
    forum_id                          int4 NULL,
    title                             varchar(100) NOT NULL,
    author_id                         int4 NULL,
    created_at                        timestamptz NOT NULL DEFAULT now(),
    last_reply_at                     timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT thread_pkey            PRIMARY KEY (thread_id),
    CONSTRAINT thread_author_id_fkey  FOREIGN KEY (author_id) REFERENCES user_account (user_account_id) ON DELETE CASCADE,
    CONSTRAINT thread_forum_id_fkey   FOREIGN KEY (forum_id) REFERENCES forum (forum_id) ON DELETE CASCADE
);

INSERT INTO thread (forum_id, title, author_id)
  VALUES  (1, 'Test Thread - Forum 1 - Author 1', 1),
          (1, 'Test Thread - Forum 1 - Author 2', 2),
          (2, 'Test Thread - Forum 2 - Author 1', 1),
          (2, 'Test Thread - Forum 2 - Author 2', 2);

CREATE TABLE IF NOT EXISTS post_body (
    post_body_id                 serial NOT NULL,
    body                        text NULL,
    CONSTRAINT post_body_pkey   PRIMARY KEY (post_body_id)
);

INSERT INTO post_body (body)
  VALUES  ('<p>Post Body 1</p>'),
          ('<p>Post Body 2</p>'),
          ('<p>Post Body 3</p>'),
          ('<p>Post Body 4</p>'),
          ('<p>Post Body 5</p>');

CREATE TABLE IF NOT EXISTS post (
    post_id                             serial NOT NULL,
    author_id                           int4 NULL,
    thread_id                           int4 NULL,
    title                               varchar(75) NULL,
    post_body_id                        int4 NULL,
    created_at                          timestamptz NOT NULL DEFAULT now(),
    last_updated_at                     timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT post_pkey                PRIMARY KEY (post_id),
    CONSTRAINT post_author_id_fkey      FOREIGN KEY (author_id) REFERENCES user_account (user_account_id) ON DELETE CASCADE,
    CONSTRAINT post_post_body_id_fkey   FOREIGN KEY (post_body_id) REFERENCES post_body (post_body_id) ON DELETE CASCADE,
    CONSTRAINT post_thread_id_fkey      FOREIGN KEY (thread_id) REFERENCES thread (thread_id) ON DELETE CASCADE
);

INSERT INTO post (author_id, thread_id, title, post_body_id)
  VALUES
    (1, 1, 'Test Article - Author 1 - Thread 1', 1),
    (2, 1, 'Test Article - Author 2 - Thread 1', 2),
    (1, 2, 'Test Article - Author 1 - Thread 2', 3),
    (2, 2, 'Test Article - Author 2 - Thread 2', 4);

CREATE OR REPLACE VIEW posts_v AS
SELECT
  p.post_id,
  ua.username,
  p.author_id,
  p.title,
  p.created_at,
  p.last_updated_at,
  pb.body
FROM post p
  LEFT JOIN post_body pb ON p.post_body_id = pb.post_body_id
  LEFT JOIN user_account ua ON p.author_id = ua.user_account_id;

CREATE OR REPLACE FUNCTION update_last_updated_column()
  RETURNS TRIGGER AS $$
  BEGIN
      NEW.last_updated_at = now();
      RETURN NEW;
  END;
  $$ LANGUAGE plpgsql;

CREATE TRIGGER update_post_modtime BEFORE
UPDATE
  ON post FOR EACH ROW EXECUTE PROCEDURE update_last_updated_column ();

CREATE OR REPLACE FUNCTION get_user (user_id INTEGER)
  RETURNS TABLE (avatar_url VARCHAR, email VARCHAR, first_name VARCHAR, user_account_id INTEGER, last_login TIMESTAMPTZ, last_name VARCHAR, LOCATION VARCHAR, post_count INTEGER, username VARCHAR) AS
$function$
SELECT
  avatar_url,
  email,
  first_name,
  user_account_id,
  last_login,
  last_name,
  LOCATION,
  post_count,
  username
FROM
  user_account
WHERE
  user_account_id = $1;
$function$
LANGUAGE SQL;

CREATE VIEW users_v AS
SELECT
  avatar_url,
  email,
  first_name,
  user_account_id,
  last_login,
  last_name,
  LOCATION,
  post_count,
  username
FROM
  user_account;

CREATE OR REPLACE FUNCTION increment_post_count()
  RETURNS TRIGGER AS
$function$
BEGIN
  UPDATE
    user_account
  SET
    post_count = post_count + 1
  WHERE
    user_account_id = new.author_id;
  RETURN new;
END;
$function$
LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS increment_user_post_count ON "post";
CREATE TRIGGER increment_user_post_count AFTER INSERT ON post FOR EACH ROW EXECUTE PROCEDURE increment_post_count ();

CREATE OR REPLACE FUNCTION create_user (email_arg character varying, f_name_arg character varying, l_name_arg character varying, location_arg character varying, username_arg character varying)
  RETURNS TABLE (avatar_url character varying, email character varying, first_name character varying, user_account_id integer, last_login timestamp WITH time zone, last_name character varying, LOCATION character varying, post_count integer, username character varying)
    LANGUAGE sql
AS $function$ WITH inserted AS (INSERT INTO user_account (email, first_name, last_name, LOCATION, username)
  VALUES ($1, $2, $3, $4, $5)
RETURNING
  *)
  SELECT
    avatar_url, email, first_name, user_account_id, last_login, last_name, LOCATION, post_count, username
  FROM
    inserted;
$function$
