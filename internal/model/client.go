package model

import (
	"time"
)

type Client struct {
	ID          int64     `json:"id"`
	ClientName  string    `json:"clientName"`
	Version     int       `json:"version"`
	Image       string    `json:"image"`
	CPU         string    `json:"cpu"`
	Memory      string    `json:"memory"`
	Priority    float64   `json:"priority"`
	NeedRestart bool      `json:"needRestart"`
	SpawnedAt   time.Time `json:"spawnedAt"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewClient(client Client) *Client {
	return &Client{
		ID:          client.ID,
		ClientName:  client.ClientName,
		Version:     client.Version,
		Image:       client.Image,
		CPU:         client.CPU,
		Memory:      client.Memory,
		Priority:    client.Priority,
		NeedRestart: client.NeedRestart,
		SpawnedAt:   time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
