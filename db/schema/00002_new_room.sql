-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

INSERT INTO rooms (name) VALUES ('Front Bedroom');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DELETE FROM rooms WHERE name = 'Front Bedroom';
-- +goose StatementEnd
