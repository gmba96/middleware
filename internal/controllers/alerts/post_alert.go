package alerts

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/config/internal/models"
	"middleware/config/internal/services/alerts"
	"net/http"
)

// CreateAlert
// @Tags         alerts
// @Summary      Créer une alerte
// @Description  Créer une alerte
// @Param        body body      CreateAlertRequest  true  "Alert data"
// @Success      201  {object}  models.Alert
// @Failure      400  "Invalid request body"
// @Failure      500  "Something went wrong"
// @Router       /alerts [post]
func PostAlert(w http.ResponseWriter, r *http.Request) {
	var req CreateAlertRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Requête invalide")
		return
	}

	createdAlert, err := alerts.CreateAlert(req.Email, req.All, req.ResourceId)
	if err != nil {
		logrus.Errorf("error: %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			respondWithError(w, customError.Code, customError.Message)
		} else {
			respondWithError(w, http.StatusInternalServerError, "Une erreur est survenue")
		}
		return
	}

	respondWithJSON(w, http.StatusCreated, createdAlert)
}

type CreateAlertRequest struct {
	Email      string     `json:"email"`
	All        bool       `json:"all"`
	ResourceId *uuid.UUID `json:"resource_id"`
}
