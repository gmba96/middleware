package alerts

import (
	"database/sql"
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

func DeleteAlert(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	result, err := db.Exec("DELETE FROM alerts WHERE id=?", id.String())
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func UpdateAlert(id uuid.UUID, alert models.Alert) (*models.Alert, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("UPDATE alerts SET email=?, all_boolean=?, resource_id=? WHERE id=?",
		alert.Email, alert.All, alert.ResourceId, id.String())
	if err != nil {
		return nil, err
	}

	updatedAlert, err := GetAlertById(id)
	if err != nil {
		return nil, err
	}

	return updatedAlert, nil
}

func CreateAlert(email string, all bool, resourceId *uuid.UUID) (*models.Alert, error) {
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

	if resourceId != nil {
		_, err = db.Exec("INSERT INTO alerts (id, email, all_boolean, resource_id) VALUES (?, ?, ?, ?)",
			id, email, all, resourceId.String())
	} else {
		_, err = db.Exec("INSERT INTO alerts (id, email, all_boolean, resource_id) VALUES (?, ?, ?, ?)",
			id, email, all, nil)
	}
	if err != nil {
		return nil, err
	}

	createdAlert, err := GetAlertById(id)
	if err != nil {
		return nil, err
	}

	return createdAlert, nil
}
