CREATE TABLE IF NOT EXISTS news (
    id uuid NOT NULL PRIMARY KEY,
    content TEXT,
    sender_email TEXT, 
    sender_email_password TEXT
); 

