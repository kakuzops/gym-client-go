package model

import (
	"time"

	"github.com/google/uuid"
)

type Deployment struct {
	ID        uuid.UUID         `json:"id"`
	labels    map[string]string `json:"label"`
	Replicas  int               `json:replicas`
	Image     string            `json:"image"`
	Name      string            `json:"name"`
	Ports     []Port            `json: "ports"`
	CreatedAt time.Time         `json:"created_at"`
}

type Port struct {
	Name   string `json:"name"`
	Number uint   `json:"number"`
}
