CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "pwd" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TYPE Currency AS ENUM ('USD', 'INR', 'CASHBACK');

CREATE TABLE "wallets" (
  "id" bigserial PRIMARY KEY,
  "owner" bigint NOT NULL,
  "currency" Currency NOT NULL,
  "balance" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "wallets" ADD FOREIGN KEY ("owner") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("from_account_id") REFERENCES "wallets" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("to_account_id") REFERENCES "wallets" ("id");

CREATE INDEX ON "wallets" ("owner");

CREATE INDEX ON "wallets" ("owner", "currency");

CREATE INDEX ON "transactions" ("from_account_id");

CREATE INDEX ON "transactions" ("to_account_id");

CREATE INDEX ON "transactions" ("from_account_id", "to_account_id");
