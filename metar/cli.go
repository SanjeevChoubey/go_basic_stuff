package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

//CLI Information
func Info() {
	app.Name = "get_metar"
	app.Usage = "You can get wether report for any airport"
	app.Author = "Sanjeev Choubey"
	app.Version = "1.0.0"
}

// CLI commands, u can add any number of commands here
// In real time we can probably place these commands at better way.
// currently i am keeping in this array
func Commands() {
	app.Commands = []cli.Command{
		{
			Name:    "metar",
			Aliases: []string{"m"},
			Usage:   "You can pass space separated upto 20 ICAO codes.",
			Action: func(c *cli.Context) {
				client := &Client{}
				// User have to pass atleat one airport code
				if c.NArg() < 1 {
					fmt.Println("At least one Icao required")
					os.Exit(0)
				}

				args := c.Args().Get(0)

				//fmt.Println("Passed args", args)
				client.Icao = append(client.Icao, args)

				// If more ICAO code found, append those
				if c.Args().Present() {
					tail := c.Args().Tail()
					client.Icao = append(client.Icao, tail...)
				}
				// Validate ICAO for 4 Character
				for _, v := range client.Icao {
					match, _ := regexp.MatchString("[A-Z]{4}|[a-z]{4}", v)
					if match == false {
						fmt.Println("ICAO codes should be 4 character length")
						os.Exit(1)
					}
				}
				// Process new Request
				err := client.NewRequest(db)
				logFatal(err)
			},
		},
	}
}
