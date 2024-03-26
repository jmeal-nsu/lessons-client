package main

import (
	"fmt"
	"jmeal_client/utils"
)

func main() {
	fmt.Println("Welcome to Timetable App!")
	var cache = utils.InitCache()
	cache.UpdateTimetable()
	cache.TimeTable.Print()
}
