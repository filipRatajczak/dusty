package port

import "dusty/internal/domain"

type Reporter interface {
	GenerateReport() (*domain.Report, error)
}
