CREATE TABLE IF NOT EXISTS network_states (
	id   SERIAL PRIMARY KEY,
	name VARCHAR(20) UNIQUE NOT NULL
);

INSERT INTO 
    network_states (name)
VALUES
    ('operational'),
    ('maintenance');
