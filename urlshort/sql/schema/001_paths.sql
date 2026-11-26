-- +goose up
CREATE TABLE paths (
    id INTEGER PRIMARY KEY,
    path TEXT NOT NULL UNIQUE,
    url TEXT NOT NULL
);
-- +goose down
DROP TABLE paths;