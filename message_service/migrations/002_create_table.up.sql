CREATE TYPE priority_type AS ENUM ('high', 'medium', 'low');

CREATE TABLE messages (
    id UUID NOT NULL PRIMARY KEY,
    content TEXT NOT NULL,
    priority priority_type NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
