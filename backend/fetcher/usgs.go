package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetRiverData(site string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://waterservices.usgs.gov/nwis/iv/?format=json&sites=%s&parameterCd=00060,00065", site)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}
