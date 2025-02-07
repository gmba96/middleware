package resources

import (
	"github.com/gofrs/uuid"
	"middleware/config/internal/helpers"
	"middleware/config/internal/models"
)

/**
 * Get all resources
 */
func GetAllResources() ([]models.Resource, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM resources")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	resources := []models.Resource{}
	for rows.Next() {
		var data models.Resource
		err = rows.Scan(&data.Id, &data.Name, &data.UcaId)
		if err != nil {
			return nil, err
		}
		resources = append(resources, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return resources, err
}

/*
 * Get a resource by its id
 */
func GetResourceById(id uuid.UUID) (*models.Resource, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM resources WHERE id=?", id.String())
	helpers.CloseDB(db)

	var resource models.Resource
	err = row.Scan(&resource.Id, &resource.Name, &resource.UcaId)
	if err != nil {
		return nil, err
	}
	return &resource, err
}
