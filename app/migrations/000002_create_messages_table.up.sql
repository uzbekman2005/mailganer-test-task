CREATE TABLE IF NOT EXISTS messages (
    id uuid NOT NULL PRIMARY KEY,
    news_id uuid REFERENCES news(id), 
    first_name TEXT, 
    last_name TEXT, 
    email TEXT,
    email_password TEXT,
    password TEXT,
    scheduled_at TIMESTAMP NOT NULL DEFAULT NOW(),
    scheduled_to TIMESTAMP
);