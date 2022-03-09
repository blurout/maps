package main

import (
	"maps/maps"
)

func main(){
	table := maps.MakeMap()
	table.Insert("Fatih")
	table.Insert("Monkey")
	table.Search("Monk")
}

