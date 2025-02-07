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
