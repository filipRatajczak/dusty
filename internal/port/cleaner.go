package port

import "dusty/internal/domain"

type Cleaner interface {
	SafeClean(path string) (*domain.Report, error)
	ComplexClean(path string) (*domain.Report, error)
}
