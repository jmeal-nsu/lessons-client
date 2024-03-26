package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://58c7d4ca-02d4-42e6-a071-82c5e296efdc.mock.pstmn.io/"

func (c *Cache) UpdateTimetable() {
	resp, err := http.Get(url + "tt")
	if err != nil {
		fmt.Println("Failed to connect to the server!")
		return
	}
	defer resp.Body.Close()

	var result []Lesson
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &result)
	for _, rec := range result {
		(*c).TimeTable.AddLesson(
			rec.Name, rec.WeekDay, rec.StartTime, rec.EndTime, rec.Teacher, rec.Classroom)
	}
}
