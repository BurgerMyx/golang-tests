package main

import (
	"fmt"

	"github.com/Altarrel/goroyale"
	"github.com/kr/pretty"
)

var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6OTkwLCJpZGVuIjoiNDU1NzEzMjA3NDU4MjAxNjAyIiwibWQiOnt9LCJ0cyI6MTUyOTk0Mzg4OTk2MX0.HQ3w4_wbC0rhkrkFJaaF9zUfATaIQawG2522MLNHFeQ"

func main() {
	c, err := goroyale.New(token, 0) // 0 will use the default request timeout of 10 seconds
	if err != nil {
		fmt.Println(err)
		return
	}

	ver, err := c.APIVersion()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("API Version:", ver)
	}

	params := map[string][]string{
		"joinable": {"1"},
	}



	tourney, err := c.OpenTournaments(params)
	//p, err := c.Player("2PLLY2RU", params)
	if err != nil {
		fmt.Println(err)
	} else {
		// will just print "Name:" as p.Name is "" because it was excluded
		// more info about this at https://docs.royaleapi.com/#/field_filter
		for i:=0;i<len(tourney);i++{
			if(tourney[i].CurrentPlayers >= tourney[i].MaxPlayers*3/4){
				continue
			}else{
				pretty.Printf("%s | %s\n", tourney[i].Name, tourney[i].Tag)
				pretty.Printf("Curr: %d | Max: %d\n", tourney[i].CurrentPlayers, tourney[i].MaxPlayers)
			}
		}
	}

	params2:=map[string][]string{}

	popplayer, err := c.PopularPlayers(params2)
	//p, err := c.Player("2PLLY2RU", params)
	if err != nil {
		fmt.Println(err)
	} else {
		// will just print "Name:" as p.Name is "" because it was excluded
		// more info about this at https://docs.royaleapi.com/#/field_filter
		for i:=0;i<len(popplayer);i++{
			pretty.Println(popplayer[i].Name)
		}
	}
}