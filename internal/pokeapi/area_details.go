package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchAreaDetails(areaName string) (RespDetailedLocations, error) {
	areaURL := baseURL + "/location-area/" + areaName
	if val, ok := c.cache.Get(areaURL); ok {
		fmt.Println("Cache Exist")
		detailedLocationResp := RespDetailedLocations{}

		err := json.Unmarshal(val, &detailedLocationResp)
		if err != nil {
			return RespDetailedLocations{}, err
		}

		return detailedLocationResp, nil
	}

	req, err := http.NewRequest("GET", areaURL, nil)
	if err != nil {
		return RespDetailedLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDetailedLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDetailedLocations{}, err
	}

	detailedLocationResp := RespDetailedLocations{}
	err = json.Unmarshal(dat, &detailedLocationResp)
	if err != nil {
		return RespDetailedLocations{}, err
	}

	c.cache.Add(areaURL, dat)
	return detailedLocationResp, nil
}
