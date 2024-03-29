-- Adding Third-Party modules
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Custom types
CREATE TYPE role AS ENUM ('user', 'admin');

-- Users table definition
CREATE TABLE IF NOT EXISTS users (
    id UUID DEFAULT uuid_generate_v4(),
    username VARCHAR(50) UNIQUE,
    password VARCHAR(128),
    phone_number VARCHAR(11),
    user_role role DEFAULT 'user',
    created_at DATE,
    updated_at DATE,

    PRIMARY KEY(id)
);

-- Contacts table definition
CREATE TABLE IF NOT EXISTS contacts (
    id UUID DEFAULT uuid_generate_v4(),
    owner_id UUID,
    first_name VARCHAR(255),
    last_name VARCHAR(255),

    PRIMARY KEY(id),
    FOREIGN KEY(owner_id) REFERENCES users(id)
);

-- Numbers table definition
CREATE TABLE IF NOT EXISTS numbers (
    id UUID DEFAULT uuid_generate_v4(),
    contact_id UUID,
    number VARCHAR(11),
    label VARCHAR(100),

    PRIMARY KEY(id),
    FOREIGN KEY(contact_id) REFERENCES contacts(id)
);

-- Emails table definition
CREATE TABLE IF NOT EXISTS emails (
    id UUID DEFAULT uuid_generate_v4(),
    contact_id UUID,
    address varchar(200),

    PRIMARY KEY(id),
    FOREIGN KEY(contact_id) REFERENCES contacts(id)
);

-- Users table constraints
ALTER TABLE users ADD UNIQUE (username);
ALTER TABLE users ALTER COLUMN username SET NOT NULL;
ALTER TABLE users ALTER COLUMN phone_number SET NOT NULL;
ALTER TABLE users ALTER COLUMN created_at SET NOT NULL;
ALTER TABLE users ALTER COLUMN updated_at SET NOT NULL;

-- Contact table constraints
ALTER TABLE contacts ALTER COLUMN first_name SET NOT NULL;

--Numbers table constraints
ALTER TABLE numbers ALTER COLUMN number SET NOT NULL;


-- Indexing tables
CREATE INDEX ON users (username);
CREATE INDEX ON contacts (first_name, last_name);
