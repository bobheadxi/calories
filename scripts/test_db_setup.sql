CREATE TABLE IF NOT EXISTS users (
    user_id TEXT UNIQUE,
    max_cal INTEGER,
    timezone INTEGER,
    name TEXT
);

CREATE TABLE IF NOT EXISTS entries (
    fuser_id TEXT UNIQUE,
    time BIGINT,
    item TEXT,
    calories INTEGER
);