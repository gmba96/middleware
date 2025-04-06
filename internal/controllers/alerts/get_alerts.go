package alerts

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/config/internal/models"
	alerts "middleware/config/internal/services/alerts"
	"net/http"
)

// Get Alerts
// @Tags         alerts
// @Summary      Récupère toutes les alertes
// @Description  Récupère toutes les alertes
// @Success      200            {array}  models.Collection
// @Failure      500             "Something went wrong"
// @Router       /alerts [get]
func GetAlerts(w http.ResponseWriter, _ *http.Request) {
	// calling service
	alerts, err := alerts.GetAllAlerts()
	if err != nil {
		// logging error
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			// writing http code in header
			w.WriteHeader(customError.Code)
			// writing error message in body
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(alerts)
	_, _ = w.Write(body)
	return
}
