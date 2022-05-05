INSERT INTO account
    (account_id, created_at, updated_at)
VALUES ('410ae8d3-1573-4f37-8b38-1bdece117db5', now(), now()),
       ('dcf5baae-ce7e-426e-8c75-6f7212aace94', now(), now()),
       ('e0c3dc07-abc9-4ec2-8994-3215b11ee935', now(), now()),
       ('44adb9a0-6575-42f4-8fce-c9f90ddd71f1', now(), now())
RETURNING *;

INSERT INTO wallet
    (wallet_id, account_id, coin, balance, created_at, updated_at)
VALUES ('81f84772-cf13-4a98-9ff7-48eabadfb5ef', '410ae8d3-1573-4f37-8b38-1bdece117db5', 'BITCOIN', 0, now(), now()),
       ('30313dad-2ba7-4572-ad6a-56379109d928', 'dcf5baae-ce7e-426e-8c75-6f7212aace94', 'DODGE-COIN', 0, now(), now()),
       ('0735ffae-f124-4d79-a614-b28452e8f14c', 'e0c3dc07-abc9-4ec2-8994-3215b11ee935', 'ETHEREUM', 0, now(), now()),
       ('73342efd-23f2-4c89-9d3c-7668e2a28925', '44adb9a0-6575-42f4-8fce-c9f90ddd71f1', 'PIETROSKI-COIN', 0, now(), now()),

       ('f9561336-0a1d-44d6-9c85-cbd4b19204d9', '410ae8d3-1573-4f37-8b38-1bdece117db5', 'BITCOIN', 0, now(), now()),
       ('6406ba42-0442-4945-ac06-21f282f50c03', 'dcf5baae-ce7e-426e-8c75-6f7212aace94', 'DODGE-COIN', 0, now(), now()),
       ('d70c69b3-f901-4d35-9b67-71b9a83fe253', 'e0c3dc07-abc9-4ec2-8994-3215b11ee935', 'ETHEREUM', 0, now(), now()),
       ('1ebc3c98-cd2b-4e15-932d-dd2b9a599bed', '44adb9a0-6575-42f4-8fce-c9f90ddd71f1', 'PIETROSKI-COIN', 0, now(), now()),

       ('798af056-9f34-4a6e-aacc-d721560838f4', '410ae8d3-1573-4f37-8b38-1bdece117db5', 'BITCOIN', 0, now(), now()),
       ('65322b22-e040-4362-b5b3-1e27a9d8d241', 'dcf5baae-ce7e-426e-8c75-6f7212aace94', 'DODGE-COIN', 0, now(), now()),
       ('bb73bc98-e0db-474f-8644-ed0e3789e521', 'e0c3dc07-abc9-4ec2-8994-3215b11ee935', 'ETHEREUM', 0, now(), now()),
       ('7b56f874-3f50-414f-8660-72b949e00f98', '44adb9a0-6575-42f4-8fce-c9f90ddd71f1', 'PIETROSKI-COIN', 0, now(), now()),

       ('573ffe3d-9f27-4f0a-a0df-d275c861181f', '410ae8d3-1573-4f37-8b38-1bdece117db5', 'BITCOIN', 0, now(), now()),
       ('e189bf34-f26f-4ce8-9ede-c25d05c96328', 'dcf5baae-ce7e-426e-8c75-6f7212aace94', 'DODGE-COIN', 0, now(), now()),
       ('562fdc01-fe18-41db-8d18-04290d7048c3', 'e0c3dc07-abc9-4ec2-8994-3215b11ee935', 'ETHEREUM', 0, now(), now()),
       ('f148e59d-fcce-413c-a1cb-a00a968737d8', '44adb9a0-6575-42f4-8fce-c9f90ddd71f1', 'PIETROSKI-COIN', 0, now(), now())
RETURNING *;
