package resources

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resourceId, err := uuid.FromString(chi.URLParam(r, "id"))
		if err != nil {
			logrus.Errorf("parsing error : %s", err.Error())
			respondWithError(w, http.StatusUnprocessableEntity, fmt.Sprintf("cannot parse id (%s) as UUID", chi.URLParam(r, "id")))
			return
		}

		ctx := context.WithValue(r.Context(), "resourceId", resourceId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
