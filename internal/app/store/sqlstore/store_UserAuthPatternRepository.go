package sqlstore

import (
	"database/sql"

	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/model"
	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/store"
)

// UserAuthPatternRepository ...
type UserAuthPatternRepository struct {
	store *Store
}

// Create ...
func (r *UserAuthPatternRepository) Create(uap *model.UserAuthPattern) error {

	if err := uap.Validate(); err != nil {
		return err
	}

	// fmt.Println("DB QUERY: create userAuthPattern")
	return r.store.db.QueryRow(
		"INSERT INTO user_auth_patterns (name) VALUES ($1) RETURNING id;",
		uap.UserAuthPatternName,
	).Scan(&uap.ID)
}

// Delete ...
func (r *UserAuthPatternRepository) Delete(id int) error {

	uap, err := r.Find(id)
	if err != nil {
		return store.ErrRecordNotFound
	}

	// fmt.Println("DB QUERY: delete userAuthPattern")
	return r.store.db.QueryRow(
		"DELETE FROM user_auth_patterns WHERE id = $1 RETURNING id",
		uap.ID,
	).Scan(&uap.ID)
}

// Find ...
func (r *UserAuthPatternRepository) Find(id int) (*model.UserAuthPattern, error) {

	uap := &model.UserAuthPattern{}

	// fmt.Println("DB QUERY: get userAuthPattern by id")
	if err := r.store.db.QueryRow(
		"SELECT id, name FROM user_auth_patterns WHERE id = $1",
		id,
	).Scan(
		&uap.ID,
		&uap.UserAuthPatternName,
	); err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return uap, nil
}

// FindAll ...
func (r *UserAuthPatternRepository) FindAll() ([]model.UserAuthPattern, error) {

	// fmt.Println("DB QUERY: get all userAuthPatterns")
	rows, err := r.store.db.Query(
		"SELECT id, name FROM user_auth_patterns;",
	)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	defer rows.Close()

	uaps := make([]model.UserAuthPattern, 0, 10)

	for rows.Next() {
		var uap model.UserAuthPattern
		if err := rows.Scan(&uap.ID, &uap.UserAuthPatternName); err != nil {
			return nil, err
		}
		uaps = append(uaps, uap)
	}

	return uaps, nil
}

// Clear ...
func (r *UserAuthPatternRepository) Clear() error {

	// fmt.Println("DB QUERY: create userAuthPattern")
	r.store.db.QueryRow(
		"DELETE FROM user_auth_patterns",
	)

	return nil
}
