-- i decided to move everything to a single sql file for simplicity sake, since its just
-- a silly project i want to avoid the verbosity of having multiple files for wathever change
-- i would want to do 

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users(
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    first_name VARCHAR(30),
    last_name VARCHAR(30),
    username VARCHAR(30) NOT NULL,
    email VARCHAR NOT NULL,
    password VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS groups(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(), 
    name VARCHAR(60) NOT NULL,
    description VARCHAR,
    creator_id UUID,
    CONSTRAINT fk_creator_id
        FOREIGN KEY(creator_id)
            REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS users_group_map(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "user_id" UUID,
    "group_id" UUID,
    CONSTRAINT fk_user_id
        FOREIGN KEY("user_id")
            REFERENCES users(id),
    CONSTRAINT fk_group_id
        FOREIGN KEY("group_id")
            REFERENCES groups(id)
);

CREATE TABLE IF NOT EXISTS tasks(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description VARCHAR,
    created_on TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    priority INT DEFAULT 0,
    status INT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS tasks_group_map(
    if UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "task_id" UUID,
    "group_id" UUID,
    CONSTRAINT fk_task_id
        FOREIGN KEY("task_id")
            REFERENCES tasks(id),
    CONSTRAINT fk_group_id
        FOREIGN KEY("group_id")
            REFERENCES groups(id)
);