package macos

import (
	"dusty/internal/domain"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

var pathsToClear []string
var homeDir string

type Cleaner struct {
	Logger *logrus.Logger
}

func init() {
	var err error
	homeDir, err = os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	pathsToClear = append(pathsToClear, filepath.Join(homeDir, "Library/Caches"))
	pathsToClear = append(pathsToClear, filepath.Join(homeDir, "Library/Logs"))
	pathsToClear = append(pathsToClear, filepath.Join(homeDir, ".Trash"))

}

func (c *Cleaner) SafeClean(safeClean bool) ([]domain.ItemReport, error) {
	results := make(chan *domain.ItemReport)
	errs := make(chan error, len(pathsToClear))

	var wg sync.WaitGroup

	for _, path := range pathsToClear {
		wg.Add(1)
		go func(p string) {
			c.Logger.Info("Cleaning: ", p)
			defer wg.Done()
			report, err := c.performSafeCleaning(path, safeClean)
			if err != nil {
				errs <- fmt.Errorf("error cleaning %s: %w", path, err)
			}
			results <- report
		}(path)
	}

	go func() {
		wg.Wait()
		close(results)
		close(errs)

	}()

	items := make([]domain.ItemReport, 0)

	for item := range results {
		items = append(items, *item)
	}

	return items, nil
}

func (c *Cleaner) performSafeCleaning(path string, safeClean bool) (*domain.ItemReport, error) {

	var totalFilesCount, totalDirCount int
	errors := make([]string, 0)
	var totalSize float64

	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			//c.Logger.Error(err)
		}

		if d.IsDir() {
			totalDirCount++
			c.Logger.Debug("Directory found", d.Name())
		} else {
			info, err := os.Stat(path)
			if err != nil {
				//c.Logger.Error("Error while reading file: ", path, err)
				errors = append(errors, "Error while reading file"+path)
			} else {
				totalFilesCount++
				totalSize += float64(info.Size())
			}

			if !safeClean {
				if rmErr := os.Remove(path); rmErr != nil {
					//c.Logger.Error("Failed to remove file", path, err)
				} else {
					c.Logger.Info("Removed file", path)
				}
			}

		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &domain.ItemReport{
		Path:   path,
		SizeMB: totalSize / 1024 / 1024,
		Files:  totalFilesCount,
		Dirs:   totalDirCount,
	}, nil

}

func (c *Cleaner) ComplexClean(path string) (*domain.Report, error) {
	return nil, nil
}

func (c *Cleaner) performComplexClean() (*domain.ItemReport, error) {
	return nil, nil
}
