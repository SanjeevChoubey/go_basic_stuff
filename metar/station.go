package main

import "regexp"

const stationIDPattern = `(\D{4})\s`

var stationIDRe *regexp.Regexp

func init() {
	stationIDRe = regexp.MustCompile(stationIDPattern)
}

func (r *Report) ParseStationID(raw string) string {
	if !stationIDRe.MatchString(raw) {
		return ""
	}
	// Return ICAO code
	return stationIDRe.FindAllStringSubmatch(raw, -1)[0][1]
}
