package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"

	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")

	if len(parts) != 3 {
		return fmt.Errorf("неверный формат данных: ожидается 3 части, получено %d", len(parts))
	}

	stepsStr := strings.TrimSpace(parts[0])
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return fmt.Errorf("ошибка при парсинге количества шагов: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным")
	}
	t.Steps = steps

	trainingType := strings.TrimSpace(parts[1])
	t.TrainingType = trainingType

	durationStr := strings.TrimSpace(parts[2])
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("ошибка при парсинге длительности: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("продолжительность должна быть положительной")
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	var (
		trainingDistance float64
		trainingSpeed    float64
		trainingCalories float64
		errCalories      error
	)

	trainingDistance = spentenergy.Distance(t.Steps, t.Height)
	trainingSpeed = spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	switch strings.ToLower(t.TrainingType) {
	case "бег":
		trainingCalories, errCalories = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if errCalories != nil {
			return "", errCalories
		}
	case "ходьба":

		trainingCalories, errCalories = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if errCalories != nil {
			return "", errCalories
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}

	result := fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		trainingDistance,
		trainingSpeed,
		trainingCalories,
	)

	return result, nil
}
