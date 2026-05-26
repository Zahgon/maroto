// Package metrics contains metrics models, constants and formatting.
package metrics

import (
	"errors"
)

var (
	ErrCannotWriteStringFile = errors.New("cannot write string to file")
	ErrCannotCreateFile      = errors.New("cannot create file")
)

type (
	// TimeScale is the time scale.
	TimeScale string
	// SizeScale is the size scale.
	SizeScale string
)

const (
	// Nano is the time scale in nanoseconds.
	Nano TimeScale = "ns"
	// Micro is the time scale in microseconds.
	Micro TimeScale = "μs"
	// Milli is the time scale in milliseconds.
	Milli TimeScale = "ms"
	// Byte is the size scale in bytes.
	Byte SizeScale = "b"
	// KiloByte is the size scale in kilobytes.
	KiloByte SizeScale = "Kb"
	// MegaByte is the size scale in megabytes.
	MegaByte SizeScale = "Mb"
	// GigaByte is the size scale in gigabytes.
	GigaByte SizeScale = "Gb"
)

// Time scales.
type Time struct {
	Value float64
	Scale TimeScale
}

// Normalize normalizes the time scale.
func (t *Time) Normalize() bool { _ = "STUB: not implemented"; return false }

// String returns the time formatted.
func (t *Time) String() string { _ = "STUB: not implemented"; return "" }

// Size scales.
type Size struct {
	Value float64
	Scale SizeScale
}

// Normalize normalizes the size scale.
func (t *Size) Normalize() bool { _ = "STUB: not implemented"; return false }

// String returns the size formatted.
func (t *Size) String() string { _ = "STUB: not implemented"; return "" }

// TimeMetric is a time metric.
type TimeMetric struct {
	Key   string
	Times []*Time
	Avg   *Time
}

// Normalize normalizes the time metric.
func (m *TimeMetric) Normalize() { _ = "STUB: not implemented"; return }

// String returns the time metric formatted.
func (m *TimeMetric) String() string { _ = "STUB: not implemented"; return "" }

func (m *TimeMetric) hasGreaterThan1000(times []*Time) bool {
	_ = "STUB: not implemented"
	return false
}

// SizeMetric is a size metric.
type SizeMetric struct {
	Key  string
	Size Size
}

// Normalize normalizes the size metric.
func (s *SizeMetric) Normalize() { _ = "STUB: not implemented"; return }

// String returns the size metric formatted.
func (s *SizeMetric) String() string { _ = "STUB: not implemented"; return "" }

// Report is a metrics report.
type Report struct {
	TimeMetrics []TimeMetric
	SizeMetric  SizeMetric
}

// Normalize normalizes the report.
func (r *Report) Normalize() *Report { _ = "STUB: not implemented"; return nil }

// String returns the report formatted.
func (r *Report) String() string { _ = "STUB: not implemented"; return "" }

// Save saves the report in a file.
func (r *Report) Save(file string) error { _ = "STUB: not implemented"; return nil }
