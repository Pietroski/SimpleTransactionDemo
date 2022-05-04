-- Migration file last updated on the 1st May 2022

-- WARNING!! This file is for documentation only
-- WARNING!! This file is to keep record and track of the database schemas
-- Every model change, should also be taken care in this file for documentation purposes
-- This file is a helper to keep track which is the current state of the most updated models

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
