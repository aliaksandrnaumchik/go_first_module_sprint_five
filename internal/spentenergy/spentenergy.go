package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть больше 0")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть больше 0")
	}
	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть больше 0")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть больше 0")
	}

	speed := MeanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	baseCalories := (weight * speed * durationInMinutes) / minInH

	calories := baseCalories * walkingCaloriesCoefficient

	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть больше 0")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть больше 0")
	}
	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть больше 0")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть больше 0")
	}

	speed := MeanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	calories := (weight * speed * durationInMinutes) / minInH

	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	dist := Distance(steps, height)

	durationHours := duration.Hours()

	speed := dist / durationHours

	return speed
}

func Distance(steps int, height float64) float64 {
	stepsFloat := float64(steps)

	stepLength := height * stepLengthCoefficient

	totalDistanceMeters := stepsFloat * stepLength

	distanceKm := totalDistanceMeters / mInKm

	return distanceKm
}
