package resources

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/config/internal/models"
	"middleware/config/internal/services/resources"
	"net/http"
)

// GetResource
// @Tags         resources
// @Summary      Get a collection.
// @Description  Get a collection.
// @Param        id           	path      string  true  "Collection UUID formatted ID"
// @Success      200            {object}  models.Resource
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /resources/{id} [get]
func GetResource(w http.ResponseWriter, r *http.Request) {
	resourceId, ok := r.Context().Value("resourceId").(uuid.UUID)
	if !ok {
		respondWithError(w, http.StatusBadRequest, "ID de resource invalide")
		return
	}

	resource, err := resources.GetResourceById(resourceId)
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		if customError, isCustom := err.(*models.CustomError); isCustom {
			respondWithError(w, customError.Code, customError.Message)
		} else {
			respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		}
		return
	}

	respondWithJSON(w, http.StatusOK, resource)
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
