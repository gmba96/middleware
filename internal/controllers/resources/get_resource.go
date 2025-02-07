package resources

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/config/internal/models"
	resources "middleware/config/internal/services/resources"
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
	ctx := r.Context()
	resourceId, _ := ctx.Value("id").(uuid.UUID)

	resource, err := resources.GetResourceById(resourceId)
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
	body, _ := json.Marshal(resource)
	_, _ = w.Write(body)
	return
}
