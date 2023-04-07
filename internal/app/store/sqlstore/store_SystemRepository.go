package sqlstore

import (
	"database/sql"
	"strings"

	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/model"
	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/store"
)

// SystemRepository ...
type SystemRepository struct {
	store *Store
}

// Get ...
func (r *SystemRepository) Get() (*model.System, error) {

	s := &model.System{}

	// If the DB cache is enabled
	var updateCache bool = false
	if r.store.useCache {
		s = r.store.c.SystemGet()
		if s.State != "" {
			return s, nil
		}
		updateCache = true
	}

	// fmt.Println("DB QUERY: get system state")
	if err := r.store.db.QueryRow(
		strings.Join([]string{
			"SELECT st.name, pl.name, threat_level, nl.name, network_threat_level FROM ztsnlac_system main",
			"INNER JOIN system_states st ON main.state_id = st.id",
			"INNER JOIN system_patch_levels pl ON main.patch_level_id = pl.id",
			"INNER JOIN network_states nl ON main.network_state_id = nl.id;",
		},
			" ",
		),
	).Scan(
		&s.State,
		&s.PatchLevel,
		&s.ThreatLevel,
		&s.NetworkState,
		&s.NetworkThreatLevel,
	); err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	// If the DB cache is enabled
	if updateCache {
		r.store.c.SystemSet(s)
	}

	return s, nil
}
