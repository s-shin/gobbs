/*
## NOTE ##

* MySQL <=> SQLite
* VARCHAR(n) <=> TEXT
* DATETIME <=> NUMBER

The time is treated as unix time.
*/

DROP TABLE IF EXISTS threads;
CREATE TABLE threads (
  id            INTEGER       PRIMARY KEY AUTOINCREMENT,
  title         VARCHAR(100)  NOT NULL,
  default_name  VARCHAR(20),
  post_num      INTEGER       DEFAULT 1,
  created_at    DATETIME      NOT NULL,
  updated_at    DATETIME      NOT NULL
);
CREATE INDEX threads_i1 ON threads (updated_at);

DROP TABLE IF EXISTS posts;
CREATE TABLE posts (
  id            INTEGER       PRIMARY KEY AUTOINCREMENT,
  thread_id     INTEGER       NOT NULL,
  name          TEXT          DEFAULT NULL,
  content       TEXT          NOT NULL,
  created_at    DATETIME      NOT NULL
);
CREATE INDEX posts_i1 ON posts (thread_id);
