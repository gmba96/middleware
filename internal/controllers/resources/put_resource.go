package resources

import (
	"encoding/json"
	"middleware/config/internal/services/resources"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/config/internal/models"
)

// PutResource
// @Tags         resources
// @Summary      Update a resource
// @Description  Update a specific resource by ID
// @Param        id            path      string          true  "Resource UUID formatted ID"
// @Param        resource      body      models.Resource  true  "Resource data"
// @Success      200           {object}  models.Resource
// @Failure      400           "Invalid request body"
// @Failure      422           "Cannot parse id"
// @Failure      500           "Something went wrong"
// @Router       /resources/{id} [put]
func PutResource(w http.ResponseWriter, r *http.Request) {
	// Récupération de l'ID de la resource dans le contexte
	resourceId, ok := r.Context().Value("resourceId").(uuid.UUID)
	if !ok {
		respondWithError(w, http.StatusBadRequest, "ID de resource invalide")
		return
	}

	// Décodage du corps de la requête
	var req UpdateResourceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logrus.Errorf("Erreur de décodage de la requête: %s", err.Error())
		respondWithError(w, http.StatusBadRequest, "Corps de requête invalide")
		return
	}

	// Appel de la couche service pour mettre à jour la resource
	updatedResource, err := resources.UpdateResource(resourceId, req.UcaId, req.Name)
	if err != nil {
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			respondWithError(w, customError.Code, customError.Message)
		} else {
			respondWithError(w, http.StatusInternalServerError, "Une erreur est survenue")
		}
		return
	}

	respondWithJSON(w, http.StatusOK, updatedResource)
}

// UpdateResourceRequest request struct
type UpdateResourceRequest struct {
	UcaId int    `json:"ucaId"`
	Name  string `json:"name"`
}
