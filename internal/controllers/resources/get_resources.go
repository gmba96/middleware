package resources

import (
	"github.com/sirupsen/logrus"
	"middleware/config/internal/models"
	resources "middleware/config/internal/services/resources"
	"net/http"
)

// GetResources
// @Tags         resources
// @Summary      Get all resources
// @Description  Retrieve a list of all resources
// @Success      200            {array}   models.Resource
// @Failure      500            "Something went wrong"
// @Router       /resources [get]
func GetResources(w http.ResponseWriter, _ *http.Request) {
	// calling service
	resources, err := resources.GetAllResources()
	if err != nil {
		// logging error
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			respondWithError(w, customError.Code, customError.Message)
		} else {
			respondWithError(w, http.StatusInternalServerError, "Une erreur est survenue")
		}
		return
	}
	respondWithJSON(w, http.StatusOK, resources)
	return
}
