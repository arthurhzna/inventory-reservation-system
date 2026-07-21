INSERT INTO users (
    uuid,
    name,
    email,
    password,
    role_id
)
VALUES (
    gen_random_uuid(),
    'Super Admin',
    'admin@example.com',
    '$2a$10$examplehashedpassword',
    1
)
ON CONFLICT (email) DO NOTHING;