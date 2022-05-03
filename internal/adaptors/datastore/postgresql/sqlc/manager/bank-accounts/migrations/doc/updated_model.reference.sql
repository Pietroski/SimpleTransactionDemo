-- Migration file last updated on the 1st May 2022

-- WARNING!! This file is for documentation only
-- WARNING!! This file is to keep record and track of the database schemas
-- Every model change, should also be taken care in this file for documentation purposes
-- This file is a helper to keep track which is the current state of the most updated models

CREATE TYPE CRYPTO_COINS AS ENUM
    (
        'BITCOIN',
        'DODGE-COIN',
        'ETHEREUM'
        );

CREATE TABLE IF NOT EXISTS transaction_record
(
    row_id         BIGSERIAL    NOT NULL,
    from_user_id   uuid         NOT NULL,
    from_wallet_id uuid         NOT NULL,
    to_user_id     uuid         NOT NULL,
    to_wallet_id   uuid         NOT NULL,
    coin           CRYPTO_COINS NOT NULL,
    amount         BIGINT       NOT NULL,
    created_at     timestamptz  NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS account
(
    row_id     BIGSERIAL   NOT NULL,
    user_id    uuid UNIQUE NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),

    CONSTRAINT pk_user_id PRIMARY KEY (user_id)
);

CREATE TABLE IF NOT EXISTS wallet
(
    row_id     BIGSERIAL    NOT NULL,
    wallet_id  uuid UNIQUE  NOT NULL,
    user_id    uuid         NOT NULL,
    coin       CRYPTO_COINS NOT NULL,
    balance    BIGINT       NOT NULL,
    created_at timestamptz  NOT NULL DEFAULT now(),
    updated_at timestamptz  NOT NULL DEFAULT now(),

    CONSTRAINT pk_wallet_id PRIMARY KEY (wallet_id)
);

ALTER TABLE wallet
    ADD CONSTRAINT fk_wallets_user_id
        FOREIGN KEY (user_id)
            REFERENCES account (user_id)
            ON DELETE CASCADE;
