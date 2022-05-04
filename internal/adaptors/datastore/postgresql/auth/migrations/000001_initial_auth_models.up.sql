-- Migration file created on the 1st May 2022

CREATE TYPE ROLES AS ENUM (
    'ADMIN',
    'COMMON'
    );

CREATE TYPE PERMISSIONS AS ENUM (
    'READ;WRITE',
    'READ'
    );

CREATE TABLE IF NOT EXISTS users
(
    row_id          BIGSERIAL   NOT NULL,
    user_id         uuid        NOT NULL,
    name            TEXT        NOT NULL,
    email           TEXT        NOT NULL,
    hashed_password TEXT        NOT NULL,
    role            ROLES       NOT NULL,
    permissions     PERMISSIONS NOT NULL,
    created_at      timestamptz NOT NULL DEFAULT now(),
    updated_at      timestamptz NOT NULL DEFAULT now(),

    CONSTRAINT user_id_pk PRIMARY KEY (user_id)
);

CREATE TABLE IF NOT EXISTS user_session
(
    user_id      uuid        NOT NULL,
    access_token TEXT        NOT NULL,
    role         ROLES       NOT NULL,
    permissions  PERMISSIONS NOT NULL,
    created_at   timestamptz NOT NULL DEFAULT now(),
    updated_at   timestamptz NOT NULL DEFAULT now(),

    CONSTRAINT user_session_user_id_pk PRIMARY KEY (user_id),
    CONSTRAINT user_session_user_id_fk FOREIGN KEY (user_id) REFERENCES users (user_id)
);

CREATE TABLE IF NOT EXISTS recovering
(
    user_id       uuid        NOT NULL,
    recovery_link TEXT        NOT NULL,
    expiry_date   timestamptz NOT NULL,
    created_at    timestamptz NOT NULL DEFAULT now(),
    updated_at    timestamptz NOT NULL DEFAULT now(),

    CONSTRAINT recovering_user_id_pk PRIMARY KEY (user_id),
    CONSTRAINT recovering_user_id_fk FOREIGN KEY (user_id) REFERENCES users (user_id)
);
