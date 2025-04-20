package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type RiverData struct {
	SiteCode  string  `json:"site"`
	FlowCFS   float64 `json:"flow_cfs"`
	Timestamp string  `json:"timestamp"`
}

func FetchRiverData(site string) (*RiverData, error) {
	url := fmt.Sprintf("https://waterservices.usgs.gov/nwis/iv/?format=json&sites=%s&parameterCd=00060", site)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var raw map[string]interface{}
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, err
	}

	// Walk the structure manually
	timeSeries := raw["value"].(map[string]interface{})["timeSeries"].([]interface{})
	if len(timeSeries) == 0 {
		return nil, fmt.Errorf("no data found for site %s", site)
	}

	entry := timeSeries[0].(map[string]interface{})
	source := entry["sourceInfo"].(map[string]interface{})
	siteCode := source["siteCode"].([]interface{})[0].(map[string]interface{})["value"].(string)

	values := entry["values"].([]interface{})[0].(map[string]interface{})["value"].([]interface{})
	latest := values[len(values)-1].(map[string]interface{})
	flowStr := latest["value"].(string)
	timestamp := latest["dateTime"].(string)

	flowCFS, err := strconv.ParseFloat(flowStr, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid flow rate: %v", flowStr)
	}

	return &RiverData{
		SiteCode:  siteCode,
		FlowCFS:   flowCFS,
		Timestamp: timestamp,
	}, nil
}
