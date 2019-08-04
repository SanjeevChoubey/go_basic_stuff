package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
)

// This method will be called for each new request from CLI
func (c *Client) NewRequest(db *sql.DB) error {

	client := &http.Client{}
	postData := make([]byte, 100)
	urlPath := os.Getenv("METAR_PATH")
	// Concatenate ICAO codes to UrlPath
	first := true
	for _, icao := range c.Icao {
		if first {
			urlPath = urlPath + icao
			first = false
			continue
		}
		urlPath = urlPath + "," + icao
	}
	//fmt.Println(urlPath)
	req, err := http.NewRequest("GET", urlPath, bytes.NewReader(postData))
	if err != nil {
		return err
	}
	req.Header.Add("X-API-Key", os.Getenv("API_KEY"))
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(bytes, &c.Result); err != nil {
		return err
	}

	count := fmt.Sprintf("%v", c.Result["results"])
	outString := fmt.Sprintf("%v", c.Result["data"])
	//fmt.Println("Return airport", count)
	//fmt.Println("Data ", outString)
	var airports []string
	// Parse data

	if count <= "1" {
		airports = append(airports, outString)
	} else {
		airports = c.ParseData(outString)
	}
	m := MetarRepository{}
	r := Report{}
	for _, v := range airports {
		// Update History table
		stationID := r.ParseStationID(v)
		m.AddWeather(db, v, stationID)
		//fmt.Println("Airport add", id, "and raw text is ", v)

		if r.IsWindSucks(v) || r.ParseWeather(v) {
			fmt.Println("Weather at ", stationID, "Sucks!!")
		} else {
			fmt.Println("Enjoy Weather at ", stationID)
		}
	}

	return nil
}

func (c *Client) ParseData(data string) []string {
	var airports []string
	var pos []int
	for i := 0; i < len(c.Icao); {
		v := strings.ToUpper(c.Icao[i])
		p := strings.Index(data, v)
		if p > 0 {
			pos = append(pos, p)
		}
		i++
	}
	sort.Ints(pos)
	var i int
	for i = 0; i < len(pos)-1; i++ {
		airports = append(airports, data[pos[i]:pos[i+1]])
	}
	airports = append(airports, data[pos[i]:])

	return airports
}
