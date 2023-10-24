package handler

import (
	"apiassignment/model"
	"apiassignment/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatePet handles creating a new pet.
func CreatePet(c *gin.Context) {
	// Parse the request body into a Pet struct
	var pet model.Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to create the pet

	createdPet, err := service.NewPetService().CreatePet(pet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create pet"})
		return
	}

	c.JSON(http.StatusCreated, createdPet)
}

// GetPets retrieves all pets.
func GetPets(c *gin.Context) {
	// Call the service to get all pets
	pets, err := service.NewPetService().GetPets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve pets"})
		return
	}

	c.JSON(http.StatusOK, pets)
}

// GetPet retrieves a specific pet by ID.
func GetPet(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Call the service to get a specific pet by ID
	pet, err := service.NewPetService().GetPetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pet not found"})
		return
	}

	c.JSON(http.StatusOK, pet)
}

// UpdatePet updates a specific pet by ID.
func UpdatePet(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Parse the request body into a Pet struct
	var updatedPet model.Pet
	if err := c.ShouldBindJSON(&updatedPet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to update the pet
	updatedPet, err = service.NewPetService().UpdatePet(id, updatedPet)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pet not found"})
		return
	}

	c.JSON(http.StatusOK, updatedPet)
}

// DeletePet deletes a specific pet by ID.
func DeletePet(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Call the service to delete the pet
	err = service.NewPetService().DeletePet(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pet not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
