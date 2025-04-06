package resources

import (
	"database/sql"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/config/internal/models"
	repository "middleware/config/internal/repositories/resources"
	"net/http"
)

func GetAllResources() ([]models.Resource, error) {
	var err error
	// calling repository
	resources, err := repository.GetAllResources()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving resources : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}
	return resources, nil
}

func GetResourceById(id uuid.UUID) (*models.Resource, error) {
	resource, err := repository.GetResourceById(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.CustomError{
				Message: "resource not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving resources : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}

	return resource, err
}

func CreateResource(ucaId int, name string) (*models.Resource, error) {
	
	resource, err := repository.CreateResource(ucaId, name)
	if err != nil {
		logrus.Errorf("error creating resource: %s", err.Error())
		return nil, &models.CustomError{
			Message: "Failed to create resource",
			Code:    http.StatusInternalServerError,
		}
	}
	return resource, nil
}

// RemoveResource effectue la suppression d'une resource après validation.
func RemoveResource(id uuid.UUID) error {
	err := repository.DeleteResource(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &models.CustomError{
				Message: "Resource non trouvée",
				Code:    http.StatusNotFound,
			}
		}
		return &models.CustomError{
			Message: "Échec de la suppression de la resource",
			Code:    http.StatusInternalServerError,
		}
	}
	return nil
}

// UpdateResource effectue des validations puis appelle la couche repository pour mettre à jour la resource.
func UpdateResource(id uuid.UUID, ucaId int, name string) (*models.Resource, error) {
	if name == "" {
		return nil, &models.CustomError{
			Message: "Le nom est requis",
			Code:    http.StatusBadRequest,
		}
	}

	resource, err := repository.UpdateResource(id, ucaId, name)
	if err != nil {
		logrus.Errorf("Erreur lors de la mise à jour de la resource: %s", err.Error())

		if err == sql.ErrNoRows || err.Error() == "no rows updated" {
			return nil, &models.CustomError{
				Message: "Resource non trouvée",
				Code:    http.StatusNotFound,
			}
		}
		return nil, &models.CustomError{
			Message: "Échec de la mise à jour de la resource",
			Code:    http.StatusInternalServerError,
		}
	}
	logrus.Infof("Resource mise à jour avec succès: %+v", resource)
	return resource, nil

}
