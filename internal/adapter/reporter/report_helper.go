package reporter

import (
	"dusty/internal/domain"
	"github.com/shirou/gopsutil/v3/host"
	"os/user"
)

func CreateSummaryMetrics(report []domain.ItemReport) *domain.SummaryMetrics {

	totalFiles := 0
	totalDirs := 0
	totalSize := 0.0
	totalItems := 0

	for _, i := range report {
		totalFiles += i.Files
		totalDirs += i.Dirs
		totalSize += i.SizeMB
		totalItems += 1
	}

	return &domain.SummaryMetrics{
		TotalSizeMB:      totalSize,
		TotalFiles:       totalFiles,
		TotalDirectories: totalDirs,
		ItemsCleaned:     totalItems,
	}
}

func CreateSystemMetadata() *domain.SystemMetadata {

	info, err := host.Info()
	if err != nil {
		panic(err)
	}

	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	return &domain.SystemMetadata{
		Hostname:   info.Hostname,
		OSVersion:  info.OS,
		User:       currentUser.Username,
		AppVersion: "0.0.1",
	}

}
