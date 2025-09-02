package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	//lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	stringSlice := strings.Split(data, ",")
	if len(stringSlice) != 3 {
		return 0, "", 0, fmt.Errorf("некорректный формат данных: %s", data)
	}

	steps, err := strconv.Atoi(stringSlice[0])

	if err != nil {
		return 0, "", 0, err
	}

	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("Steps Errors")
	}

	if len(stringSlice[1]) < 2 {
		return 0, "", 0, fmt.Errorf("Action type Errors")
	}

	action := stringSlice[1]

	duration, err := time.ParseDuration(stringSlice[2])

	if err != nil {
		return 0, action, 0, err
	}

	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("Steps Errors")
	}

	return steps, action, duration, nil
}

func distance(steps int, height float64) float64 {

	lenght := height * stepLengthCoefficient

	result := float64(steps) * lenght / mInKm

	return result
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {

	if duration <= 0 {
		return 0
	}
	speed := distance(steps, height) / duration.Hours()

	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, action, duration, error := parseTraining(data)

	if error != nil {
		fmt.Println(error)
		return "", error
	}

	dist := distance(steps, height)

	speed := meanSpeed(steps, height, duration)

	str := "Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n"

	switch action {
	case "Бег":
		caloriesRun, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
		return fmt.Sprintf(str, action, duration.Hours(), dist, speed, caloriesRun), nil
	case "Ходьба":
		caloriesWalk, err := WalkingSpentCalories(steps, weight, height, duration)

		if err != nil {
			log.Println(err)
			return "", err
		}
		return fmt.Sprintf(str, action, duration.Hours(), dist, speed, caloriesWalk), nil

	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("incorrect data: %d, %.2f, %.2f, %s", steps, weight, height, duration)
	}

	speed := meanSpeed(steps, height, duration)

	minutes := duration.Minutes()

	result := weight * speed * float64(minutes) / float64(minInH)

	return result, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Некоректные данные")
	}

	speed := meanSpeed(steps, height, duration)

	result := walkingCaloriesCoefficient * weight * speed * (duration.Minutes() / float64(minInH))

	return result, nil
}
