package utils

import "fmt"

type Teacher struct {
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic"`
}

type Teachers []Teacher

func (t Teachers) PrintTeachers() {
	for id, teacher := range t {
		teacher.PrintTeacher(id)
	}
}

func (t Teacher) PrintTeacher(id int) {
	if t.Patronymic == "" {
		fmt.Println(id, t.Name, t.Surname)
	} else {
		fmt.Println(id, t.Name, t.Patronymic, t.Surname)
	}
}
