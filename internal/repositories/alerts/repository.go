package alerts

import (
	"github.com/gofrs/uuid"
	"middleware/config/internal/helpers"
	"middleware/config/internal/models"
)

/**
 * Get all alerts
 */
func GetAllAlerts() ([]models.Alert, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM alerts")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	alerts := []models.Alert{}
	for rows.Next() {
		var data models.Alert
		err = rows.Scan(&data.Id, &data.Email, &data.All, &data.ResourceId)
		if err != nil {
			return nil, err
		}
		alerts = append(alerts, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return alerts, err
}

/*
 * Get a alert by its id
 */
func GetAlertById(id uuid.UUID) (*models.Alert, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM alerts WHERE id=?", id.String())
	helpers.CloseDB(db)

	var alert models.Alert
	err = row.Scan(&alert.Id, &alert.Email, &alert.All, &alert.ResourceId)
	if err != nil {
		return nil, err
	}
	return &alert, err
}
