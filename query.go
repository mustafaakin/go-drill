package drill

import (
    "bytes"
    "encoding/json"
)

// QueryResponse is a https://drill.apache.org/docs/rest-api/#query
type QueryResponse struct {
	Columns []string
	Rows    []interface{}
}

// RawQuery given an SQL, returns the official QueryResponse format
func (d *Drillbit) RawQuery(sql string) (*QueryResponse, error) {
	body := make(map[string]string)
	body["queryType"] = "SQL"
	body["query"] = sql

	var resp QueryResponse
	err := d.post("/query.json", body, &resp)
	return &resp, err
}

// Query is an extension to RawQuery method, which assumes the response is given object, so that we can assume that
// we can convert to given interface easily so that the caller can use it in a type-safe like manner.
func (d *Drillbit) Query(v interface{}, sql string) error {
    res, err := d.RawQuery(sql)
    if err != nil {
        return err
    }

    // TODO: Quick hack here, re-encoding already received data, can we make a better solution?
    b := new(bytes.Buffer)
    err = json.NewEncoder(b).Encode(res.Rows)
    if err != nil {
        return err
    }

    // Now try to decode it to given object
    err = json.NewDecoder(b).Decode(v)
    if err != nil {
        return err
    }
    return nil
}