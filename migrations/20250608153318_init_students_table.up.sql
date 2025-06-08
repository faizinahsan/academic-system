CREATE TABLE IF NOT EXISTS students (
    id VARCHAR(32) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    phone VARCHAR(32),
    profile_picture TEXT,
    gender VARCHAR(16),
    major VARCHAR(100),
    faculty VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);