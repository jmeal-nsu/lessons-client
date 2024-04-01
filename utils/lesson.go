package utils

import (
	"errors"
	//"fmt"
	"fmt"
	"strconv"
)

type Lesson struct {
	Name      string `json:"name"`
	WeekDay   int    `json:"weekDay"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Teacher   string `json:"teacher"`
	Classroom string `json:"classroom"`
}

type TimeTable [7][]Lesson

const header = "|       |     mon     |     tue     |     wed     |     thu     |     fri     |     sat     |     sun     |"
const horizSeparator = "+-------+-------------+-------------+-------------+-------------+-------------+-------------+-------------+"

var times = []string{"09:00", "10:50", "12:40", "14:30", "16:20", "18:20", "20:00"}

func (t *TimeTable) AddLesson(name string, weekDay int,
	start, end string,
	teacher string, classroom string) {

	var lessons = t[weekDay]

	(*t)[weekDay] = append(lessons,
		Lesson{name, weekDay, start, end,
			teacher, classroom})
}

func (t *TimeTable) PrintLesson(weekDayS string, time string) {

	weekDay, err := parseDay(weekDayS)
	if err != nil {
		return
	}

	lessons := (*t)[weekDay]

	for _, les := range lessons {
		if les.StartTime == time {
			fmt.Println(les)
		}
	}
}

func (t *TimeTable) PrintDay(weekDayS string) {

	for _, time := range times {
		t.PrintLesson(weekDayS, time)
	}
}

func (t *TimeTable) Print() {

	var lessons [7]string

	fmt.Println(horizSeparator)
	fmt.Println(header)
	fmt.Println(horizSeparator)

	for _, time := range times {
		for i := 0; i < 7; i++ {
			for _, les := range t[i] {
				if time == les.StartTime {
					lessons[i] = les.Name
				} else {
					lessons[i] = ""
				}
			}
		}

		fmt.Print(fmt.Sprintf("| %-6s| %-11.11s | %-11.11s | %-11.11s | %-11.11s | %-11.11s | %-11.11s | %-11.11s |\n", time,
			lessons[0], lessons[1], lessons[2], lessons[3], lessons[4], lessons[5], lessons[6]))
		fmt.Println(horizSeparator)
	}
}

func parseDay(day string) (int, error) {
	weekDay, err := strconv.Atoi(day)
	if weekDay <= 0 || weekDay > 7 || err != nil {
		return -1, errors.New("")
	}
	return weekDay - 1, nil
}
