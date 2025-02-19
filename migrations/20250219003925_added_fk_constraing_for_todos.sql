-- +goose Up
-- +goose StatementBegin
ALTER TABLE todos
ADD COLUMN user_id BIGINT REFERENCES users (id) ON DELETE CASCADE;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE todos
DROP COLUMN user_id;

-- +goose StatementEnd
