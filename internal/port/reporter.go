package port

import "mac-cleaner/internal/domain"

type Reporter interface {
	GenerateReport() (*domain.Report, error)
}
