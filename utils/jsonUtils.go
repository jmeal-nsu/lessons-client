package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (t *TimeTable) UpdateTimetable() error {

	resp, err := getJson(getTT)
	if err != nil {
		return err
	}

	var result []Lesson
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &result)
	for _, rec := range result {
		(*t).AddLesson(rec.StartTime, rec.WeekDay, rec.Type, rec.Teacher, rec.Classroom,
			rec.Name, rec.Pavilion)
	}
	defer resp.Body.Close()
	return nil
}

func GetPlaces() Places {
	resp, err := getJson(places)
	if err != nil {
		return nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	var jsonData []Place
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}

	defer resp.Body.Close()

	return jsonData
}

func GetTeachers() Teachers {
	resp, err := getJson(teachers)
	if err != nil {
		return nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	var jsonData []Teacher
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}

	defer resp.Body.Close()

	return jsonData
}
