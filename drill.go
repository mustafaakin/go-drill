package drill

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Drillbit struct {
	endpoint string
}

// NewDrill connects to a given Drillbit by its REST API
func NewDrillbit(endpoint string) (*Drillbit, error) {
	d := Drillbit{
		endpoint: endpoint,
	}

	// Test connection
	return &d, nil
}

// NewDrillFromZK connects to Zookeeper and queries the drillbits
func NewDrillFromZK(zkpath string) (*Drillbit, error) {
	return nil, nil
}

func (d *Drillbit) request(urlPath string, method string, body interface{}, resp interface{}) error {
	path := d.endpoint + urlPath
	var res *http.Response
	var err error

	switch method {
	case "GET":
		res, err = http.Get(path)
	case "POST":
		if body != nil {
			b := new(bytes.Buffer)
			json.NewEncoder(b).Encode(body)
			res, err = http.Post(path, "application/json", b)
		}
	case "DELETE":
		req, err := http.NewRequest("DELETE", path, nil)
		if err != nil {
			return err
		}
		client := &http.Client{}
		res, err = client.Do(req)
	}

	if err != nil {
		return err
	}

	if res.Body == nil {
		return errors.New("The response body was empty")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		} else {
			return fmt.Errorf("Drill returned an error, Status: %d Message: \n%s", res.StatusCode, string(b))
		}
	}

	err = json.NewDecoder(res.Body).Decode(&resp)
	return err
}

// get makes a GET request
func (d *Drillbit) get(urlPath string, resp interface{}) (error) {
	return d.request(urlPath, "GET", nil, &resp)
}

// post makes a POST request
func (d *Drillbit) post(urlPath string, body interface{}, resp interface{}) error {
	return d.request(urlPath, "POST", body, &resp)
}

// delete makes a DELETE request
func (d *Drillbit) delete(urlPath string, resp interface{}) error {
	return d.request(urlPath, "DELETE", nil, &resp)
}
