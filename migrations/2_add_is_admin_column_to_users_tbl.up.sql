alter table users
    ADD COLUMN is_admin BOOLEAN NOT NULL DEFAULT FALSE;