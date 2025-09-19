CREATE TABLE IF NOT EXISTS movies (
    id UUID PRIMARY KEY NOT NULL,
    title VARCHAR(255) NOT NULL,
    omdbid VARCHAR(255) NOT NULL,
    runtime_in_minutes INT NOT NULL
);