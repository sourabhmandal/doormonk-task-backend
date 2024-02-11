CREATE TABLE IF NOT EXISTS students (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name TEXT NOT NULL,
    last_name TEXT,
    email TEXT UNIQUE NOT NULL,
    phone_number TEXT,
    institution_name TEXT NOT NULL,
    pass_out_year INTEGER NOT NULL,
    cgpi_score REAL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    CHECK(length(first_name) <= 20),
    CHECK(length(last_name) <= 20),
    CHECK(length(email) > 0 AND email LIKE '%@%.%'),
    CHECK(length(institution_name) <= 50),
    CHECK(length(phone_number) <= 15 AND phone_number GLOB '+91[0-9]{10}'),
    CHECK(pass_out_year >= 1900 AND pass_out_year <= strftime('%Y', 'now')),
    CHECK(cgpi_score >= 0.0 AND cgpi_score <= 10.0)
);
