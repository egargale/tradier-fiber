CREATE TYPE "Account_Type" AS ENUM (
  'cash',
  'margin',
  'pdt'
);

CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "tradier_id" varchar NOT NULL,
  "name" varchar NOT NULL
);

CREATE TABLE "profiles" (
  "id" SERIAL PRIMARY KEY,
  "user_id" bigint,
  "balances_id" bigint,
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
  "position_id" int PRIMARY KEY,
  "profile_id" bigint,
  "cost_basis" bigint,
  "date_acquired" timestamptz,
  "quantity" int,
  "symbol" varchar,
  "strategy_id" bigserial
);

CREATE TABLE "strategies" (
  "strategy_id" SERIAL PRIMARY KEY,
  "profile_id" bigint,
  "position_id" bigi
);

CREATE TABLE "balances" (
  "balances_id" bigserial PRIMARY KEY,
  "option_short_value" int,
  "total_equit" bigint,
  "account_number" bigint,
  "account_type" varchar,
  "close_pl" bigint,
  "current_requirement" bigint,
  "equity" bigint,
  "long_market_value" bigint,
  "market_value" bigint,
  "open_pl" bigint,
  "option_long_value" bigint,
  "option_requirement" int,
  "pending_orders_count" int,
  "short_market_value" bigint,
  "stock_long_value" bigint,
  "total_cash" bigint,
  "uncleared_funds" bigint,
  "pending_cash" bigint
);

CREATE TABLE "balance_margin" (
  "balances_type" varchar PRIMARY KEY,
  "fed_call" int,
  "maintenance_call" int,
  "option_buying_power" bigint,
  "stock_buying_power" bigint,
  "stock_short_value" int,
  "sweep" int
);

CREATE TABLE "balance_cash" (
  "balances_type" varchar PRIMARY KEY,
  "cash_available" bigint,
  "sweep" int,
  "unsettled_funds" bigint
);

CREATE TABLE "balance_pdt" (
  "balances_type" varchar PRIMARY KEY,
  "fed_call" int,
  "maintenance_call" int,
  "option_buying_power" bigint,
  "stock_buying_power" bigint,
  "stock_short_value" int
);

ALTER TABLE "profiles" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "profiles" ADD FOREIGN KEY ("balances_id") REFERENCES "balances" ("balances_id");

ALTER TABLE "positions" ADD FOREIGN KEY ("profile_id") REFERENCES "profiles" ("id");

ALTER TABLE "strategies" ADD FOREIGN KEY ("profile_id") REFERENCES "profiles" ("id");

ALTER TABLE "positions" ADD FOREIGN KEY ("position_id") REFERENCES "strategies" ("position_id");

ALTER TABLE "profiles" ADD FOREIGN KEY ("type") REFERENCES "balance_margin" ("balances_type");

ALTER TABLE "profiles" ADD FOREIGN KEY ("type") REFERENCES "balance_pdt" ("balances_type");

ALTER TABLE "profiles" ADD FOREIGN KEY ("type") REFERENCES "balance_cash" ("balances_type");
