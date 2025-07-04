package reporter

import (
	"dusty/internal/domain"
	"encoding/json"
	"time"
)

type JsonReport struct {
	Items     []domain.ItemReport
	StartTime time.Time
	EndTime   time.Time
}

func (j *JsonReport) GenerateReport() (*domain.Report, []byte) {
	r := &domain.Report{
		StartedAt:  j.StartTime,
		FinishedAt: j.EndTime,
		Duration:   j.EndTime.Sub(j.StartTime).Seconds(),
		Summary:    *CreateSummaryMetrics(j.Items),
		Items:      j.Items,
		System:     *CreateSystemMetadata(),
	}

	bArr, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	return r, bArr

}
