package main

import "fmt"

func weekDay(day int) string {
	switch day { // there's no need for break command
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid day"
	}
}

func weekDay2(day int) string {
	switch {
	case day == 1:
		return "Sunday"
	case day == 2:
		return "Monday"
	case day == 3:
		return "Tuesday"
	case day == 4:
		return "Wednesday"
	case day == 5:
		return "Thursday"
	case day == 6:
		return "Friday"
	case day == 7:
		return "Saturday"
	default:
		return "Invalid day"
	}
}

func weekDay3(day int) string {
	var dayOfTheWeek string

	switch {
	case day == 1:
		dayOfTheWeek = "Sunday"
	case day == 2:
		dayOfTheWeek = "Monday"
	case day == 3:
		dayOfTheWeek = "Tuesday"
	case day == 4:
		dayOfTheWeek = "Wednesday"
	case day == 5:
		dayOfTheWeek = "Thursday"
	case day == 6:
		dayOfTheWeek = "Friday"
	case day == 7:
		dayOfTheWeek = "Saturday"
	default:
		dayOfTheWeek = "Invalid day"
	}

	return dayOfTheWeek
}

func weekDay4(day int) string {
	var dayOfTheWeek string

	switch {
	case day == 1:
		dayOfTheWeek = "Sunday"
		fallthrough // goes directly no next case
	case day == 2:
		dayOfTheWeek = "Monday"
	case day == 3:
		dayOfTheWeek = "Tuesday"
	case day == 4:
		dayOfTheWeek = "Wednesday"
	case day == 5:
		dayOfTheWeek = "Thursday"
	case day == 6:
		dayOfTheWeek = "Friday"
	case day == 7:
		dayOfTheWeek = "Saturday"
	default:
		dayOfTheWeek = "Invalid day"
	}

	return dayOfTheWeek
}

func main() {
	fmt.Println("Switch")

	day := weekDay(2)
	fmt.Println(day)

	day = weekDay(200)
	fmt.Println(day)

	day = weekDay2(3)
	fmt.Println(day)

	day = weekDay2(200)
	fmt.Println(day)

	day = weekDay3(4)
	fmt.Println(day)

	day = weekDay3(200)
	fmt.Println(day)

	day = weekDay4(1)
	fmt.Println(day)
}
