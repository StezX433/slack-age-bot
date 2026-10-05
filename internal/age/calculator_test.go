package age

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	dob := time.Date(2004, 7, 15, 0, 0, 0, 0, time.UTC)
	now := time.Date(2026, 10, 6, 0, 0, 0, 0, time.UTC)

	got := CalculateAge(dob, now)
	want := 22

	if got != want {
		t.Errorf("CalculateAge() = %d, want %d", got, want)
	}
}

func TestCalculateAgeBeforeBirthday(t *testing.T) {
	dob := time.Date(2004, 7, 15, 0, 0, 0, 0, time.UTC)
	now := time.Date(2026, 7, 14, 0, 0, 0, 0, time.UTC)

	got := CalculateAge(dob, now)
	want := 21

	if got != want {
		t.Errorf("CalculateAge() = %d, want %d", got, want)
	}
}

func TestDaysAlive(t *testing.T) {
	dob := time.Date(2004, 7, 15, 0, 0, 0, 0, time.UTC)
	now := time.Date(2004, 7, 20, 0, 0, 0, 0, time.UTC)

	got := DaysAlive(dob, now)
	want := 5

	if got != want {
		t.Errorf("DaysAlive() = %d, want %d", got, want)
	}
}

func TestNextBirthday(t *testing.T) {
	dob := time.Date(2004, 7, 15, 0, 0, 0, 0, time.UTC)
	now := time.Date(2026, 10, 6, 0, 0, 0, 0, time.UTC)

	got := NextBirthday(dob, now)

	want := time.Date(2027, 7, 15, 0, 0, 0, 0, time.UTC)

	if !got.Equal(want) {
		t.Errorf("NextBirthday() = %v, want %v", got, want)
	}
}

func TestNextMilestone(t *testing.T) {
	tests := []struct {
		age  int
		want int
	}{
		{18, 25},
		{22, 25},
		{25, 30},
		{31, 35},
		{42, 45},
	}

	for _, test := range tests {
		got := NextMilestone(test.age)

		if got != test.want {
			t.Errorf(
				"NextMilestone(%d) = %d, want %d",
				test.age,
				got,
				test.want,
			)
		}
	}
}
