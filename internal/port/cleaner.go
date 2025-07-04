package port

import "mac-cleaner/internal/domain"

type Cleaner interface {
	SafeClean(path string) (*domain.Report, error)
	ComplexClean(path string) (*domain.Report, error)
}
