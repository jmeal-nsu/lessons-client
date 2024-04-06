package utils

import (
	"errors"
	"fmt"
	"strconv"
)

type Lesson struct {
	Name      string `json:"subject"`
	WeekDay   string `json:"week_day"`
	StartTime string `json:"start"`
	Teacher   string `json:"teacher"`
	Classroom string `json:"cabinet"`
	Type      string `json:"type"`
	Pavilion  string `json:"pavilion"`
}

type TimeTable [7][]Lesson

const header = "|          |     mon     |     tue     |     wed     |     thu     |     fri     |     sat     |     sun     |"
const horizSeparator = "+----------+-------------+-------------+-------------+-------------+-------------+-------------+-------------+"

var times = []string{"09:00", "10:50", "12:40", "14:30", "16:20", "18:10", "20:00"}
var dayToNumber = map[string]int{
	"Monday":    0,
	"Tuesday":   1,
	"Wednesday": 2,
	"Thursday":  3,
	"Friday":    4,
	"Saturday":  5,
	"Sunday":    6,
}

func (t *TimeTable) AddLesson(start, weekDay, typ, teacher, cabinet, subject, pavilion string) {
	day := dayToNumber[weekDay]

	(*t)[day] = append(t[day], Lesson{
		Name:      subject,
		WeekDay:   weekDay,
		StartTime: start,
		Teacher:   teacher,
		Classroom: cabinet,
		Type:      typ,
		Pavilion:  pavilion,
	})
}

func (t *TimeTable) FindLesson(day, start string) (Lesson, error) {
	num, err := getInt(day)
	if err != nil {
		fmt.Println("Wrong day format!")
		return Lesson{}, err
	}

	for _, lesson := range t[num-1] {
		if lesson.StartTime == start {
			return lesson, nil
		}
	}

	return Lesson{}, errors.New("")
}

func PrintLesson(lesson Lesson) {
	fmt.Println(lesson.Name)
	fmt.Println(lesson.Type)
	fmt.Println(lesson.StartTime)
	fmt.Println(lesson.Teacher)
	fmt.Println(lesson.Classroom, lesson.Pavilion)
}

func (t *TimeTable) PrintDay(day string) {
	num, err := getInt(day)
	if err != nil || num <= 0 || num >= 8 {
		fmt.Println("Wrong day format!")
		return
	}

	for _, lesson := range t[num-1] {
		PrintLesson(lesson)
		fmt.Println("------o------")
	}
}

func (t *TimeTable) Print() {

	var lessons [7]string

	fmt.Println(horizSeparator)
	fmt.Println(header)
	fmt.Println(horizSeparator)

	for _, s := range times {
		for i := 0; i < 7; i++ {
			for _, les := range t[i] {
				if les.StartTime == s {
					lessons[i] = les.Name
				}
			}
		}

		fmt.Print(fmt.Sprintf("| %-9s| %-11.11s | %-11.11s | %-11.11s | %-11.11s | %-11.11s | %-11.11s | %-11.11s |\n",
			s, lessons[0], lessons[1], lessons[2], lessons[3], lessons[4], lessons[5], lessons[6]))
		fmt.Println(horizSeparator)
		lessons = [7]string{}
	}
}

func getInt(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return -1, errors.New("")
	}
	return num, nil
}
