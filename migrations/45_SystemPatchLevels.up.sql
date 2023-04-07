CREATE TABLE IF NOT EXISTS system_patch_levels (
	id   SERIAL PRIMARY KEY,
	name VARCHAR(20) UNIQUE NOT NULL
);

INSERT INTO 
    system_patch_levels (name)
VALUES
    ('outdated'),
    ('patched');
