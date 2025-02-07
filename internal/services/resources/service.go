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
			Code:    500,
		}
	}

	return resource, err
}
