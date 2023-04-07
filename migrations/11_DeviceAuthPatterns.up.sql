CREATE TABLE IF NOT EXISTS device_auth_patterns (
	id   SERIAL PRIMARY KEY,
	name VARCHAR(100) UNIQUE NOT NULL
);

INSERT INTO 
    device_auth_patterns (name)
VALUES
    ('CertAuth'),
    ('HWTokenAuth');
