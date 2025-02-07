package resources

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/config/internal/models"
	resources "middleware/config/internal/services/resources"
	"net/http"
)

// GetCollections
// @Tags         collections
// @Summary      Get collections.
// @Description  Get collections.
// @Success      200            {array}  models.Collection
// @Failure      500             "Something went wrong"
// @Router       /collections [get]
func GetResources(w http.ResponseWriter, _ *http.Request) {
	// calling service
	resources, err := resources.GetAllResources()
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
	body, _ := json.Marshal(resources)
	_, _ = w.Write(body)
	return
}
