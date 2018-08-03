package repository

import (
	"database/sql"
	"gopostgres/src/modules/profile/model"
)

type profileRepositoryPostgres struct {
	db *sql.DB
}

func NewProfileRepositoryPostgres(db *sql.DB) *profileRepositoryPostgres {
	return &profileRepositoryPostgres{db}
}

func (r *profileRepositoryPostgres) Save(profile *model.Profile) error {
	query := `INSERT INTO "profile"("id", "first_name", "last_name", "email", "password", "created_at", "updated_at") VALUES($1, $2, $3, $4, $5, $6, $7)`

	steatment, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer steatment.Close()

	_, err = steatment.Exec(profile.ID, profile.FirstName, profile.LastName, profile.Email, profile.Password, profile.CreatedAt, profile.UpdatedAt)

	if err == nil {
		return err
	}

	return nil
}

func (r *profileRepositoryPostgres) Update(id string, profile *model.Profile) error {
	query := `UPDATE "profile" SET "first_name"=$1, "last_name"=$2, "email"=$3, "password"=$4, "updated_at"=$5 WHERE "id"=$6`

	steatment, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer steatment.Close()

	_, err = steatment.Exec(profile.FirstName, profile.LastName, profile.Email, profile.Password, profile.UpdatedAt, id)

	if err != nil {
		return err
	}

	return nil

}

func (r *profileRepositoryPostgres) Delete(id string) error {
	query := `DELETE FROM "profile" WHERE "id" = $1`

	steatment, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer steatment.Close()

	_, err = steatment.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

func (r *profileRepositoryPostgres) FindByID(id string) (*model.Profile, error) {
	query := `SELECT * FROM "profile" WHERE "id" = $1`

	var profile model.Profile

	steatment, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer steatment.Close()

	err = steatment.QueryRow(id).Scan(&profile.ID, &profile.FirstName, &profile.LastName, &profile.Email, &profile.Password, &profile.CreatedAt, &profile.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &profile, nil

}

func (r *profileRepositoryPostgres) FindAll() (model.Profiles, error) {
	query := `SELECT * FROM "profile"`

	var profiles model.Profiles

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var profile model.Profile

		err = rows.Scan(&profile.ID, &profile.FirstName, &profile.LastName, &profile.Email, &profile.Password, &profile.CreatedAt, &profile.UpdatedAt)

		if err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}

	return profiles, nil
}
