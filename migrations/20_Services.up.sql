CREATE TABLE IF NOT EXISTS services (
	id   SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
    sni VARCHAR(255) UNIQUE NOT NULL,
    data_sensitivity INT NOT NULL,
    software_patch_level_id INT NOT NULL,
    CONSTRAINT fk_software_patch_level
      FOREIGN KEY(software_patch_level_id) 
	  REFERENCES service_software_patch_levels(id)
);

INSERT INTO 
    services (name, sni, data_sensitivity, software_patch_level_id)
VALUES
    ('overleaf', 'overleaf.ztsnlac.com', '5', '1'),
    ('overleaf', 'overleaf.example.com', '7', '2'),
    ('gitlab', 'gitlab.ztsnlac.com', '10', '1'),
    ('gitlab', 'gitlab.com', '5', '1');
