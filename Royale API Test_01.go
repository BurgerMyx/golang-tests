package main

import (
	"fmt"

	"github.com/Altarrel/goroyale"
	"github.com/kr/pretty"
)

var token = "Insert API Key Here"

func main() {
	c, err := goroyale.New(token, 0)
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
		//Prints out available tournaments
		for i:=0;i<len(tourney);i++{
			if(tourney[i].CurrentPlayers >= tourney[i].MaxPlayers*3/4){ //Will only print tourneys of players 0.75 of max player
				continue
			}else{
				pretty.Printf("%s | %s\n", tourney[i].Name, tourney[i].Tag)
				pretty.Printf("Curr: %d | Max: %d\n", tourney[i].CurrentPlayers, tourney[i].MaxPlayers)
			}
		}
	}

	params2:=map[string][]string{}

	popplayer, err := c.PopularPlayers(params2)
	if err != nil {
		fmt.Println(err)
	} else {
		for i:=0;i<len(popplayer);i++{
			pretty.Println(popplayer[i].Name) //Prints out popular players
		}
	}
}
