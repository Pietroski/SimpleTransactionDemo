-- Migration file created on the 1st May 2022

CREATE TYPE CRYPTO_CURRENCIES AS ENUM
    (
        'BITCOIN',
        'DODGE-COIN',
        'ETHEREUM',
        'PIETROSKI-COIN'
        );

CREATE TABLE IF NOT EXISTS account
(
    row_id     BIGSERIAL   NOT NULL,
    account_id uuid UNIQUE NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

ALTER TABLE wallet
    ADD CONSTRAINT pk_account_id
        PRIMARY KEY (account_id);

CREATE TABLE IF NOT EXISTS wallet
(
    row_id     BIGSERIAL         NOT NULL,
    wallet_id  uuid UNIQUE       NOT NULL,
    account_id uuid              NOT NULL,
    coin       CRYPTO_CURRENCIES NOT NULL,
    balance    BIGINT            NOT NULL,
    created_at timestamptz       NOT NULL DEFAULT now(),
    updated_at timestamptz       NOT NULL DEFAULT now(),

    CONSTRAINT pk_wallet_id PRIMARY KEY (wallet_id),
    UNIQUE (wallet_id, coin)
);

ALTER TABLE wallet
    ADD CONSTRAINT pk_wallet_id
        PRIMARY KEY (wallet_id),
    ADD CONSTRAINT unique_wallet_id_and_coin
        UNIQUE (wallet_id, coin),
    ADD CONSTRAINT fk_wallets_account_id
        FOREIGN KEY (account_id)
            REFERENCES account (account_id)
            ON DELETE CASCADE;

CREATE TABLE IF NOT EXISTS transaction_record
(
    row_id          BIGSERIAL         NOT NULL,
    from_account_id uuid              NOT NULL,
    from_wallet_id  uuid              NOT NULL,
    to_account_id   uuid              NOT NULL,
    to_wallet_id    uuid              NOT NULL,
    coin            CRYPTO_CURRENCIES NOT NULL,
    amount          BIGINT            NOT NULL,
    created_at      timestamptz       NOT NULL DEFAULT now()
);

ALTER TABLE transaction_record
    ADD CONSTRAINT fk_from_account_id
        FOREIGN KEY (from_account_id)
            REFERENCES account (account_id)
            ON DELETE CASCADE,
    ADD CONSTRAINT fk_from_wallet_id
        FOREIGN KEY (from_wallet_id)
            REFERENCES wallet (wallet_id)
            ON DELETE NO ACTION,
    ADD CONSTRAINT fk_to_account_id
        FOREIGN KEY (to_account_id)
            REFERENCES account (account_id)
            ON DELETE NO ACTION,
    Add CONSTRAINT fk_to_wallet_id
        FOREIGN KEY (to_wallet_id)
            REFERENCES wallet (wallet_id)
            ON DELETE NO ACTION;

CREATE TABLE IF NOT EXISTS entry_record
(
    row_id     BIGSERIAL         NOT NULL,
    account_id uuid              NOT NULL,
    wallet_id  uuid              NOT NULL,
    coin       CRYPTO_CURRENCIES NOT NULL,
    amount     BIGINT            NOT NULL,
    created_at timestamptz       NOT NULL DEFAULT now()
);

ALTER TABLE entry_record
    ADD CONSTRAINT fk_account_id
        FOREIGN KEY (account_id)
            REFERENCES account (account_id)
            ON DELETE CASCADE,
    ADD CONSTRAINT fk_wallet_id
        FOREIGN KEY (wallet_id)
            REFERENCES wallet (wallet_id)
            ON DELETE NO ACTION;
