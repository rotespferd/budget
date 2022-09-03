CREATE TABLE IF NOT EXISTS budget (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    amount INTEGER NOT NULL
);
CREATE UNIQUE INDEX idx_budget_id ON budget(id);