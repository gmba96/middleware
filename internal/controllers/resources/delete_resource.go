package resources

import (
	"github.com/gofrs/uuid"
	"middleware/config/internal/models"
	"middleware/config/internal/services/resources"
	"net/http"
)

// DeleteResource
// @Tags resources
// @Summary Supprime une resource
// @Description Supprime la resource identifiée par son ID
// @Produce json
// @Success 204 "Aucun contenu"
// @Failure 400 "ID de resource invalide"
// @Failure 404 "Resource non trouvée"
// @Failure 500 "Une erreur est survenue"
// @Router /resources/{id} [delete]
func DeleteResource(w http.ResponseWriter, r *http.Request) {
	resourceId, ok := r.Context().Value("resourceId").(uuid.UUID)
	if !ok {
		respondWithError(w, http.StatusBadRequest, "ID de resource invalide")
		return
	}

	err := resources.RemoveResource(resourceId)
	if err != nil {
		if customError, isCustom := err.(*models.CustomError); isCustom {
			respondWithError(w, customError.Code, customError.Message)
		} else {
			respondWithError(w, http.StatusInternalServerError, "Une erreur est survenue")
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
