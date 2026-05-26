package maroto

import (
	"github.com/johnfercher/go-tree/node"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/metrics"
)

type MetricsDecorator struct {
	addRowsTime    []*metrics.Time
	addRowTime     []*metrics.Time
	addAutoRowTime []*metrics.Time
	addPageTime    []*metrics.Time
	headerTime     *metrics.Time
	footerTime     *metrics.Time
	generateTime   *metrics.Time
	structureTime  *metrics.Time
	inner          core.Maroto
}

// NewMetricsDecorator is responsible to create the metrics decorator
// for the maroto instance.
func NewMetricsDecorator(inner core.Maroto) core.Maroto {
	_ = "STUB: not implemented"
	return *new(core.Maroto)
}

// FitlnCurrentPage decoratess the FitlnCurrentPage method of maroto instance.
func (m *MetricsDecorator) FitlnCurrentPage(heightNewLine float64) bool {
	_ = "STUB: not implemented"
	return false
}

// GetCurrentConfig decorates the GetCurrentConfig method of maroto instance.
func (m *MetricsDecorator) GetCurrentConfig() *entity.Config { _ = "STUB: not implemented"; return nil }

// Generate decorates the Generate method of maroto instance.
func (m *MetricsDecorator) Generate() (core.Document, error) {
	_ = "STUB: not implemented"
	return *new(core.Document), nil
}

// AddPages decorates the AddPages method of maroto instance.
func (m *MetricsDecorator) AddPages(pages ...core.Page) { _ = "STUB: not implemented"; return }

// AddRows decorates the AddRows method of maroto instance.
func (m *MetricsDecorator) AddRows(rows ...core.Row) { _ = "STUB: not implemented"; return }

// AddRow decorates the AddRow method of maroto instance.
func (m *MetricsDecorator) AddRow(rowHeight float64, cols ...core.Col) core.Row {
	_ = "STUB: not implemented"
	return *new(core.Row)
}

// AddRow decorates the AddRow method of maroto instance.
func (m *MetricsDecorator) AddAutoRow(cols ...core.Col) core.Row {
	_ = "STUB: not implemented"
	return *new(core.Row)
}

// RegisterHeader decorates the RegisterHeader method of maroto instance.
func (m *MetricsDecorator) RegisterHeader(rows ...core.Row) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterFooter decorates the RegisterFooter method of maroto instance.
func (m *MetricsDecorator) RegisterFooter(rows ...core.Row) error {
	_ = "STUB: not implemented"
	return nil
}

// GetStructure decorates the GetStructure method of maroto instance.
func (m *MetricsDecorator) GetStructure() *node.Node[core.Structure] {
	_ = "STUB: not implemented"
	return nil
}

func (m *MetricsDecorator) buildMetrics(bytesSize int) *metrics.Report {
	_ = "STUB: not implemented"
	return nil
}

func (m *MetricsDecorator) getAVG(times []*metrics.Time) *metrics.Time {
	_ = "STUB: not implemented"
	return nil
}
