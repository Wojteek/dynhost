package app

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

// Data - the structure of file with the data
type Data struct {
	LastIP    string    `json:"last_ip"`
	PrevIP    string    `json:"prev_ip"`
	ChangedAt time.Time `json:"changed_at"`
}

// GetData returns a data from file
func GetData(dataPath string) (Data, error) {
	content, err := ioutil.ReadFile(dataPath)

	if err != nil {
		return Data{}, err
	}

	var d Data

	if err := json.Unmarshal(content, &d); err != nil {
		return Data{}, err
	}

	return d, nil
}

// SaveData saves the data into the file
func (d *Data) SaveData(dataPath string) (*Data, error) {
	jsonData, err := json.Marshal(d)

	if err != nil {
		return d, err
	}

	if err := ioutil.WriteFile(dataPath, jsonData, 0644); err != nil {
		return d, err
	}

	return d, nil
}
