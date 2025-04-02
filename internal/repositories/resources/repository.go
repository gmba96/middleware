package resources

import (
	"database/sql"
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
		err = rows.Scan(&data.Id, &data.UcaId, &data.Name)
		if err != nil {
			return nil, err
		}
		resources = append(resources, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return resources, err
}

/**
 * Get a resource by its id
 */
func GetResourceById(id uuid.UUID) (*models.Resource, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM resources WHERE id = ?", id.String())
	helpers.CloseDB(db)

	var resource models.Resource
	err = row.Scan(&resource.Id, &resource.UcaId, &resource.Name)
	if err != nil {
		return nil, err
	}

	return &resource, nil
}

// CreateResource inserts a new resource into the database
func CreateResource(ucaId int, name string) (*models.Resource, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	// Generate new UUID
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	// Execute insert query
	_, err = db.Exec(
		"INSERT INTO resources (id, uca_id, name) VALUES (?, ?, ?)",
		id.String(),
		ucaId,
		name,
	)
	if err != nil {
		return nil, err
	}

	return &models.Resource{
		Id:    &id,
		UcaId: ucaId,
		Name:  name,
	}, nil
}

// DeleteResource supprime une resource dans la base de données.
func DeleteResource(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	result, err := db.Exec("DELETE FROM resources WHERE id=?", id.String())
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// UpdateResource met à jour une resource dans la base de données et retourne la resource mise à jour.
func UpdateResource(id uuid.UUID, ucaId int, name string) (*models.Resource, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	// Exécution de la requête de mise à jour
	result, err := db.Exec("UPDATE resources SET uca_id=?, name=? WHERE id=?", ucaId, name, id.String())
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	resource, err := GetResourceById(id)
	if err != nil {
		return nil, err
	}
	return resource, nil
}
