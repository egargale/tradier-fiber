create table IF NOT EXISTS accounts (
    account_id SERIAL PRIMARY KEY,
	id VARCHAR(11) UNIQUE NOT NULL,
	name VARCHAR(50) NOT NULL,
	account_number VARCHAR(50) NOT NULL,
	classification VARCHAR(15) NOT NULL,
	date_created DATE NOT NULL,
	day_trader boolean NOT NULL,
	option_level INT NOT NULL,
	status VARCHAR(6) NOT NULL,
	type VARCHAR(6) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now())
);
create table IF NOT EXISTS positions (
	position_id SERIAL PRIMARY KEY,
    profile_id VARCHAR(11),
	cost_basis DECIMAL(6,2),
	date_acquired DATE,
    id INT,
	quantity INT,
	symbol VARCHAR(50)
);

ALTER TABLE "positions" ADD FOREIGN KEY ("profile_id") REFERENCES "accounts" ("id");

CREATE INDEX IF NOT EXISTS idx_tradier_id ON accounts(id);
CREATE INDEX IF NOT EXISTS idx_position ON positions(symbol,date_acquired);

CREATE OR REPLACE VIEW positions_master AS
	SELECT accounts.id, accounts."name", accounts.classification, positions.symbol,  positions.quantity, positions.cost_basis, positions.date_acquired from accounts
	JOIN positions ON accounts.id = positions.profile_id;