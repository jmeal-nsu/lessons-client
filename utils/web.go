package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const url = "http://jmeal-nsu.ru/"
const getTT = "lessons/timetable"
const healthcheck = "healthcheck"
const places = "places"
const teachers = "teachers"

const delaySec = 120

func getJson(arg string) (*http.Response, error) {
	resp, err := http.Get(url + arg)
	if err != nil {
		fmt.Println("Failed to connect to the server!")
		return nil, err
	}
	return resp, nil
}

func Healthcheck() bool {
	resp, err := http.Get(url + healthcheck)
	if err != nil || resp.StatusCode != 200 {
		return false
	}
	defer resp.Body.Close()
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

func Put(data []map[string]interface{}, where string) error {
	payload, err := json.Marshal(data)
	req, err := http.NewRequest("PUT", where, bytes.NewBuffer(payload))
	if err != nil {
		return errors.New("error while sending put request")
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)

	return nil
}

func PutPlace(line []string) error {

	data := []map[string]interface{}{
		{
			"cabinet":  argDecide(line[2]),
			"pavilion": argDecide(line[3]),
		},
	}
	return Put(data, url+"places")
}

func argDecide(a string) interface{} {
	if a == "_" {
		return nil
	}
	return a
}

func PutTeacher(line []string) error {
	data := []map[string]interface{}{
		{
			"name":       argDecide(line[2]),
			"surname":    argDecide(line[3]),
			"patronymic": argDecide(line[4]),
		},
	}
	return Put(data, url+"teachers")
}

func PutLesson(line []string) error {

	tid, err := strconv.Atoi(line[6])
	if err != nil {
		return err
	}
	pid, err := strconv.Atoi(line[7])
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"subject":    line[2],
		"type":       line[3],
		"week_day":   line[4],
		"start":      line[5],
		"teacher_id": tid,
		"place_id":   pid,
	}

	payload, err := json.Marshal(data)
	req, err := http.NewRequest("PUT", url+"lessons/", bytes.NewBuffer(payload))
	if err != nil {
		return errors.New("error while sending put request")
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)

	return nil
}
