CREATE TABLE user
(
    id                  BINARY(16) NOT NULL,
    email               VARCHAR(321) NOT NULL UNIQUE,
    password            VARCHAR(256) NOT NULL,
    is_active           BOOLEAN NOT NULL DEFAULT 0,
    updated_at          TIMESTAMP NOT NULL,
    created_at          TIMESTAMP NOT NULL,

    PRIMARY KEY (id)
) CHARACTER SET utf8mb4
  COLLATE utf8mb4_bin;

