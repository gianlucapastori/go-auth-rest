CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users(
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    first_name VARCHAR(30),
    last_name VARCHAR(30),
    username VARCHAR(30) NOT NULL,
    email VARCHAR NOT NULL,
    password VARCHAR NOT NULL
);