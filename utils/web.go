package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const urlTest = "https://58c7d4ca-02d4-42e6-a071-82c5e296efdc.mock.pstmn.io/"
const url = "http://jmeal-nsu.ru/"
const getTT = "lessons/timetable"
const healthcheck = "healthcheck"

const delaySec = 120

func (t *TimeTable) UpdateTimetable() error {
	resp, err := http.Get(url + getTT)
	if err != nil {
		fmt.Println("Failed to connect to the server!")
		return err
	}
	defer resp.Body.Close()

	var result []Lesson
	body, err := ioutil.ReadAll(resp.Body) //todo deprecated
	err = json.Unmarshal(body, &result)
	for _, rec := range result {
		(*t).AddLesson(rec.StartTime, rec.WeekDay, rec.Type, rec.Teacher, rec.Classroom,
			rec.Name, rec.Pavilion)
	}
	return nil
}

func Healthcheck() bool {
	resp, err := http.Get(url + healthcheck)
	if err != nil || resp.StatusCode != 200 {
		return false
	}
	resp.Body.Close()
	return true
}

func (t *TimeTable) UpdateTimer() {
	waitTime := delaySec * time.Second

	for {
		time.Sleep(waitTime)
		err := t.UpdateTimetable()
		if err != nil {
			println("Error while refreshing the timetable!")
			return
		}
	}
}
