package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("неверный формат данных: %s", data)
	}

	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка преобразования количества шагов: %w", err)
	}

	if steps <= 0 {
		return 0, 0, fmt.Errorf("количество шагов должно быть больше нуля")
	}

	duration, err := time.ParseDuration(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка преобразования продолжительности: %w", err)
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	// Дистанция в метрах
	distanceMeters := float64(steps) * stepLength
	// Дистанция в километрах
	distanceKm := distanceMeters / mInKm

	// Калории
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return fmt.Sprintf(`Количество шагов: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.`, steps, distanceKm, calories)
}
