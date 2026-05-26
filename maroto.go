package maroto

import (
	"errors"

	"github.com/johnfercher/maroto/v2/internal/cache"

	"github.com/johnfercher/maroto/v2/pkg/core/entity"

	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/core"
)

var (
	ErrCannotGenerateInLowMemoryMode       = errors.New("an error has occurred while trying to generate PDFs in low memory mode")
	ErrCannotGenerateInParallelMode        = errors.New("an error has occurred while trying to generate PDFs concurrently")
	ErrFooterHeightIsGreaterThanUsefulArea = errors.New("footer height is greater than page useful area")
	ErrHeaderHeightIsGreaterThanUsefulArea = errors.New("header height is greater than page useful area")
)

type Maroto struct {
	config   *entity.Config
	provider core.Provider
	cache    cache.Cache

	// Building
	cell          entity.Cell
	pages         []core.Page
	rows          []core.Row
	header        []core.Row
	footer        []core.Row
	headerHeight  float64
	footerHeight  float64
	currentHeight float64
}

// GetCurrentConfig is responsible for returning the current settings from the file
func (m *Maroto) GetCurrentConfig() *entity.Config {
	_ = "STUB: not implemented"

	// New is responsible for create a new instance of core.Maroto.
	// It's optional to provide an *entity.Config with customizations
	// those customization are created by using the config.Builder.
	return nil
}

func New(cfgs ...*entity.Config) core.Maroto { _ = "STUB: not implemented"; return *new(core.Maroto) }

// AddPages is responsible for add pages directly in the document.
// By adding a page directly, the current cursor will reset and the
// new page will appear as the next. If the page provided have
// more rows than the maximum useful area of a page, maroto will split
// that page in more than one.
func (m *Maroto) AddPages(pages ...core.Page) { _ = "STUB: not implemented"; return }

// AddRows is responsible for add rows in the current document.
// By adding a row, if the row will extrapolate the useful area of a page,
// maroto will automatically add a new page. Maroto use the information of
// PageSize, PageMargin, FooterSize and HeaderSize to calculate the useful
// area of a page.
func (m *Maroto) AddRows(rows ...core.Row) {
	_ = "STUB: not implemented"

	// AddRow is responsible for add one row in the current document.
	// By adding a row, if the row will extrapolate the useful area of a page,
	// maroto will automatically add a new page. Maroto use the information of
	// PageSize, PageMargin, FooterSize and HeaderSize to calculate the useful
	// area of a page.
	return
}

func (m *Maroto) AddRow(rowHeight float64, cols ...core.Col) core.Row {
	_ = "STUB: not implemented"
	return *new(core.Row)
}

// AddAutoRow is responsible for adding a line with automatic height to the
// current document.
// The row height will be calculated based on its content.
func (m *Maroto) AddAutoRow(cols ...core.Col) core.Row {
	_ = "STUB: not implemented"
	return *new(core.Row)
}

// FitlnCurrentPage is responsible to validating whether a line fits on
// the current page.
func (m *Maroto) FitlnCurrentPage(heightNewLine float64) bool {
	_ = "STUB: not implemented"
	return false
}

// RegisterHeader is responsible to define a set of rows as a header
// of the document. The header will appear in every new page of the document.
// The header cannot occupy an area greater than the useful area of the page,
// it this case the method will return an error.
func (m *Maroto) RegisterHeader(rows ...core.Row) error { _ = "STUB: not implemented"; return nil }

// RegisterFooter is responsible to define a set of rows as a footer
// of the document. The footer will appear in every new page of the document.
// The footer cannot occupy an area greater than the useful area of the page,
// it this case the method will return an error.
func (m *Maroto) RegisterFooter(rows ...core.Row) error { _ = "STUB: not implemented"; return nil }

// Generate is responsible to compute the component tree created by
// the usage of all other Maroto methods, and generate the PDF document.
func (m *Maroto) Generate() (core.Document, error) {
	_ = "STUB: not implemented"
	return *new(core.Document), nil
}

// GetStructure is responsible for return the component tree, this is useful
// on unit tests cases.
func (m *Maroto) GetStructure() *node.Node[core.Structure] { _ = "STUB: not implemented"; return nil }

func (m *Maroto) addRows(rows ...core.Row) { _ = "STUB: not implemented"; return }

func (m *Maroto) addRow(r core.Row) { _ = "STUB: not implemented"; return }

// Row smaller than the remain space on page

// As row will extrapolate page, we will add empty space
// on the page to force a new page

// AddRows row on the new page

func (m *Maroto) addHeader() { _ = "STUB: not implemented"; return }

func (m *Maroto) fillPageToAddNew() { _ = "STUB: not implemented"; return }

// Truncate space to 9 decimal places to avoid rounding errors

func (m *Maroto) setConfig() { _ = "STUB: not implemented"; return }

func (m *Maroto) generate() (core.Document, error) {
	_ = "STUB: not implemented"
	return *new(core.Document), nil
}

func (m *Maroto) generateConcurrently() (core.Document, error) {
	_ = "STUB: not implemented"
	return *new(core.Document), nil
}

func (m *Maroto) generateLowMemory() (core.Document, error) {
	_ = "STUB: not implemented"
	return *new(core.Document), nil
}

func (m *Maroto) processPage(pages []core.Page) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Maroto) getRowsHeight(rows ...core.Row) float64 { _ = "STUB: not implemented"; return 0 }

func getConfig(configs ...*entity.Config) *entity.Config { _ = "STUB: not implemented"; return nil }

func getProvider(cache cache.Cache, cfg *entity.Config) core.Provider {
	_ = "STUB: not implemented"
	return *new(core.Provider)
}
