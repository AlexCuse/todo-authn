CREATE DATABASE todoauth;

\c todoauth;

-- Below copied from ./schema.sql

-- Create a table to store TODO items
CREATE TABLE todos (
                       id TEXT PRIMARY KEY DEFAULT gen_random_uuid()::TEXT,
                       user_id int NOT NULL, -- we're cheating and creating in same DB as authn
                       title VARCHAR(100) NOT NULL,
                       description TEXT,
                       completed BOOLEAN DEFAULT FALSE NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create an index on the user_id column in the todo_items table for faster queries
CREATE INDEX user_id_index ON todos(user_id);