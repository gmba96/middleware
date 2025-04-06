package models

import (
	"github.com/gofrs/uuid"
)

type Resource struct {
	Id    *uuid.UUID `json:"id"`
	UcaId int        `json:"uca_id"`
	Name  string     `json:"name"`
}

type Alert struct {
	Id         *uuid.UUID `json:"id"`
	Email      string     `json:"email"`
	All        bool       `json:"all"`
	ResourceId *uuid.UUID `json:"resource_id"`
}
