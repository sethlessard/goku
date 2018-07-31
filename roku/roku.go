package roku

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

// Roku struct
type Roku struct {
	APIAddress string `json:"api_address"`
	USN        string `json:"usn"`
}

type AppDictionary struct {
	XMLName xml.Name `xml:"apps"`
	Apps    []struct {
		Name string `xml:",chardata"`
		ID   string `xml:"id,attr"`
	} `xml:"app"`
}

type ActiveApp struct {
	XMLName xml.Name `xml:"active-app"`
	App     []struct {
		Name    string `xml:",chardata"`
		ID      string `xml:"id,attr"`
		Version string `xml:"version,attr"`
	} `xml:"app"`
}

func (r Roku) GetApps() (AppDictionary, error) {
	resp, err := http.Get(fmt.Sprintf("%squery/apps", r.APIAddress))
	var dict AppDictionary
	if err != nil {
		return dict, err
	}

	defer resp.Body.Close()

	if err := xml.NewDecoder(resp.Body).Decode(&dict); err != nil {
		return dict, err
	}

	return dict, nil
}

func (r Roku) GetActiveApp() (ActiveApp, error) {
	res, err := http.Get(fmt.Sprintf("%squery/active-app", r.APIAddress))
	var activeApp ActiveApp
	if err != nil {
		return activeApp, err
	}

	defer res.Body.Close()

	if err := xml.NewDecoder(res.Body).Decode(&activeApp); err != nil {
		return activeApp, err
	}

	return activeApp, nil
}
