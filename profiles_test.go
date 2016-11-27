package drill

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"os"
)

func TestDrillbit_Profiles(t *testing.T) {
	d, err := NewDrillbit("http://172.16.3.17:8047")
	assert.NoError(t, err)
	assert.NotNil(t, d)

	res, err := d.Profiles()
	assert.NoError(t, err)
	assert.NotNil(t, res)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(res)
}
