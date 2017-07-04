package drill

type Storage struct {
	Name   string `json:"name"`
	Config interface{} `json:"config"`
}

type ResultState struct {
	Result string `json:"result"`
}

// Storage returns all storage as Storage format array
func (d *Drillbit) Storage() (*[]Storage, error) {
	res := make([]Storage, 0)
	err := d.get("/storage.json", &res)
	return &res, err
}

// GetStorage given a storage name, returns the storage as Storage format
func (d *Drillbit) GetStorage(name string) (*Storage, error) {
	path := "/storage/" + name + ".json"
	var res Storage
	err := d.get(path, &res)
	return &res, err
}

// EnableStorage given a storage name, returns the result as ResultState format
func (d *Drillbit) EnableStorage(name string) (*ResultState, error) {
	path := "/storage/" + name + "/enable/true"
	var res ResultState
	err := d.get(path, &res)
	return &res, err
}

// DisableStorage given a storage name, returns the result as ResultState format
func (d *Drillbit) DisableStorage(name string) (*ResultState, error) {
	path := "/storage/" + name + "/enable/false"
	var res ResultState
	err := d.get(path, &res)
	return &res, err
}

// CreateStorage given a storage name and Storage instance, returns the result as ResultState format
func (d *Drillbit) CreateStorage(name string, storage Storage) (*ResultState, error) {
	path := "/storage/" + name + ".json"
	var res ResultState
	err := d.post(path, storage, &res)
	return &res, err
}

// DeleteStorage given a storage name, returns the result as ResultState format
func (d *Drillbit) DeleteStorage(name string) (*ResultState, error) {
	path := "/storage/" + name + ".json"
	var res ResultState
	err := d.delete(path, &res)
	return &res, err
}
