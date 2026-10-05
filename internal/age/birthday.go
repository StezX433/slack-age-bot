package age

import "time"

func NextBirthday(dob time.Time, now time.Time) time.Time {
	birthday := time.Date(
		now.Year(),
		dob.Month(),
		dob.Day(),
		0,
		0,
		0,
		0,
		now.Location(),
	)

	today := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0,
		0,
		0,
		0,
		now.Location(),
	)

	if birthday.Before(today) {
		birthday = time.Date(
			now.Year()+1,
			dob.Month(),
			dob.Day(),
			0,
			0,
			0,
			0,
			now.Location(),
		)
	}

	return birthday
}
