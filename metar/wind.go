package main

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	windDirectionSpeedPattern     = `(\d{3})(\d{2,3})KT`
	windDirectionSpeedGustPattern = `(\d{3})(\d{2,3})G(\d{2,3})KT`
	windCalmPattern               = `00000KT`
)

var (
	windDirectionSpeedRe     *regexp.Regexp
	windDirectionSpeedGustRe *regexp.Regexp
)

func init() {

	windDirectionSpeedRe = regexp.MustCompile(windDirectionSpeedPattern)
	windDirectionSpeedGustRe = regexp.MustCompile(windDirectionSpeedGustPattern)

}

func (r *Report) IsWindSucks(raw string) bool {
	// if no wind them its calm
	if strings.Contains(raw, "00000KT") {
		r.wind.IsCalm = true
	} else {
		if windDirectionSpeedGustRe.MatchString(raw) {
			matches := windDirectionSpeedGustRe.FindAllStringSubmatch(raw, -1)
			r.wind.Speed = matches[0][2]
			r.wind.Gust = matches[0][3]
		} else if windDirectionSpeedRe.MatchString(raw) {
			matches := windDirectionSpeedRe.FindAllStringSubmatch(raw, -1)
			r.wind.Speed = matches[0][2]
		}
	}
	speed, _ := strconv.Atoi(r.wind.Speed)
	// If Speed of wing is more than 30KT, then this weather sucks
	if speed > 30 {
		return true
	}
	return false
}
