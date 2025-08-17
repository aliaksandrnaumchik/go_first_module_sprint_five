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
	// 1. Проверка входных параметров на корректность
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

	// 2. Вычисление средней скорости
	speed := MeanSpeed(steps, height, duration)

	// 3. Перевод длительности в минуты
	durationInMinutes := duration.Minutes()

	// 4. Базовый расчет калорий
	baseCalories := (weight * speed * durationInMinutes) / minInH

	// 5. Применение корректирующего коэффициента
	calories := baseCalories * walkingCaloriesCoefficient

	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// 1. Проверка входных параметров на корректность
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

	// 2. Вычисление средней скорости
	speed := MeanSpeed(steps, height, duration)

	// 3. Перевод длительности в минуты
	durationInMinutes := duration.Minutes()

	// 4. Расчет потраченных калорий
	calories := (weight * speed * durationInMinutes) / minInH

	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// 1. Проверяем продолжительность
	if duration <= 0 {
		return 0
	}

	// 2. Вычисляем дистанцию
	dist := Distance(steps, height)

	// 3. Переводим длительность в часы
	durationHours := duration.Hours()

	// 4. Вычисляем среднюю скорость
	// Скорость = дистанция (км) / время (часы)
	speed := dist / durationHours

	return speed
}

func Distance(steps int, height float64) float64 {
	// 1. Приводим steps к float64
	stepsFloat := float64(steps)

	// 2. Вычисляем длину шага
	stepLength := height * stepLengthCoefficient

	// 3. Вычисляем общую дистанцию в метрах
	totalDistanceMeters := stepsFloat * stepLength

	// 4. Переводим метры в километры
	distanceKm := totalDistanceMeters / mInKm

	return distanceKm
}
