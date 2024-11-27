-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS rooms (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS measurements (
    id SERIAL PRIMARY KEY,
    room_id INT REFERENCES rooms(id) ON DELETE CASCADE,
    temperature FLOAT,
    humidity FLOAT,
    timestamp BIGINT
);

INSERT INTO rooms (name) VALUES ('Master Bedroom');
INSERT INTO rooms (name) VALUES ('Office');
INSERT INTO rooms (name) VALUES ('Nursery');
INSERT INTO rooms (name) VALUES ('Living Room');
INSERT INTO rooms (name) VALUES ('Playroom');
INSERT INTO rooms (name) VALUES ('Basement');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE measurements;
DROP TABLE rooms;
-- +goose StatementEnd
