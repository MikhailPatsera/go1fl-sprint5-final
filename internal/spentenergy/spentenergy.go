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

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distanceMeters := float64(steps) * stepLength
	return distanceMeters / mInKm
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	if steps < 0 {
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validateCaloriesParams(steps, weight, height, duration); err != nil {
		return 0, err
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationInMinutes) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validateCaloriesParams(steps, weight, height, duration); err != nil {
		return 0, err
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationInMinutes) / minInH
	calories *= walkingCaloriesCoefficient

	return calories, nil
}

func validateCaloriesParams(steps int, weight, height float64, duration time.Duration) error {
	switch {
	case steps <= 0:
		return fmt.Errorf("количество шагов должно быть положительным")
	case weight <= 0:
		return fmt.Errorf("вес должен быть положительным")
	case height <= 0:
		return fmt.Errorf("рост должен быть положительным")
	case duration <= 0:
		return fmt.Errorf("длительность должна быть положительной")
	default:
		return nil
	}
}