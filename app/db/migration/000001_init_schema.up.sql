CREATE TABLE IF NOT EXISTS "users" (
  "id" serial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "is_active" boolean NOT NULL DEFAULT false,
  "is_staff"  boolean NOT NULL DEFAULT false,
  "is_admin"  boolean NOT NULL DEFAULT false
);
CREATE TABLE IF NOT EXISTS "pages" (
  "id" serial PRIMARY KEY,
  "title" varchar NOT NULL DEFAULT '',
  "descr" varchar NOT NULL DEFAULT '',
  "parent_id" int NOT NULL DEFAULT 0,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "is_active" boolean NOT NULL DEFAULT false
);

CREATE TABLE IF NOT EXISTS "posts" (
  "id" serial PRIMARY KEY,
  "title" varchar NOT NULL,
  "descr" varchar NOT NULL DEFAULT '',
  "body" text NOT NULL,
  "thumbnail" varchar NOT NULL DEFAULT '',
  "image" varchar NOT NULL DEFAULT '',
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "user_id" int NOT NULL DEFAULT 0,
  "page_id" int NOT NULL DEFAULT 0
);





CREATE TABLE IF NOT EXISTS "config" (
  "seo_title" varchar NOT NULL DEFAULT '',
  "seo_descript" varchar NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS "images" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL DEFAULT '',
  "dist" varchar NOT NULL DEFAULT '',
  "size" float NOT NULL DEFAULT 0,
  "width" float NOT NULL DEFAULT 0,
  "height" float NOT NULL DEFAULT 0
);

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "posts" ADD FOREIGN KEY ("page_id") REFERENCES "pages" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "pages" ADD FOREIGN KEY ("parent_id") REFERENCES "pages" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
