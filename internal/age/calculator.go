package age

import "time"

func CalculateAge(dob time.Time, now time.Time) int {
	age := now.Year() - dob.Year()

	birthdayThisYear := time.Date(
		now.Year(),
		dob.Month(),
		dob.Day(),
		0,
		0,
		0,
		0,
		now.Location(),
	)

	if now.Before(birthdayThisYear) {
		age--
	}

	return age
}

func DaysAlive(dob time.Time, now time.Time) int {
	start := time.Date(
		dob.Year(),
		dob.Month(),
		dob.Day(),
		0,
		0,
		0,
		0,
		time.UTC,
	)

	end := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0,
		0,
		0,
		0,
		time.UTC,
	)

	return int(end.Sub(start).Hours() / 24)
}
