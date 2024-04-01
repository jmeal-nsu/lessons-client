package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const cmdPointer = ">> "
const argPointer = "  -> "

func Start() {

	fmt.Println("Trying to connect...")

	if !Healthcheck() {
		fmt.Println("Failed to connect!")
		return
	}

	fmt.Println("Welcome!")

	var timetable = new(TimeTable)
	if timetable.UpdateTimetable() != nil {
		fmt.Println("Error while timetable refresh")
		return
	}
	timetable.Print()

	go timetable.UpdateTimer()

	cmdListen(timetable)

}

func cmdListen(timetable *TimeTable) {

	fmt.Println("Type your command:")
	var cmd []string

	for {
		read(&cmd, cmdPointer)

		if (len(cmd)) == 0 {
			continue
		}

		if cmd[0] == "quit" {
			fmt.Println("Bye!")
			return
		}
		parseCmd(&cmd, timetable)
		cmd = []string{}
	}

}

func parseCmd(cmd *[]string, tt *TimeTable) {

	switch (*cmd)[0] {
	case "help":
		{
			help()
			break
		}
	case "show":
		{
			show(*cmd, tt)
			break
		}
	default:
		fmt.Println("Unknown command! Type 'help' to get list of available commands.")

	}
}

func help() {
	fmt.Println("Help:\n" +
		"	help - Print commands and their description.\n" +
		"	healthcheck - Check the connection to the server. If the connection is interrupted, the program will warn about it and terminate.\n" +
		"	quit - quit program\n" +
		"	show <args...> - display information\n" +
		"		timetable - print whole timetable\n" +
		"		lesson <1...7(day)> <hh:mm(start time)> - print information about exact lesson\n" +
		"		day <1...7(day)> - print timetable for the exact day")
}

func show(line []string, tt *TimeTable) {

	for len(line) <= 1 {
		read(&line, argPointer)
	}

	switch line[1] {
	case "timetable":
		{
			tt.Print()
		}
	case "lesson":
		{
			for len(line) < 4 {
				read(&line, argPointer)
			}
			//show lesson <day> <start>
			tt.PrintLesson(line[2], line[3])
		}
	case "day":
		{
			for len(line) < 3 {
				read(&line, argPointer)
			}
			//show day <day>
			tt.PrintDay(line[2])
		}
	default:
		fmt.Println("Wrong args: " + line[1])
	}

}

func readUntilValidInput(line *[]string, argPointer string, validArgs int) {
	for len(*line) <= validArgs {
		read(line, argPointer)
	}
}

func read(args *[]string, pointer string) {
	fmt.Print(pointer)

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	lines := strings.Fields(scanner.Text())

	*args = append(*args, lines...)

}
