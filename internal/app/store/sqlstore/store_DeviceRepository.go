package sqlstore

import (
	"database/sql"

	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/model"
	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/store"
)

// DeviceRepository ...
type DeviceRepository struct {
	store *Store
}

// Create ...
func (r *DeviceRepository) Create(d *model.Device) error {

	if err := d.Validate(); err != nil {
		return err
	}

	// If the DB cache is enabled
	if r.store.useCache {
		r.store.c.DeviceCreate(d)
	}

	// fmt.Println("DB QUERY: create device")
	return r.store.db.QueryRow(
		"INSERT INTO devices (name, cert_cn, last_access_time, expected) VALUES ($1, $2) RETURNING id;",
		d.DeviceName,
		d.DeviceCertCN,
		d.LastAccessTime,
		d.Expected,
	).Scan(&d.ID)
}

// Delete ...
func (r *DeviceRepository) Delete(id int) error {

	// Delete the device from the cache
	if r.store.c != nil {
		r.store.c.DeviceDelete(id)
	}

	d, err := r.Find(id)
	if err != nil {
		return store.ErrRecordNotFound
	}

	// fmt.Println("DB QUERY: delete device")
	return r.store.db.QueryRow(
		"DELETE FROM devices WHERE id = $1 RETURNING id",
		d.ID,
	).Scan(&d.ID)
}

// Find ...
func (r *DeviceRepository) Find(id int) (*model.Device, error) {

	d := &model.Device{}

	// If the DB cache is enabled
	var updateCache bool = false
	if r.store.useCache {
		d, err := r.store.c.DeviceFind(id)
		if err == nil {
			return d, nil
		}
		updateCache = true
	}

	// fmt.Println("DB QUERY: find device by id")
	if err := r.store.db.QueryRow(
		"SELECT id, name, cert_cn, last_access_time, expected FROM devices WHERE id = $1",
		id,
	).Scan(
		&d.ID,
		&d.DeviceName,
		&d.DeviceCertCN,
		&d.LastAccessTime,
		&d.Expected,
	); err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	// If the DB cache is enabled
	if updateCache {
		r.store.c.DeviceCreate(d)
	}

	return d, nil
}

// FindByCN ...
func (r *DeviceRepository) FindByCN(cn string) (*model.Device, error) {

	d := &model.Device{}

	// If the DB cache is enabled
	var updateCache bool = false
	if r.store.useCache {
		d, err := r.store.c.DeviceFindByCN(cn)
		if err == nil {
			return d, nil
		}
		updateCache = true
	}

	// fmt.Println("DB QUERY: find device by certificate cn")
	if err := r.store.db.QueryRow(
		"SELECT id, name, cert_cn, last_access_time, expected FROM devices WHERE cert_cn = $1",
		cn,
	).Scan(
		&d.ID,
		&d.DeviceName,
		&d.DeviceCertCN,
		&d.LastAccessTime,
		&d.Expected,
	); err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	// If the DB cache is enabled
	if updateCache {
		r.store.c.DeviceCreate(d)
	}

	return d, nil
}

// FindAll ...
func (r *DeviceRepository) FindAll() ([]model.Device, error) {

	// fmt.Println("DB QUERY: find all devices")
	rows, err := r.store.db.Query(
		"SELECT id, name, cert_cn, last_access_time, expected FROM devices;",
	)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	defer rows.Close()

	devices := make([]model.Device, 0, 10)

	for rows.Next() {
		var d model.Device
		if err := rows.Scan(
			&d.ID,
			&d.DeviceName,
			&d.DeviceCertCN,
			&d.LastAccessTime,
			&d.Expected,
		); err != nil {
			return nil, err
		}
		devices = append(devices, d)

		// If the DB cache is enabled
		if r.store.useCache {
			// It's safe to call just DeviceCreate() in this case
			r.store.c.DeviceCreate(&d)
		}
	}

	return devices, nil
}

// Clear ...
func (r *DeviceRepository) Clear() error {

	// fmt.Println("DB QUERY: remove all devices")
	r.store.db.QueryRow(
		"DELETE FROM devices",
	)

	// If the DB cache is enabled
	if r.store.useCache {
		r.store.c.DeviceClear()
	}

	return nil
}

// FindUsedAuthPatterns ...
func (r *DeviceRepository) FindUsedAuthPatterns(cn string) ([]string, error) {

	var deviceAuthPatterns []string
	var err error

	// If the DB cache is enabled
	var updateCache bool = false
	if r.store.useCache {
		deviceAuthPatterns, err = r.store.c.DeviceFindUsedAuthPatterns(cn)
		if err == nil {
			return deviceAuthPatterns, nil
		}
		updateCache = true
	}

	// fmt.Println("DB QUERY: get device authentication patterns history")
	rows, err := r.store.db.Query(
		"SELECT name FROM device_auth_patterns WHERE id IN (SELECT DISTINCT device_auth_pattern_id FROM auth_attempts WHERE device_id=(SELECT id FROM devices WHERE cert_cn = $1));",
		cn,
	)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	defer rows.Close()

	//! ToDo: check if absence of make causes any errors
	// deviceAuthPatterns = make([]string, 0)

	for rows.Next() {
		var s string
		if err := rows.Scan(&s); err != nil {
			return nil, err
		}

		deviceAuthPatterns = append(deviceAuthPatterns, s)
	}

	// If the DB cache is enabled
	if updateCache {
		err = r.store.c.DeviceSetUsedAuthPatterns(cn, deviceAuthPatterns)
		if err != nil {
			return nil, err
		}
	}

	return deviceAuthPatterns, nil
}

// FindTrustHistory ...
func (r *DeviceRepository) FindTrustHistory(cn string) ([]float32, error) {

	var trustHistory []float32
	var err error

	// If the DB cache is enabled
	var updateCache bool = false
	if r.store.useCache {
		trustHistory, err = r.store.c.DeviceFindTrustHistory(cn)
		if err == nil {
			return trustHistory, nil
		}
		updateCache = true
	}

	// fmt.Println("DB QUERY: get device trust scores history")
	rows, err := r.store.db.Query(
		"SELECT time, calculated_device_trust FROM auth_attempts WHERE device_id=(SELECT id FROM devices WHERE cert_cn = $1) ORDER BY time DESC LIMIT 10;",
		cn,
	)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	defer rows.Close()

	//! ToDo: check if absence of make causes any errors
	// trustHistory := make([]float32, 0)

	for rows.Next() {
		var t string
		var tr float32
		if err := rows.Scan(&t, &tr); err != nil {
			return nil, err
		}
		trustHistory = append(trustHistory, tr)
	}

	// If the DB cache is enabled
	if updateCache {
		err = r.store.c.DeviceSetTrustHistory(cn, trustHistory)
		if err != nil {
			return nil, err
		}
	}

	return trustHistory, nil
}

// FindLocationIPHistory ...
func (r *DeviceRepository) FindLocationIPHistory(cn string) ([]string, error) {

	var locationIPHistory []string
	var err error

	// If the DB cache is enabled
	var updateCache bool = false
	if r.store.useCache {
		locationIPHistory, err = r.store.c.DeviceFindLocationIPHistory(cn)
		if err == nil {
			return locationIPHistory, nil
		}
		updateCache = true
	}

	// fmt.Println("DB QUERY: get device location ip history")
	rows, err := r.store.db.Query(
		"SELECT DISTINCT location_ip FROM auth_attempts WHERE device_id=(SELECT id FROM devices WHERE cert_cn = $1);",
		cn,
	)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	defer rows.Close()

	// locationIPHistory := make([]string, 0)

	for rows.Next() {
		var ip string
		if err := rows.Scan(&ip); err != nil {
			return nil, err
		}
		locationIPHistory = append(locationIPHistory, ip)
	}

	if updateCache {
		err = r.store.c.DeviceSetLocationIPHistory(cn, locationIPHistory)
		if err != nil {
			return nil, err
		}
	}

	return locationIPHistory, nil
}

// FindServiceUsageHistory ...
func (r *DeviceRepository) FindServiceUsageHistory(cn string) ([]string, error) {

	var serviceUsageHistory []string
	var err error

	// If the DB cache is enabled
	var updateCache bool = false
	if r.store.useCache {
		serviceUsageHistory, err = r.store.c.DeviceFindServiceUsageHistory(cn)
		if err == nil {
			return serviceUsageHistory, nil
		}
		updateCache = true
	}

	// fmt.Println("DB QUERY: get device service usage history")
	rows, err := r.store.db.Query(
		"SELECT sni FROM services WHERE id IN (SELECT DISTINCT service_id FROM auth_attempts WHERE device_id=(SELECT id FROM devices WHERE cert_cn = $1));",
		cn,
	)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	defer rows.Close()

	//! ToDo: check if absence of make causes any errors
	// serviceUsageHistory = make([]string, 0, 10)

	for rows.Next() {
		var s string
		if err := rows.Scan(&s); err != nil {
			return nil, err
		}
		serviceUsageHistory = append(serviceUsageHistory, s)
	}

	// If the DB cache is enabled
	if updateCache {
		err = r.store.c.DeviceSetServiceUsageHistory(cn, serviceUsageHistory)
		if err != nil {
			return nil, err
		}
	}

	return serviceUsageHistory, nil
}

// FindUserUsage ...
func (r *DeviceRepository) FindUserUsageHistory(cn string) ([]string, error) {

	var userUsage []string
	var err error

	// If the DB cache is enabled
	var updateCache bool = false
	if r.store.useCache {
		userUsage, err = r.store.c.DeviceFindUserUsageHistory(cn)
		if err == nil {
			return userUsage, nil
		}
		updateCache = true
	}

	// fmt.Println("DB QUERY: get device user usage history")
	rows, err := r.store.db.Query(
		"SELECT name FROM users WHERE id IN (SELECT DISTINCT user_id FROM auth_attempts WHERE device_id=(SELECT id FROM devices WHERE cert_cn = $1));",
		cn,
	)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	defer rows.Close()

	//! ToDo: check if absence of make causes any errors
	// userUsage = make([]string, 0, 10)

	for rows.Next() {
		var user string
		if err := rows.Scan(&user); err != nil {
			return nil, err
		}
		userUsage = append(userUsage, user)
	}

	if updateCache {
		err = r.store.c.DeviceSetUserUsageHistory(cn, userUsage)
		if err != nil {
			return nil, err
		}
	}

	return userUsage, nil
}
