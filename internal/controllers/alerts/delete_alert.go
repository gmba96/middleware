package alerts

import (
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/config/internal/models"
	"middleware/config/internal/services/alerts"
	"net/http"
)

// DeleteAlert
// @Tags         alerts
// @Summary      Delete an alert.
// @Description  Delete an alert.
// @Param        id   path      string  true  "Alert UUID formatted ID"
// @Success      204  "Alert deleted successfully"
// @Failure      404  "Alert not found"
// @Failure      500  "Something went wrong"
// @Router       /alerts/{id} [delete]
func DeleteAlert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	alertId, _ := ctx.Value("alertId").(uuid.UUID)

	err := alerts.DeleteAlert(alertId)
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

	respondWithJSON(w, http.StatusOK, "")
}
