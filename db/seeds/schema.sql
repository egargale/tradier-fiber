-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2022-01-02T08:09:34.038Z

CREATE TYPE "account_type" AS ENUM (
  'cash',
  'margin',
  'pdt'
);

CREATE TABLE "users" (
  "user_id" SERIAL PRIMARY KEY,
  "tradier_id" varchar NOT NULL,
  "name" varchar NOT NULL
);

CREATE TABLE "profiles" (
  "profile_id" SERIAL PRIMARY KEY,
  "user_id" int,
  "balances_id" int,
  "account_number" varchar NOT NULL,
  "classification" varchar NOT NULL,
  "date_created" timestamptz NOT NULL,
  "day_trader" boolean NOT NULL,
  "option_level" int NOT NULL,
  "status" varchar NOT NULL,
  "type" Account_Type NOT NULL,
  "last_update_date" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "positions" (
  "position_id" SERIAL PRIMARY KEY,
  "profile_id" int,
  "cost_basis" int,
  "date_acquired" timestamptz,
  "quantity" int,
  "symbol" varchar,
  "strategy_id" bigint
);

CREATE TABLE "strategies" (
  "strategy_id" SERIAL PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "profiles" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "positions" ADD FOREIGN KEY ("profile_id") REFERENCES "profiles" ("profile_id");

ALTER TABLE "positions" ADD FOREIGN KEY ("strategy_id") REFERENCES "strategies" ("strategy_id");
