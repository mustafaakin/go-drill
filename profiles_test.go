package drill

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"os"
)

func TestDrillbit_Profiles(t *testing.T) {
	d, err := NewDrillbit("http://localhost:8047")
	assert.NoError(t, err)
	assert.NotNil(t, d)

	res, err := d.Profiles()
	assert.NoError(t, err)
	assert.NotNil(t, res)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(res)
}

func TestDrillbit_GetProfile(t *testing.T) {
	d, err := NewDrillbit("http://localhost:8047")
	assert.NoError(t, err)
	assert.NotNil(t, d)

	// get all profiles to obtain one of id
	res, err := d.Profiles()
	assert.NoError(t, err)
	assert.NotNil(t, res)

	result, err := d.GetProfile(res.FinishedQueries[0].QueryID)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(result)
}

func TestDrillbit_DeleteProfile(t *testing.T) {
	d, err := NewDrillbit("http://localhost:8047")
	assert.NoError(t, err)
	assert.NotNil(t, d)

	// get all profiles to obtain one of id
	res, err := d.Profiles()
	assert.NoError(t, err)
	assert.NotNil(t, res)

	result, err := d.DeleteProfile(res.FinishedQueries[0].QueryID)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(result)
}
