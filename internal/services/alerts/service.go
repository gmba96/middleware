package alerts

import (
	"database/sql"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/config/internal/models"
	repository "middleware/config/internal/repositories/alerts"
	"net/http"
)

func GetAllAlerts() ([]models.Alert, error) {
	var err error
	// calling repository
	alerts, err := repository.GetAllAlerts()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving alerts: %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}
	return alerts, nil
}

func GetAlertById(id uuid.UUID) (*models.Alert, error) {
	alert, err := repository.GetAlertById(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.CustomError{
				Message: "alert not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving alerts : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return alert, err
}

func DeleteAlert(id uuid.UUID) error {
	err := repository.DeleteAlert(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &models.CustomError{
				Message: "Alert not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error deleting alert: %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}
	return nil
}

func UpdateAlert(id uuid.UUID, email string, all bool, resourceId *uuid.UUID) (*models.Alert, error) {
	if email == "" {
		return nil, &models.CustomError{Message: "Email is required", Code: http.StatusBadRequest}
	}
	if !all && resourceId == nil {
		return nil, &models.CustomError{Message: "ResourceId is required when All is false", Code: http.StatusBadRequest}
	}

	existingAlert, err := repository.GetAlertById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &models.CustomError{Message: "Alert not found", Code: http.StatusNotFound}
		}
		logrus.Errorf("error retrieving alert: %s", err.Error())
		return nil, &models.CustomError{Message: "Something went wrong", Code: http.StatusInternalServerError}
	}

	existingAlert.Email = email
	existingAlert.All = all
	existingAlert.ResourceId = resourceId

	updatedAlert, err := repository.UpdateAlert(id, *existingAlert)
	if err != nil {
		logrus.Errorf("error updating alert: %s", err.Error())
		return nil, &models.CustomError{Message: "Something went wrong", Code: http.StatusInternalServerError}
	}

	return updatedAlert, nil
}

func CreateAlert(email string, all bool, resourceId *uuid.UUID) (*models.Alert, error) {
	if email == "" {
		return nil, &models.CustomError{Message: "Email is required", Code: http.StatusBadRequest}
	}
	if !all && resourceId == nil {
		return nil, &models.CustomError{Message: "ResourceId is required when All is false", Code: http.StatusBadRequest}
	}

	createdAlert, err := repository.CreateAlert(email, all, resourceId)
	if err != nil {
		logrus.Errorf("error creating alert: %s", err.Error())
		return nil, &models.CustomError{Message: "Something went wrong", Code: http.StatusInternalServerError}
	}

	return createdAlert, nil
}
