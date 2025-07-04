package main

import (
	"dusty/internal/adapter/reporter"
	"dusty/internal/adapter/system/macos"
	"fmt"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var rootCmd = &cobra.Command{
	Use:   "dusty",
	Short: "Dusty - Safe macOS cleaner & reporter",
	Run: func(cmd *cobra.Command, args []string) {
		log := logrus.New()
		green := color.New(color.FgGreen).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()
		red := color.New(color.FgRed).SprintFunc()

		fmt.Printf("%s Starting Safe Clean...\n", green("[INFO]"))

		startTime := time.Now()
		cleaner := &macos.Cleaner{Logger: log}
		report, err := cleaner.SafeClean(true)
		if err != nil {
			fmt.Printf("%s SafeClean failed: %v\n", red("[ERROR]"), err)
			return
		}
		fmt.Printf("%s Safe Clean finished. Generating JSON report...\n", green("[INFO]"))
		endTime := time.Now()

		jsonReporter := &reporter.JsonReport{
			Items:     report,
			StartTime: startTime,
			EndTime:   endTime,
		}
		_, rbytes := jsonReporter.GenerateReport()

		file := "clean_report.json"
		if err := os.WriteFile(file, rbytes, 0644); err != nil {
			fmt.Printf("%s Could not write report: %v\n", red("[ERROR]"), err)
			return
		}
		fmt.Printf("%s JSON report saved to %s\n", yellow("[SUCCESS]"), file)
	},
}

func Execute() error {
	return rootCmd.Execute()
}
