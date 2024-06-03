package report

import (
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/jpodlasnisky/stresstest/constants"
	"github.com/jpodlasnisky/stresstest/models"
)

func GenerateReport(totalResult models.TotalResult) {
	stats := calculateStats(totalResult)
	printReport(totalResult.URL, totalResult.TotalDuration, stats)
}

func calculateStats(totalResult models.TotalResult) models.ReportStats {
	stats := models.ReportStats{
		TotalRequests:      len(totalResult.Results),
		StatusDistribution: make(map[int]int),
		MinDuration:        math.MaxFloat64,
	}

	for _, result := range totalResult.Results {
		if result.Error {
			stats.ErrorCount++
			log.Printf("Error count: %d - Falha ao executar requisição: %s \n ", stats.ErrorCount, result.ErrorMessage)

		} else {
			if result.StatusCode == http.StatusOK {
				stats.SuccessCount++
			}
		}

		if result.Duration < stats.MinDuration {
			stats.MinDuration = result.Duration
		}

		if result.Duration > stats.MaxDuration {
			stats.MaxDuration = result.Duration
		}

		stats.StatusDistribution[result.StatusCode]++
	}

	return stats
}

func printReport(url string, totalDuration float64, stats models.ReportStats) {
	fmt.Println(constants.ReportHeader)
	fmt.Printf("URL Alvo: %s\n", url)
	fmt.Printf("Requisições: %d\n", stats.TotalRequests)
	fmt.Printf("Bem-Sucedidas: %d\n", stats.SuccessCount)
	fmt.Printf("Erros: %d\n", stats.ErrorCount)
	fmt.Printf("Duração Total: %.2f ms\n", totalDuration)
	fmt.Printf("Média por Requisição: %.2f ms\n", totalDuration/float64(stats.TotalRequests))
	fmt.Printf("Duração Mínima da Requisição: %.2f ms\n", stats.MinDuration)
	fmt.Printf("Duração Máxima da Requisição: %.2f ms\n", stats.MaxDuration)
	fmt.Println("\nDistribuição dos Códigos de Status HTTP:")
	for status, count := range stats.StatusDistribution {
		fmt.Printf("  %d: %d\n", status, count)
	}
	fmt.Println(constants.ReportFooter)
}
