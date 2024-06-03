package cmd

import (
	"fmt"

	"github.com/jpodlasnisky/stresstest/loadtester"
	"github.com/jpodlasnisky/stresstest/report"
	"github.com/spf13/cobra"
)

var stressTestCmd = &cobra.Command{
	Use:   "stressTest",
	Short: "Teste de carga para uma URL específica",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetInt("requests")
		concurrency, _ := cmd.Flags().GetInt("concurrency")
		header, _ := cmd.Flags().GetStringSlice("header")

		result := loadtester.RunLoadTest(url, requests, concurrency, header)

		fmt.Printf("Teste de carga concluído para %s\n", url)

		report.GenerateReport(result)
	},
}

func init() {

	var url string
	var header []string
	var requests, concurrency int

	rootCmd.AddCommand(stressTestCmd)

	stressTestCmd.Flags().StringVarP(&url, "url", "u", "", "URL to be tested")
	stressTestCmd.Flags().IntVarP(&requests, "requests", "r", 100, "Total number of requests")
	stressTestCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 10, "Number of concurrent requests")
	stressTestCmd.Flags().StringSliceVarP(&header, "header", "H", []string{}, "Header to be included in the request")

	stressTestCmd.MarkFlagRequired("url")
}
