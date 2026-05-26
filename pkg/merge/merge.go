// Package merge implements PDF merge.
package merge

import (
	"errors"
	"io"
)

var ErrCannotMergePDFs = errors.New("cannot merge PDFs")

// Bytes merges PDFs from byte slices.
func Bytes(pdfs ...[]byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func mergePdfs(readers []io.ReadSeeker, writer io.Writer, dividerPage bool) error {
	_ = "STUB: not implemented"
	return nil
}
