CREATE TABLE users (
    user_id TEXT UNIQUE,
    max_cal INTEGER,
    timezone INTEGER,
    name TEXT
);

CREATE TABLE entries (
    fuser_id TEXT UNIQUE,
    time BIGINT,
    item TEXT,
    calories INTEGER
);