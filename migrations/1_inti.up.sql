CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY, -- ИСПРАВЛЕНО ТУТ: теперь он автоинкрементный
    email TEXT NOT NULL UNIQUE,
    pash_hash BYTEA NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_email ON users (email);

CREATE TABLE IF NOT EXISTS apps
(
    id INTEGER PRIMARY KEY ,
    name TEXT not null unique ,
    secret text not null UNIQUE
);