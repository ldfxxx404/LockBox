-- +goose Up
-- +goose StatementBegin
INSERT INTO users (id,email, name, password, is_admin, storage_limit)
VALUES
    (1,'admin@admin.f', 'admin', '$2a$10$j3hKqcP043zG3Uo32dN3S.z.Q50P6BXf/X09F2hBAJTtSqyQBtl/S', true, 20),
    (2,'user@user.d', 'user', '$2a$10$Z73cr5Y5c.ef3m.aOYDjy.7pafVPXfU31GHkFl.zIgV2.Zu7hrXIO', false, 20)
    ON CONFLICT (email) DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE email IN ('admin@admin.f', 'user@user.d');
-- +goose StatementEnd
