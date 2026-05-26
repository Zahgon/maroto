// nolint:errchkjson // not needed
package test

import (
	"errors"
	"testing"

	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/core"
)

var (
	ErrCannotReadDir       = errors.New("cannot read directory")
	ErrCannotReadFile      = errors.New("cannot read file")
	ErrCannotUnmarshallYML = errors.New("cannot unmarshall yaml")
	ErrMarotoYMLNotFound   = errors.New("found go.mod but not .maroto.yml")
)

var (
	marotoFile      = ".maroto.yml"
	goModFile       = "go.mod"
	configSingleton *Config
)

type Node struct {
	Value   any            `json:"value,omitempty"`
	Type    string         `json:"type"`
	Details map[string]any `json:"details,omitempty"`
	Nodes   []*Node        `json:"nodes,omitempty"`
}

// MarotoTest is the unit test instance.
type MarotoTest struct {
	t    *testing.T
	node *node.Node[core.Structure]
}

// New creates the MarotoTest instance to unit tests.
func New(t *testing.T) *MarotoTest { _ = "STUB: not implemented"; return nil }

// Assert validates if the structure is the same as defined by Equals method.
func (m *MarotoTest) Assert(structure *node.Node[core.Structure]) *MarotoTest {
	_ = "STUB: not implemented"
	return nil
}

// Equals defines which file will be loaded to do the comparison.
func (m *MarotoTest) Equals(file string) *MarotoTest { _ = "STUB: not implemented"; return nil }

// Save is an auxiliary method to update the file to be asserted.
func (m *MarotoTest) Save(file string) *MarotoTest { _ = "STUB: not implemented"; return nil }

func (m *MarotoTest) buildNode(node *node.Node[core.Structure]) *Node {
	_ = "STUB: not implemented"
	return nil
}

func getMarotoConfigFilePath() (string, error) { _ = "STUB: not implemented"; return "", nil }

func loadMarotoConfigFile(path string) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

func getMarotoConfigFilePathRecursive(path string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func hasFileInPath(file string, path string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func getParentDir(path string) string { _ = "STUB: not implemented"; return "" }
