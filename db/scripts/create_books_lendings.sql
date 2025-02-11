CREATE TABLE IF NOT EXISTS book_lendings (
    id SERIAL PRIMARY KEY,
    book_id INT NOT NULL,
    user_id INT NOT NULL,
    lend_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    return_due_date TIMESTAMP,
    returned_date timestamp,
    FOREIGN KEY (book_id) REFERENCES books (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);