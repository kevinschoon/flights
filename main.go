package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jawher/mow.cli"
)

func main() {
	app := cli.App("flights", "✈️:✈️ ✈️ ✈️ ✈️ ")
	app.Command("airports", "list all possible airports", func(cmd *cli.Cmd) {
		var (
			asJson = cmd.BoolOpt("json", false, "use json encoding")
		)
		cmd.Action = func() {
			airports := Ordered(Airports())
			if *asJson {
				maybe(json.NewEncoder(os.Stdout).Encode(airports))
			} else {
				for _, airport := range airports {
					fmt.Println(airport)
				}
			}
		}
	})

	app.Command("route", "find all possible routes", func(cmd *cli.Cmd) {
		cmd.Spec = "[OPTIONS] ARRIVAL DESTINATION"
		var (
			asJson        = cmd.BoolOpt("json", false, "use json encoding")
			threshold     = cmd.IntOpt("threshold", 100, "airport distance threshold")
			arrivalCode   = cmd.StringArg("ARRIVAL", "", "starting airport")
			departureCode = cmd.StringArg("DESTINATION", "", "ending airport")
		)
		cmd.Action = func() {
			airports := Airports()
			if _, ok := airports[*arrivalCode]; !ok {
				maybe(fmt.Errorf("bad arrival code: %s", *arrivalCode))
			}
			if _, ok := airports[*departureCode]; !ok {
				maybe(fmt.Errorf("bad departure code: %s", *departureCode))
			}
			departures, arrivals := FindRoutes(float64(*threshold), airports, airports[*arrivalCode], airports[*departureCode])
			if *asJson {
				data := [][]Airport{departures, arrivals}
				maybe(json.NewEncoder(os.Stdout).Encode(data))
			} else {
				fmt.Println("Departures: ")
				for _, departure := range departures {
					fmt.Println(departure)
				}
				fmt.Println("Arrivals: ")
				for _, arrival := range arrivals {
					fmt.Println(arrival)
				}
			}
		}
	})
	maybe(app.Run(os.Args))
}
