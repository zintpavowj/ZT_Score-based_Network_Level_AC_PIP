CREATE TABLE IF NOT EXISTS system_states (
	id   SERIAL PRIMARY KEY,
	name VARCHAR(20) UNIQUE NOT NULL
);

INSERT INTO 
    system_states (name)
VALUES
    ('operational'),
    ('maintenance');
