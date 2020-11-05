CREATE TABLE IF NOT EXISTS "posts" (
  "id" serial PRIMARY KEY,
  "title" varchar NOT NULL,
  "descr" varchar  NOT NULl DEFAULT '',
  "body" text NOT NULL,
  "thumbnail" varchar  NOT NULl DEFAULT '',
  "image" varchar  NOT NULl DEFAULT '',
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "user_id" int NOT NULL DEFAULT 0,
  "cate_id" int NOT NULL DEFAULT 0
);
