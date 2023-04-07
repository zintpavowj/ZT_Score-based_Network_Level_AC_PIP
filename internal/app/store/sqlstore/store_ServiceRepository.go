package sqlstore

import (
	"database/sql"
	"strings"

	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/model"
	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/store"
)

// ServiceRepository ...
type ServiceRepository struct {
	store *Store
}

// Create ...
func (r *ServiceRepository) Create(s *model.Service) error {

	if err := s.Validate(); err != nil {
		return err
	}

	if r.store.useCache {
		r.store.c.ServiceCreate(s)
	}

	// fmt.Println("DB QUERY: create service")
	return r.store.db.QueryRow(
		strings.Join([]string{
			"INSERT INTO services (name, sni, data_sensitivity, software_patch_level_id)",
			"VALUES ($1, $2, $3, (SELECT id FROM service_software_patch_levels WHERE name = $4))",
			"RETURNING id;",
		},
			" ",
		),
		s.ServiceName,
		s.ServiceSNI,
		s.DataSensitivity,
		s.SoftwarePatchLevel,
	).Scan(&s.ID)
}

// Delete ...
func (r *ServiceRepository) Delete(id int) error {

	// Delete the device from the cache
	if r.store.c != nil {
		r.store.c.ServiceDelete(id)
	}

	s, err := r.Find(id)
	if err != nil {
		return store.ErrRecordNotFound
	}

	// fmt.Println("DB QUERY: delete service")
	return r.store.db.QueryRow(
		"DELETE FROM services WHERE id = $1 RETURNING id",
		s.ID,
	).Scan(&s.ID)
}

// Find ...
func (r *ServiceRepository) Find(id int) (*model.Service, error) {

	s := &model.Service{}

	// If the DB cache is enabled
	var updateCache bool = false
	if r.store.useCache {
		s, err := r.store.c.ServiceFind(id)
		if err == nil {
			return s, nil
		}
		updateCache = true
	}

	// fmt.Println("DB QUERY: get service by id")
	if err := r.store.db.QueryRow(strings.Join([]string{
		"SELECT s.id, s.name, s.sni, s.data_sensitivity, pl.name FROM services s",
		"INNER JOIN service_software_patch_levels pl ON s.software_patch_level_id = pl.id",
		"WHERE s.id = $1",
	},
		" ",
	),
		id,
	).Scan(
		&s.ID,
		&s.ServiceName,
		&s.ServiceSNI,
		&s.DataSensitivity,
		&s.SoftwarePatchLevel,
	); err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	// If the DB cache is enabled
	if updateCache {
		r.store.c.ServiceCreate(s)
	}

	return s, nil
}

// FindBySNI ...
func (r *ServiceRepository) FindBySNI(sni string) (*model.Service, error) {

	s := &model.Service{}

	// If the DB cache is enabled
	var updateCache bool = false
	if r.store.useCache {
		s, err := r.store.c.ServiceFindBySNI(sni)
		if err == nil {
			return s, nil
		}
		updateCache = true
	}

	// fmt.Println("DB QUERY: get service by sni")
	if err := r.store.db.QueryRow(strings.Join([]string{
		"SELECT s.id, s.name, s.sni, s.data_sensitivity, pl.name FROM services s",
		"INNER JOIN service_software_patch_levels pl ON s.software_patch_level_id = pl.id",
		"WHERE s.sni = $1",
	},
		" ",
	), sni).Scan(
		&s.ID,
		&s.ServiceName,
		&s.ServiceSNI,
		&s.DataSensitivity,
		&s.SoftwarePatchLevel,
	); err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	// If the DB cache is enabled
	if updateCache {
		r.store.c.ServiceCreate(s)
	}

	return s, nil
}

// FindAll ...
func (r *ServiceRepository) FindAll() ([]model.Service, error) {

	// fmt.Println("DB QUERY: get all services")
	rows, err := r.store.db.Query(
		strings.Join([]string{
			"SELECT s.id, s.name, sni, data_sensitivity, l.name FROM services s",
			"INNER JOIN service_software_patch_levels l ON l.id = s.software_patch_level_id;",
		},
			" ",
		),
	)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	defer rows.Close()

	services := make([]model.Service, 0, 10)

	for rows.Next() {
		var s model.Service
		if err := rows.Scan(
			&s.ID,
			&s.ServiceName,
			&s.ServiceSNI,
			&s.DataSensitivity,
			&s.SoftwarePatchLevel,
		); err != nil {
			return nil, err
		}
		services = append(services, s)

		// If the DB cache is enabled
		if r.store.useCache {
			// It's safe to call just DeviceCreate() in this case
			r.store.c.ServiceCreate(&s)
		}
	}

	return services, nil
}

// Clear ...
func (r *ServiceRepository) Clear() error {

	// fmt.Println("DB QUERY: delete all services")
	r.store.db.QueryRow(
		"DELETE FROM services",
	)

	// If the DB cache is enabled
	if r.store.useCache {
		r.store.c.ServiceClear()
	}

	return nil
}
