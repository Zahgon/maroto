package core

import (
	"errors"

	"github.com/johnfercher/maroto/v2/pkg/metrics"
)

var (
	ErrCannotMergeBytes = errors.New("cannot merge bytes")
	ErrCannotWriteFile  = errors.New("cannot write file")
)

type Pdf struct {
	bytes  []byte
	report *metrics.Report
}

// NewPDF is responsible to create a new instance of PDF.
func NewPDF(bytes []byte, report *metrics.Report) Document {
	_ = "STUB: not implemented"
	return *new(Document)
}

// GetBytes returns the PDF bytes.
func (p *Pdf) GetBytes() []byte {
	_ = "STUB: not implemented"

	// GetBase64 returns the PDF bytes in base64.
	return nil
}

func (p *Pdf) GetBase64() string { _ = "STUB: not implemented"; return "" }

// GetReport returns the metrics.Report.
func (p *Pdf) GetReport() *metrics.Report {
	_ = "STUB: not implemented"

	// Save saves the PDF in a file.
	return nil
}

func (p *Pdf) Save(file string) error { _ = "STUB: not implemented"; return nil }

// Merge merges the PDF with another PDF.
func (p *Pdf) Merge(bytes []byte) error { _ = "STUB: not implemented"; return nil }

func (p *Pdf) appendMetric(timeSpent *metrics.Time) { _ = "STUB: not implemented"; return }
