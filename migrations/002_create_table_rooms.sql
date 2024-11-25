-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the rooms table
CREATE TABLE rooms (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,  -- Generate UUID v4 by default
    host_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255),
    FOREIGN KEY (host_id) REFERENCES users (id)
);

-- +migrate Down
DROP TABLE rooms;
