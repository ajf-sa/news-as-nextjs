CREATE TABLE IF NOT EXISTS "users" (
  "id" serial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar NOT NULL DEFAULT '',
  "create_at" timestamptz NOT NULL DEFAULT (now())
);