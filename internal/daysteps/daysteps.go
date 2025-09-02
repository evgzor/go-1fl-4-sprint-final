package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/evgzor/go-1fl-4-sprint-final/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	stringSlice := strings.Split(data, ",")
	if len(stringSlice) != 2 {
		return 0, 0, fmt.Errorf("incorrect data format: %s", data)
	}

	steps, err := strconv.Atoi(stringSlice[0])

	if err != nil {
		return 0, 0, err
	}

	if steps <= 0 {
		return 0, 0, fmt.Errorf("steps errors")
	}

	duration, err := time.ParseDuration(stringSlice[1])

	if err != nil {
		return 0, 0, err
	}

	if duration.Seconds() <= 0 {
		return 0, 0, fmt.Errorf("steps errors")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {

	steps, duration, err := parsePackage(data)

	if err != nil {
		log.Println(err)
		return ""
	}

	if steps <= 0 || duration.Seconds() == 0 {
		return ""
	}

	distance := float64(steps) * stepLength / float64(mInKm)

	caloriesWalk, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	if err != nil {
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, caloriesWalk)
}
