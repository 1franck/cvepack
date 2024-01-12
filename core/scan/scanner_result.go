package scan

import (
	"cvepack/core/ecosystem"
	"time"
)

type ScannerResult struct {
	Source    ecosystem.Source
	StartedAt time.Time
	EndedAt   time.Time
	Projects  []ecosystem.Project
}

func NewScannerResult(source ecosystem.Source) *ScannerResult {
	return &ScannerResult{Source: source, StartedAt: time.Now()}
}

func (s *ScannerResult) End() {
	s.EndedAt = time.Now()
}

func (s *ScannerResult) Duration() time.Duration {
	return s.EndedAt.Sub(s.StartedAt)
}
