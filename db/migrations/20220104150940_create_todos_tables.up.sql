create table IF NOT EXISTS todos  (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    completed BOOLEAN DEFAULT FALSE
);