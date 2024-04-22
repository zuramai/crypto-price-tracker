CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email varchar(255) NOT NULL, 
    password varchar(255) NOT NULL 
);