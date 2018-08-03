package main

import (
	"fmt"
	"gopostgres/config"
	"gopostgres/src/modules/profile/model"
	"gopostgres/src/modules/profile/repository"
)

func main() {
	fmt.Println("Go Postgres SQL")

	db, err := config.GetPostgresDB()

	if err != nil {
		fmt.Println(err)
	}

	profileRepositoryPostgres := repository.NewProfileRepositoryPostgres(db)

	// EXAMPLE INSERT DATA
	// dans := model.NewProfile()
	// dans.ID = "D3"
	// dans.FirstName = "Hamdan"
	// dans.LastName = "Muhammad Al Fatih"
	// dans.Email = "hamdan@qwords.co.id"
	// dans.Password = "1234567890"

	// err = saveProfile(dans, profileRepositoryPostgres)

	// EXAMPLE UPDATE DATA
	// dans := model.NewProfile()
	// dans.ID = "D1"
	// dans.FirstName = "Dans"
	// dans.LastName = "Muhammad"
	// dans.Email = "hamdan@qwords.com"
	// dans.Password = "1234567890"

	// err = updateProfile(dans, profileRepositoryPostgres)

	//EXAMPLE DELETE
	//err = deleteProfile("D1", profileRepositoryPostgres)

	//GET PROFILE BY ID
	// profile, err := getProfile("D2", profileRepositoryPostgres)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("===============================")
	// fmt.Println(profile)

	//GET ALL DATA PROFILES
	profiles, err := getProfiles(profileRepositoryPostgres)

	if err != nil {
		fmt.Println(err)
	} else {
		for _, v := range profiles {
			fmt.Println(v)
		}
	}

}

func saveProfile(p *model.Profile, repo repository.ProfileRepository) error {
	err := repo.Save(p)

	if err != nil {
		return err
	}

	return nil
}

func updateProfile(p *model.Profile, repo repository.ProfileRepository) error {
	err := repo.Update(p.ID, p)

	if err != nil {
		return err
	}

	return nil
}

func deleteProfile(id string, repo repository.ProfileRepository) error {
	err := repo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func getProfile(id string, repo repository.ProfileRepository) (*model.Profile, error) {
	profile, err := repo.FindByID(id)

	if err != nil {
		return nil, err
	}

	return profile, nil
}

func getProfiles(repo repository.ProfileRepository) (model.Profiles, error) {
	profiles, err := repo.FindAll()

	if err != nil {
		return nil, err
	}

	return profiles, nil
}
