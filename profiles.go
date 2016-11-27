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

func (d *Drillbit) Profiles() (*Profiles, error) {
	var res Profiles
	err := d.get("/profiles.json", &res)
	return &res, err
}
