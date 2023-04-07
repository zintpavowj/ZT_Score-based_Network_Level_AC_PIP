CREATE TABLE IF NOT EXISTS ztsnlac_system (
	state_id              INT UNIQUE NOT NULL,
  patch_level_id        INT UNIQUE NOT NULL,
  threat_level          INT UNIQUE NOT NULL,
  network_state_id      INT UNIQUE NOT NULL,
  network_threat_level  INT UNIQUE NOT NULL,
  CONSTRAINT fk_state
    FOREIGN KEY(state_id) 
  REFERENCES system_states(id),
  CONSTRAINT fk_patch_level
    FOREIGN KEY(patch_level_id) 
  REFERENCES system_patch_levels(id),
  CONSTRAINT fk_network_state
    FOREIGN KEY(network_state_id) 
  REFERENCES network_states(id)
);

INSERT INTO 
    ztsnlac_system (state_id, patch_level_id, threat_level, network_state_id, network_threat_level)
VALUES
    (1, 1, 1, 1, 0);
