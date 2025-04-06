package alerts

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/config/internal/models"
	alerts "middleware/config/internal/services/alerts"
	"net/http"
)

// GetAlert
// @Tags         alerts
// @Summary      Récupère une alerte
// @Description  Récupère une alerte par son ID
// @Param        id           	path      string  true  "Alert UUID formatted ID"
// @Success      200            {object}  models.Alert
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /alerts/{id} [get]
func GetAlert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	alertId, _ := ctx.Value("id").(uuid.UUID)

	alert, err := alerts.GetAlertById(alertId)
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(alert)
	_, _ = w.Write(body)
	return
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	body, _ := json.Marshal(&models.CustomError{Message: message, Code: code})
	_, _ = w.Write(body)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.WriteHeader(code)
	body, _ := json.Marshal(payload)
	_, _ = w.Write(body)
}
