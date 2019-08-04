package main

import (
	"regexp"
)

const weatherpattern = `(\+|-|VC)?((MI|PR|BC|DR|BL|SH|TS|FZ)|(DZ|RA|SN|SG|IC|PL|GR|GS|UP)|(BR|FG|FU|VA|DU|SA|HZ|PY)|(PO|SQ|FC|SS|DS))`

var weatherRe *regexp.Regexp

func init() {
	weatherRe = regexp.MustCompile(weatherpattern)
}

func (r *Report) ParseWeather(raw string) bool {
	if !weatherRe.MatchString(raw) {
		return false
	}

	// We are handling only these Metar code hance kept it in an array
	meterCodes := []string{
		"FG", "GR", "SN", "FZRA", "FZDZ", "RA", "TS", "PL", "DZ", "VCTS",
	}

	matches := weatherRe.FindAllStringSubmatch(raw, -1)

	// Keep in array for all weather report
	for _, match := range matches {
		for _, n := range meterCodes {
			if match[0] == n {
				return true
			}
		}
	}
	return false
}
