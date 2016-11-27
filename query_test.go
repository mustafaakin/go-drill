package drill

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDrillbit_RawQuery(t *testing.T) {
	d, err := NewDrillbit("http://172.16.3.17:8047")
	assert.NoError(t, err)
	assert.NotNil(t, d)

	res, err := d.RawQuery("show files in dfs")
	assert.NoError(t, err)
	assert.NotNil(t, res)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(res)
}

func TestDrillbit_Query(t *testing.T) {
	d, err := NewDrillbit("http://172.16.3.17:8047")
	assert.NoError(t, err)
	assert.NotNil(t, d)

	type fileQuery struct {
		Name             string
		IsDirectory      string
		IsFile           bool   `json:",string"`
		Length           uint64 `json:",string"`
		Owner            string
		Group            string
		Permissions      string
		AccessTime       string
		ModificationTime string
	}

	res := make([]fileQuery, 0)

	err = d.Query(&res, "show files in dfs")
	assert.NoError(t, err)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(res)
}
