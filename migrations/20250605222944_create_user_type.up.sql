CREATE TABLE IF NOT EXISTS user_types (
    id SERIAL PRIMARY KEY,
    code VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    created TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO user_types (code, name) VALUES
('01', 'Admin'),
('02', 'SBA'),
('03', 'Professors'),
('04', 'Students');