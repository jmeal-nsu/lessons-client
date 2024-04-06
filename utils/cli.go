package utils

import (
	"bufio"
	"errors"
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
		if parseCmd(&cmd, timetable) != nil {
			return
		}
		cmd = []string{}
	}

}

func parseCmd(cmd *[]string, tt *TimeTable) error {

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
	case "healthcheck":
		{
			return check()
		}
	case "add":
		{
			add(*cmd)
		}
	default:
		fmt.Println("Unknown command! Type 'help' to get list of available commands.")
	}
	return nil
}

func help() {
	fmt.Println("Help:\n" +
		"	help - Print commands and their description.\n" +
		"	healthcheck - Check the connection to the server. If the connection is interrupted, the program will warn about it and terminate.\n" +
		"	quit - quit program\n" +
		"	show <args...> - display information\n" +
		"		timetable - print whole timetable\n" +
		"		lesson <1...7(day)> <hh:mm(start time)> - print information about exact lesson\n" +
		"		day <1...7(day)> - print timetable for the exact day\n" +
		"		teachers - print list of all teachers\n" +
		"		places - print list of all places\n" +
		"	add <args...> - add new item\n" +
		"		teacher <name> <surname> <patronymic|_>\n" +
		"		place <cabinet_number|_> <_|new|main|laboratory|sport_new|sport_old|stadium>\n" +
		"		lesson <name> <Lecture|Practice|Laboratory> <Monday|...|Sunday> <hh:mm(start time)> <teacher_id> <place_id>\n" +
		"\nINFO: '_' = null")
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
			les, err := tt.FindLesson(line[2], line[3])
			if err != nil {
				return
			}
			PrintLesson(les)
		}
	case "day":
		{
			for len(line) < 3 {
				read(&line, argPointer)
			}
			//show day <day>
			tt.PrintDay(line[2])
		}
	case "places":
		{
			//show places
			p := GetPlaces()
			p.PrintPlaces()
		}
	case "teachers":
		{
			//show teachers
			t := GetTeachers()
			t.PrintTeachers()
		}

	default:
		fmt.Println("Wrong args: " + line[1])
	}
}

func check() error {
	res := Healthcheck()
	if !res {
		fmt.Println("Connection failed!")
		return errors.New("")
	}
	fmt.Println("Connected!")
	return nil
}

func read(args *[]string, pointer string) {
	fmt.Print(pointer)

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	lines := strings.Fields(scanner.Text())

	*args = append(*args, lines...)
}

func add(line []string) {
	for len(line) <= 1 {
		read(&line, argPointer)
	}

	switch line[1] {
	case "lesson":
		{
			//add lesson <subject> <type> <week_day> <start> <teacher_id> <place_id>
			for len(line) < 8 {
				read(&line, argPointer)
			}
			err := PutLesson(line)
			if err != nil {

				fmt.Println(err)
				return
			}
		}
	case "place":
		{
			//add place <cabinet> <pavilion>
			for len(line) < 4 {
				read(&line, argPointer)
			}
			err := PutPlace(line)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	case "teacher":
		{
			//add teacher <name> <surname> <pat>
			for len(line) < 5 {
				read(&line, argPointer)
			}
			err := PutTeacher(line)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

}
