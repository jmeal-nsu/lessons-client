package utils

func Start() {

	println("Trying to connect...")

	if !Healthcheck() {
		println("Failed to connect!")
		return
	}

	println("Welcome!")

	var timetable = new(TimeTable)
	timetable.UpdateTimetable()
	timetable.Print()

}
