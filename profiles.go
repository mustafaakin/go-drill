package drill

type Profiles struct {
	RunningQueries  []Profile
	FinishedQueries []Profile
}

type Profile struct {
	QueryID  string
	Time     string
	Location string
	Foreman  string
	Query    string
	State    string
	User     string
}

// Profiles returns all running and finished queries as Profile format array
func (d *Drillbit) Profiles() (*Profiles, error) {
	var res Profiles
	err := d.get("/profiles.json", &res)
	return &res, err
}

// GetProfile given a query id, returns query profile
func (d *Drillbit) GetProfile(queryId string) (interface{}, error) {
	var res interface{}
	path := "/profiles/" + queryId + ".json"
	err := d.get(path, &res)
	return &res, err
}

// DeleteProfile given a query id, returns success if the query is running, otherwise
// it returns an error as a string so it cannot be decoded
func (d *Drillbit) DeleteProfile(queryId string) (interface{}, error) {
	var res ResultState
	path := "/profiles/cancel/" + queryId
	err := d.get(path, &res)
	return &res, err
}