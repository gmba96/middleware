package resources

import (
	"encoding/json"
	_ "encoding/json"
	_ "github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	_ "github.com/sirupsen/logrus"
	"middleware/config/internal/models"
	"middleware/config/internal/services/resources"
	"net/http"
)

// PostResource
// @Tags         resources
// @Summary      Create a new resource
// @Description  Create a new resource
// @Param        resource      body      models.Resource  true  "Resource data"
// @Success      201           {object}  models.Resource
// @Failure      400           "Invalid request body"
// @Failure      500           "Something went wrong"
// @Router       /resources [post]
func PostResource(w http.ResponseWriter, r *http.Request) {
	var req CreateResourceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logrus.Errorf("error decoding request: %s", err.Error())
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	newResource, err := resources.CreateResource(req.UcaId, req.Name)
	if err != nil {
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			respondWithError(w, customError.Code, customError.Message)
		} else {
			respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		}
		return
	}
	respondWithJSON(w, http.StatusCreated, newResource)
}

// CreateResourceRequest request struct
type CreateResourceRequest struct {
	Name  string `json:"name"`
	UcaId int    `json:"uca_id"`
}
