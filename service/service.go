package service

import (
	"apiassignment/dal"
	"apiassignment/model"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type PetService struct {
	DB *sql.DB
}

func NewPetService() *PetService {
	return &PetService{}
}

// CreatePet Insert a new pet into the database
func (s *PetService) CreatePet(pet model.Pet) (model.Pet, error) {
	_, err := dal.GetSQLDBConn().Exec("INSERT INTO pets (name, age, breed) VALUES (?, ?, ?)", pet.Name, pet.Age, pet.Breed)
	if err != nil {
		log.Printf("Error creating pet: %v", err) // Log the error
		return model.Pet{}, err
	}

	log.Printf("Pet created: ID=%d, Name=%s, Age=%d, Breed=%s", pet.ID, pet.Name, pet.Age, pet.Breed) // Log successful operation
	return pet, nil
}

func (s *PetService) GetPets() ([]model.Pet, error) {
	rows, err := dal.GetSQLDBConn().Query("SELECT id, name, age, breed FROM pets")
	if err != nil {
		log.Printf("Error retrieving pets: %v", err) // Log the error
		return nil, err
	}
	defer rows.Close()

	var pets []model.Pet
	for rows.Next() {
		var pet model.Pet
		if err := rows.Scan(&pet.ID, &pet.Name, &pet.Age, &pet.Breed); err != nil {
			log.Printf("Error scanning row: %v", err) // Log the error
			return nil, err
		}
		pets = append(pets, pet)
	}

	log.Printf("Retrieved %d pets", len(pets)) // Log successful operation
	return pets, nil
}

// GetPetByID Retrieve a specific pet by ID from the database
func (s *PetService) GetPetByID(id int) (model.Pet, error) {
	var pet model.Pet

	err := dal.GetSQLDBConn().QueryRow("SELECT id, name, age, breed FROM pets WHERE id = ?", id).
		Scan(&pet.ID, &pet.Name, &pet.Age, &pet.Breed)
	if err != nil {
		log.Printf("Error retrieving pet by ID: %v", err) // Log the error
		return model.Pet{}, err
	}

	log.Printf("Retrieved pet by ID: ID=%d, Name=%s, Age=%d, Breed=%s", pet.ID, pet.Name, pet.Age, pet.Breed) // Log successful operation
	return pet, nil
}

// UpdatePet Update a pet in the database
func (s *PetService) UpdatePet(id int, pet model.Pet) (model.Pet, error) {
	_, err := dal.GetSQLDBConn().Exec("UPDATE pets SET name = ?, age = ?, breed = ? WHERE id = ?", pet.Name, pet.Age, pet.Breed, id)
	if err != nil {
		log.Printf("Error updating pet: %v", err) // Log the error
		return model.Pet{}, err
	}

	log.Printf("Updated pet: ID=%d, Name=%s, Age=%d, Breed=%s", pet.ID, pet.Name, pet.Age, pet.Breed) // Log successful operation
	return pet, nil
}

// DeletePet Delete a pet from the database
func (s *PetService) DeletePet(id int) error {
	_, err := dal.GetSQLDBConn().Exec("DELETE FROM pets WHERE id = ?", id)
	if err != nil {
		log.Printf("Error deleting pet: %v", err) // Log the error
		return err
	}

	log.Printf("Deleted pet by ID: ID=%d", id) // Log successful operation
	return nil
}
