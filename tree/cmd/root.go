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
var pipeMap map[int]bool
var directoriesCount, filesCount int

func printTreeStructure(directory string, nestLevel int, isBase bool) {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return
	}

	balence := 0

	for i := 0; i < nestLevel; i++ {
		balence+= 1
	}

	for i, entry := range entries {
		var j int
		for j = 0; j < balence; j++ {
			if pipeMap[j] {
				fmt.Print("│   ")
			}else {
				fmt.Print("    ")
			}
		}
		if i == len(entries) - 1 {
			fmt.Print("└── ",entry.Name(), "\n")
			pipeMap[j] = false
		}else {
			fmt.Print("├── ",entry.Name(), "\n")
			pipeMap[j] = true
		}

		if entry.IsDir() {
			directoriesCount += 1
			printTreeStructure(directory + "/" + entry.Name(), nestLevel + 1, false)
		}else {
			filesCount += 1
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
		pipeMap = make(map[int]bool)
		directory := args[0]
		var nestLevel int 
		fmt.Println(directory)
		printTreeStructure(directory, nestLevel, true)
		fmt.Print("\n")
		fmt.Println(directoriesCount, "directories,", filesCount, "files")
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


