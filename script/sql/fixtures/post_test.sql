
INSERT INTO threads
  (title, default_name, post_num, created_at, updated_at)
VALUES
  ("foo", "", 1, 100, 100),
  ("bar", "bar", 2, 200, 300);

INSERT INTO posts
  (thread_id, name, content, created_at)
VALUES
  (1, "foo", "foo", 100),
  (2, "", "foofoo", 200),
  (2, "barbar", "barbar", 300)
