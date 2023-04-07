CREATE TABLE IF NOT EXISTS service_software_patch_levels (
	id   SERIAL PRIMARY KEY,
	name VARCHAR(100) UNIQUE NOT NULL
);

INSERT INTO 
    service_software_patch_levels (name)
VALUES
    ('up-to-date'),
    ('outdated');
