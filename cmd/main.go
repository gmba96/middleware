package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/config/internal/controllers/alerts"
	"middleware/config/internal/controllers/resources"
	"middleware/config/internal/helpers"
	_ "middleware/config/internal/models"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	// Routes pour resources
	r.Route("/resources", func(r chi.Router) {
		r.Get("/", resources.GetResources)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(resources.Ctx)
			r.Get("/", resources.GetResource)
		})
	})

	// Routes pour alerts
	r.Route("/alerts", func(r chi.Router) {
		r.Get("/", alerts.GetAlerts)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(alerts.Ctx)
			r.Get("/", alerts.GetAlert)
		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8080", r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}

	schemes := []string{
		`CREATE TABLE IF NOT EXISTS resources (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			ucaId INT NOT NULL,
			name VARCHAR(255) NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS alerts (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			email VARCHAR(255) NOT NULL,
			all_boolean BOOLEAN NOT NULL DEFAULT FALSE,
			resourceId VARCHAR(255) NOT NULL
		);`,
	}

	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table! Error was: " + err.Error())
		}
	}

	helpers.CloseDB(db)
}
