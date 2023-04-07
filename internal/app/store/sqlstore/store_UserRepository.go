package sqlstore

import (
	"database/sql"

	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/model"
	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/store"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {

	if err := u.Validate(); err != nil {
		return err
	}

	// If the DB cache is enabled
	if r.store.useCache {
		r.store.c.UserCreate(u)
	}

	// fmt.Println("DB QUERY: create user")
	return r.store.db.QueryRow(
		"INSERT INTO users (name, email, last_access_time, expected, access_time_min, access_time_max, database_update_time, password_failed_attempts) VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP, &7) RETURNING id;",
		u.Name,
		u.Email,
		u.LastAccessTime,
		u.Expected,
		u.AccessTimeMin,
		u.AccessTimeMax,
		u.PasswordFailedAttempts,
	).Scan(&u.ID)
}

// Delete ...
func (r *UserRepository) Delete(id int) error {

	// Delete the user from the cache
	if r.store.c != nil {
		r.store.c.UserDelete(id)
	}

	u, err := r.Find(id)
	if err != nil {
		return store.ErrRecordNotFound
	}

	// fmt.Println("DB QUERY: delete user")
	return r.store.db.QueryRow(
		"DELETE FROM users WHERE id = $1 RETURNING id",
		u.ID,
	).Scan(&u.ID)
}

// Find ...
func (r *UserRepository) Find(id int) (*model.User, error) {

	u := &model.User{}

	// If the DB cache is enabled
	var updateCache bool = false
	if r.store.useCache {
		u, err := r.store.c.UserFind(id)
		if err == nil {
			return u, nil
		}
		updateCache = true
	}

	// fmt.Println("DB QUERY: get user by id")
	if err := r.store.db.QueryRow(
		"SELECT id, name, email, last_access_time, expected, access_time_min, access_time_max, database_update_time, password_failed_attempts FROM users WHERE id = $1",
		id,
	).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.LastAccessTime,
		&u.Expected,
		&u.AccessTimeMin,
		&u.AccessTimeMax,
		&u.DatabaseUpdateTime,
		&u.PasswordFailedAttempts,
	); err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	// If the DB cache is enabled
	if updateCache {
		r.store.c.UserCreate(u)
	}

	return u, nil
}

// FindByName ...
func (r *UserRepository) FindByName(name string) (*model.User, error) {

	u := &model.User{}

	// If the DB cache is enabled
	var updateCache bool = false
	if r.store.useCache {
		u, err := r.store.c.UserFindByName(name)
		if err == nil {
			return u, nil
		}
		updateCache = true
	}

	// fmt.Println("DB QUERY: get user by name")
	if err := r.store.db.QueryRow(
		"SELECT id, name, email, last_access_time, expected, access_time_min, access_time_max, database_update_time, password_failed_attempts FROM users WHERE name = $1",
		name,
	).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.LastAccessTime,
		&u.Expected,
		&u.AccessTimeMin,
		&u.AccessTimeMax,
		&u.DatabaseUpdateTime,
		&u.PasswordFailedAttempts,
	); err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	// If the DB cache is enabled
	if updateCache {
		r.store.c.UserCreate(u)
	}

	return u, nil
}

// FindAll ...
func (r *UserRepository) FindAll() ([]model.User, error) {

	// fmt.Println("DB QUERY: get all users")
	rows, err := r.store.db.Query(
		"SELECT id, name, email, last_access_time, expected, access_time_min, access_time_max, database_update_time, password_failed_attempts FROM users;",
	)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	defer rows.Close()

	users := make([]model.User, 0, 10)

	for rows.Next() {
		var u model.User
		if err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Email,
			&u.LastAccessTime,
			&u.Expected,
			&u.AccessTimeMin,
			&u.AccessTimeMax,
			&u.DatabaseUpdateTime,
			&u.PasswordFailedAttempts,
		); err != nil {
			return nil, err
		}
		users = append(users, u)

		// If the DB cache is enabled
		if r.store.useCache {
			// It's safe to call just UserCreate() in this case
			r.store.c.UserCreate(&u)
		}
	}

	return users, nil
}

// Clear ...
func (r *UserRepository) Clear() error {

	// fmt.Println("DB QUERY: delete all users")
	r.store.db.QueryRow(
		"DELETE FROM users",
	)

	// If the DB cache is enabled
	if r.store.useCache {
		r.store.c.UserClear()
	}

	return nil
}

// FindUsedAuthPatterns ...
func (r *UserRepository) FindUsedAuthPatterns(name string) ([]string, error) {

	var userAuthPatterns []string
	var err error

	// If the DB cache is enabled
	var updateCache bool = false
	if r.store.useCache {
		userAuthPatterns, err = r.store.c.UserFindUsedAuthPatterns(name)
		if err == nil {
			return userAuthPatterns, nil
		}
		updateCache = true
	}

	// fmt.Println("DB QUERY: get user authentication patterns history")
	rows, err := r.store.db.Query(
		"SELECT name FROM user_auth_patterns WHERE id IN (SELECT DISTINCT user_auth_pattern_id FROM auth_attempts WHERE user_id=(SELECT id FROM users WHERE name = $1));",
		name,
	)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	defer rows.Close()

	//! ToDo: check if absence of make causes any errors
	// userAuthPatterns = make([]string, 0, 10)

	for rows.Next() {
		var s string
		if err := rows.Scan(&s); err != nil {
			return nil, err
		}
		userAuthPatterns = append(userAuthPatterns, s)
	}

	// If the DB cache is enabled
	if updateCache {
		err = r.store.c.UserSetUsedAuthPatterns(name, userAuthPatterns)
		if err != nil {
			return nil, err
		}
	}

	return userAuthPatterns, nil
}

// FindTrustHistory ...
func (r *UserRepository) FindTrustHistory(name string) ([]float32, error) {

	var trustHistory []float32
	var err error

	// If the DB cache is enabled
	var updateCache bool = false
	if r.store.useCache {
		trustHistory, err = r.store.c.UserFindTrustHistory(name)
		if err == nil {
			return trustHistory, nil
		}
		updateCache = true
	}

	// fmt.Println("DB QUERY: get user trust values history")
	rows, err := r.store.db.Query(
		"SELECT time, calculated_user_trust FROM auth_attempts WHERE user_id=(SELECT id FROM users WHERE name = $1) ORDER BY time DESC LIMIT 10;",
		name,
	)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	defer rows.Close()

	//! ToDo: check if absence of make causes any errors
	// trustHistory = make([]float32, 0, 10)

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
		err = r.store.c.UserSetTrustHistory(name, trustHistory)
		if err != nil {
			return nil, err
		}
	}

	return trustHistory, nil
}

// FindAccessRateHistory ...
func (r *UserRepository) FindAccessRateHistory(name string) ([]float32, error) {

	var accessRateHistory []float32
	var err error

	// If the DB cache is enabled
	var updateCache bool = false
	if r.store.useCache {
		accessRateHistory, err = r.store.c.UserFindAccessRateHistory(name)
		if err == nil {
			return accessRateHistory, nil
		}
		updateCache = true
	}

	// fmt.Println("DB QUERY: get user access rate history")
	rows, err := r.store.db.Query(
		"SELECT time, access_rate FROM auth_attempts WHERE user_id=(SELECT id FROM users WHERE name = $1) ORDER BY time DESC LIMIT 10;",
		name,
	)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	defer rows.Close()

	//! ToDo: check if absence of make causes any errors
	// accessRateHistory = make([]float32, 0, 10)

	for rows.Next() {
		var t string
		var ar float32
		if err := rows.Scan(&t, &ar); err != nil {
			return nil, err
		}
		accessRateHistory = append(accessRateHistory, ar)
	}

	// If the DB cache is enabled
	if updateCache {
		err = r.store.c.UserSetAccessRateHistory(name, accessRateHistory)
		if err != nil {
			return nil, err
		}
	}

	return accessRateHistory, nil
}

// FindInputBehaviorHistory ...
func (r *UserRepository) FindInputBehaviorHistory(name string) ([]float32, error) {

	var inputBehaviorHistory []float32
	var err error

	// If the DB cache is enabled
	var updateCache bool = false
	if r.store.useCache {
		inputBehaviorHistory, err = r.store.c.UserFindInputBehaviorHistory(name)
		if err == nil {
			return inputBehaviorHistory, nil
		}
		updateCache = true
	}

	// fmt.Println("DB QUERY: get user input behavior history")
	rows, err := r.store.db.Query(
		"SELECT time, input_behavior FROM auth_attempts WHERE user_id=(SELECT id FROM users WHERE name = $1) ORDER BY time DESC LIMIT 10;",
		name,
	)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	defer rows.Close()

	//! ToDo: check if absence of make causes any errors
	// inputBehaviorHistory := make([]float32, 0, 10)

	for rows.Next() {
		var t string
		var ar float32
		if err := rows.Scan(&t, &ar); err != nil {
			return nil, err
		}
		inputBehaviorHistory = append(inputBehaviorHistory, ar)
	}

	// If the DB cache is enabled
	if updateCache {
		err = r.store.c.UserSetInputBehaviorHistory(name, inputBehaviorHistory)
		if err != nil {
			return nil, err
		}
	}

	return inputBehaviorHistory, nil
}

// FindServiceUsage ...
func (r *UserRepository) FindServiceUsageHistory(name string) ([]string, error) {

	var serviceUsageHistory []string
	var err error

	// If the DB cache is enabled
	var updateCache bool = false
	if r.store.useCache {
		serviceUsageHistory, err = r.store.c.UserFindServiceUsageHistory(name)
		if err == nil {
			return serviceUsageHistory, nil
		}
		updateCache = true
	}

	// fmt.Println("DB QUERY: get user service usage history")
	rows, err := r.store.db.Query(
		"SELECT sni FROM services WHERE id IN (SELECT DISTINCT service_id FROM auth_attempts WHERE user_id=(SELECT id FROM users WHERE name = $1));",
		name,
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
		err = r.store.c.UserSetServiceUsageHistory(name, serviceUsageHistory)
		if err != nil {
			return nil, err
		}
	}

	return serviceUsageHistory, nil
}

// FindDeviceUsageHistory ...
func (r *UserRepository) FindDeviceUsageHistory(name string) ([]string, error) {

	var deviceUsageHistory []string
	var err error

	// If the DB cache is enabled
	var updateCache bool = false
	if r.store.useCache {
		deviceUsageHistory, err = r.store.c.UserFindDeviceUsageHistory(name)
		if err == nil {
			return deviceUsageHistory, nil
		}
		updateCache = true
	}

	// fmt.Println("DB QUERY: get user device usage history")
	rows, err := r.store.db.Query(
		"SELECT name, cert_cn FROM devices WHERE id IN (SELECT DISTINCT device_id FROM auth_attempts WHERE user_id=(SELECT id FROM users WHERE name = $1));",
		name,
	)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	defer rows.Close()

	//! ToDo: check if absence of make causes any errors
	// deviceUsageHistory := make([]string, 0, 10)

	for rows.Next() {
		var d, cn string
		if err := rows.Scan(&d, &cn); err != nil {
			return nil, err
		}
		deviceUsageHistory = append(deviceUsageHistory, cn)
	}

	// If the DB cache is enabled
	if updateCache {
		err = r.store.c.UserSetDeviceUsageHistory(name, deviceUsageHistory)
		if err != nil {
			return nil, err
		}
	}

	return deviceUsageHistory, nil
}
