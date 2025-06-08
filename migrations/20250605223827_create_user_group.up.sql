CREATE TABLE IF NOT EXISTS user_groups (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    user_type_id INTEGER NOT NULL,
    created TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_user_type FOREIGN KEY (user_type_id) REFERENCES user_types (id) ON DELETE CASCADE
);