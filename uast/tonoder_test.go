package uast

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	fixtureDir = "fixtures"
)

func TestToNoderJava(t *testing.T) {
	require := require.New(t)

	f, err := getFixture("java_example_1.json")
	require.NoError(err)

	var c ToNoder = &BaseToNoder{
		InternalTypeKey: "internalClass",
		LineKey:         "line",
	}
	n, err := c.ToNode(f)
	require.NoError(err)
	require.NotNil(n)
}

func getFixture(name string) (map[string]interface{}, error) {
	path := filepath.Join(fixtureDir, name)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	d := json.NewDecoder(f)
	data := map[string]interface{}{}
	if err := d.Decode(&data); err != nil {
		_ = f.Close()
		return nil, err
	}

	if err := f.Close(); err != nil {
		return nil, err
	}

	return data, nil
}
