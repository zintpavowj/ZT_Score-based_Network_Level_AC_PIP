CREATE TABLE IF NOT EXISTS user_auth_patterns (
	id   SERIAL PRIMARY KEY,
	name VARCHAR(100) UNIQUE NOT NULL
);

INSERT INTO 
    user_auth_patterns (name)
VALUES
    ('PasswAuth'),
    ('HWTokenAuth'),
    ('FaceIDAuth');
