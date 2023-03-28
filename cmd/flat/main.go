package main

import (
	"encoding/json"
	"fmt"
)

var greyColor = `grey`
var blueColor = `blue`
var violetColor = `violet`
var greenColor = `green`

type Room struct {
	Name  string
	Area  float64
	color string
	Room  *Room
}

type MyFlat struct {
	Bathroom Room `json:"bathroom"`
	Lounge   Room `json:"lounge"`
	Bedroom  Room `json:"bedroom"`
}

func main() {
	greyColor = `dark grey`

	bathroom := Room{
		Name:  `Chamber of Secrets`,
		Area:  3.5,
		color: blueColor,
	}
	office := &Room{
		color: greyColor,
		Area:  4.2,
		Name:  `My Kingdom`,
	}
	bedroom := Room{
		Name:  `Sleeping`,
		color: violetColor,
		Area:  12,
	}
	lounge := Room{`Living room`, 25, greenColor, office}

	myAppartment := MyFlat{bathroom, lounge, bedroom}

	res, _ := json.Marshal(myAppartment)
	fmt.Printf("\n%v \n\n %s\n", myAppartment, string(res))

	// exercise
	// make two couples, which have and don't have a child
}
