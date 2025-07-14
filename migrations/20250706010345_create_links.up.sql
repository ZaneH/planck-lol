CREATE EXTENSION IF NOT EXISTS citus;

CREATE TABLE links (
    id BIGSERIAL PRIMARY KEY,
    short_code VARCHAR(255) NOT NULL,
    long_url TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP WITH TIME ZONE
);

SELECT create_distributed_table('links', 'id')
