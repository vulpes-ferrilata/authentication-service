CREATE TABLE IF NOT EXISTS user_credentials(
   id VARCHAR(36) PRIMARY KEY,
   user_id VARCHAR(36) NOT NULL UNIQUE,
   email VARCHAR(255) NOT NULL UNIQUE,
   hash_password BLOB NOT NULL,
   version INT
);