/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math"
	"strconv"

	"github.com/spf13/cobra"
)

// findAreaCircleCmd represents the findAreaCircle command
var findAreaCircleCmd = &cobra.Command{
	Use:   "findAreaCircle",
	Short: "calculate circular area from flag",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("findAreaCircle called")

		if len(args) == 0 {
			fmt.Printf("Pleas to Add your parameter")
			return
		}

		res := areaCircle(args) * 3.14
		fmt.Printf("Result %v \n", res)
	},
}

func init() {
	rootCmd.AddCommand(findAreaCircleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findAreaCircleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findAreaCircleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func areaCircle(arg []string) float64 {
	a, _ := strconv.ParseFloat(arg[0], 64)
	b := math.Pow(a, 2)

	return b
}
