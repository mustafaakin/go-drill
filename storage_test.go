package drill

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"os"
)

func TestDrillbit_Storage(t *testing.T) {
	d, err := NewDrillbit("http://localhost:8047")
	assert.NoError(t, err)
	assert.NotNil(t, d)

	res, err := d.Storage()
	assert.NoError(t, err)
	assert.NotNil(t, res)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(res)
}

func TestDrillbit_GetStorage(t *testing.T) {
	d, err := NewDrillbit("http://localhost:8047")
	assert.NoError(t, err)
	assert.NotNil(t, d)

	res, err := d.GetStorage("mongo")
	assert.NoError(t, err)
	assert.NotNil(t, res)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(res)
}

func TestDrillbit_EnableStorage(t *testing.T) {
	d, err := NewDrillbit("http://localhost:8047")
	assert.NoError(t, err)
	assert.NotNil(t, d)

	res, err := d.EnableStorage("mongo")
	assert.NoError(t, err)
	assert.NotNil(t, res)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(res)
}

func TestDrillbit_DisableStorage(t *testing.T) {
	d, err := NewDrillbit("http://localhost:8047")
	assert.NoError(t, err)
	assert.NotNil(t, d)

	res, err := d.DisableStorage("mongo")
	assert.NoError(t, err)
	assert.NotNil(t, res)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(res)
}

func TestDrillbit_CreateStorage(t *testing.T) {
	d, err := NewDrillbit("http://localhost:8047")
	assert.NoError(t, err)
	assert.NotNil(t, d)

	var storage Storage
	storage.Name = "myplugin"
	storage.Config = map[string]interface{}{
		"type": "file",
		"enabled": false,
		"connection" : "file:///",
		"workspaces" : map[string]interface{}{"root" : map[string]interface{}{
			"location" : "/",
			"writable" : false,
			"defaultInputFormat" : nil,
		}},
		"formats" : nil,
	}

	res, err := d.CreateStorage("mongo",storage)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(res)
}

func TestDrillbit_DeleteStorage(t *testing.T) {
	d, err := NewDrillbit("http://localhost:8047")
	assert.NoError(t, err)
	assert.NotNil(t, d)

	res, err := d.DeleteStorage("myplugin")
	assert.NoError(t, err)
	assert.NotNil(t, res)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(res)
}
