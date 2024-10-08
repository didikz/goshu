CREATE TABLE IF NOT EXISTS urls(
    id BIGSERIAL PRIMARY KEY,
    slug VARCHAR(50) NOT NULL UNIQUE,
    original_url VARCHAR(255) NOT NULL,
    user_id BIGINT NULL,
    visit_count INT DEFAULT 0 NULL,
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp NULL,
    deleted_at timestamp NULL
);