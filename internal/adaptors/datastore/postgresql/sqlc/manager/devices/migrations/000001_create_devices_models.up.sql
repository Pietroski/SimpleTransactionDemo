-- Migration file created on the 1st May 2022

CREATE TABLE IF NOT EXISTS devices
(
    row_id       BIGSERIAL,
    device_id    uuid UNIQUE NOT NULL,
    user_id      uuid        NOT NULL,
    device_name  TEXT        NOT NULL,
    device_brand TEXT        NOT NULL,
    created_at   timestamptz DEFAULT now(),
    updated_at   timestamptz DEFAULT now(),

    CONSTRAINT pk_device PRIMARY KEY (device_id)
);
