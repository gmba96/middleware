package alerts

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/config/internal/models"
	"middleware/config/internal/services/alerts"
	"net/http"
)

// UpdateAlert
// @Tags         alerts
// @Summary      Met à jour une alerte
// @Description  Met à jour une alerte par son ID
// @Param        id   path      string  true  "Alert UUID formatted ID"
// @Param        body body      models.Alert  true  "Alert data"
// @Success      200  {object}  models.Alert
// @Failure      400  "Invalid request body"
// @Failure      404  "Alert not found"
// @Failure      500  "Something went wrong"
// @Router       /alerts/{id} [put]
func PutAlert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	alertId, _ := ctx.Value("alertId").(uuid.UUID)

	var req UpdateAlertRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Requête invalide")
		return
	}

	updatedAlert, err := alerts.UpdateAlert(alertId, req.Email, req.All, req.ResourceId)
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

	respondWithJSON(w, http.StatusOK, updatedAlert)
}

// UpdateAlertRequest struct request
type UpdateAlertRequest struct {
	Email      string     `json:"email"`
	All        bool       `json:"all"`
	ResourceId *uuid.UUID `json:"resource_id"`
}
