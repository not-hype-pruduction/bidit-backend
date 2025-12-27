package utils

import (
	"strings"
)

// PBNToSlice переводит руку из формата PBN (S.H.D.C) в слайс []string
// Пример входа: "AKJ.T98.632.AQ4" или "N:AKJ.T98.632.AQ4"
// Пример выхода: ["SA", "SK", "SJ", "H10", "H9", "H8", "D6", "D3", "D2", "CA", "CQ", "C4"]
func PBNToSlice(pbnHand string) []string {
	if strings.Contains(pbnHand, ":") {
		parts := strings.Split(pbnHand, ":")
		if len(parts) > 1 {
			pbnHand = parts[1]
		}
	}

	suits := strings.Split(pbnHand, ".")
	if len(suits) != 4 {
		return nil // Ошибка формата
	}

	suitPrefixes := []string{"S", "H", "D", "C"}
	var result []string

	for i, ranks := range suits {
		suitPrefix := suitPrefixes[i]

		for _, char := range ranks {
			rank := string(char)

			if rank == "T" {
				rank = "10"
			}

			result = append(result, suitPrefix+rank)
		}
	}

	return result
}
