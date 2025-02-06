/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var relativePath bool

func printTreeStructure(directory string, nestLevel int, isBase bool) {

	entries, err := os.ReadDir(directory)
	if err != nil {
		return
	}
	for i, entry := range entries {
		for i := 1; i<nestLevel; i++ {
				fmt.Print(" │  ")
		}
		var arrow string
		if i == len(entries)-1{
			arrow = " └── "
		}else{
			arrow = " ├── "
		}
		if !entry.IsDir() {
			fmt.Print(arrow, entry.Name() + "\n")
		}else{
			if !isBase{
				fmt.Print(arrow + directory + "\n")
			}else {
				fmt.Print(directory + "\n")
			}
			printTreeStructure(directory + "/" + entry.Name(), nestLevel+1, false)
		}
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tree",
	Short: "A simple command line program that implements Unix tree like functionality.",
	Long: ``,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		directory := args[0]
		var nestLevel int 
		printTreeStructure(directory, nestLevel, true)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tree.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().BoolVarP(&relativePath, "Relative Path", "f", false, "Print relative path")
}


