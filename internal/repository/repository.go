package repository

import (
	"compile-and-run-sandbox/internal/sandbox"
	"time"
)

type Execution struct {
	ID string `gorm:"primarykey"`

	Source     string
	Output     string
	Status     sandbox.ContainerStatus
	TestStatus sandbox.ContainerTestStatus

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Repository struct {
}
