/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var relativePath, pemission, directoryOnly bool
var pipeMap map[int]bool
var directoriesCount, filesCount, nestedLevel int

func printEntry(entry os.DirEntry, directory string) {
	if relativePath {
		fmt.Print(directory, "/", entry.Name(), "\n")
	} else if directoryOnly {
		if entry.IsDir() {
			fmt.Print(entry.Name(), "\n")
		}
	}else if pemission{
		info, _ := entry.Info()
		pemissions := info.Mode().Perm()
		fmt.Print("[",pemissions, "]  ", entry.Name(),"\n")
	} else {
		fmt.Print(entry.Name(), "\n")
	}
}

func isEntryrRelevent(entry os.DirEntry) bool {
	if directoryOnly {
		return entry.IsDir()
	} else {
		return true
	}
}

func printStru(stru string, entry os.DirEntry) {
	if isEntryrRelevent(entry) {
		fmt.Print(stru)
	}
}

func printTreeStructure(directory string, nestLevel int) {
	if nestedLevel!=0 &&nestLevel >= nestedLevel {
		return
	}
	entries, err := os.ReadDir(directory)
	if err != nil {
		return
	}

	balence := 0

	for i := 0; i < nestLevel; i++ {
		balence += 1
	}

	for i, entry := range entries {
		var j int
		for j = 0; j < balence; j++ {
			if pipeMap[j] {
				printStru("│   ", entry)
			} else {
				printStru("    ", entry)
			}
		}
		if i == len(entries)-1 {
			printStru("└── ", entry)
			pipeMap[j] = false
		} else {
			printStru("├── ", entry)
			pipeMap[j] = true
		}
		printEntry(entry, directory)

		if entry.IsDir() {
			directoriesCount += 1
			printTreeStructure(directory+"/"+entry.Name(), nestLevel+1)
		} else {
			filesCount += 1
		}
	}
}

func printSummery() {

	var a, b string

	if directoriesCount == 1 {
		a = "directory"
	} else {
		a = "directories"
	}

	if filesCount == 1 {
		b = "file"
	} else {
		b = "files"
	}

	if directoryOnly {
		fmt.Println(directoriesCount, a)
	} else {
		fmt.Print(directoriesCount, " ", a)
		fmt.Println(",", filesCount, b)

	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tree",
	Short: "A simple command line program that implements Unix tree like functionality.",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// relativePath, directoryOnly = false, false
		pipeMap = make(map[int]bool)
		directory := args[0]
		var nestLevel int
		fmt.Println(directory)
		printTreeStructure(directory, nestLevel)
		fmt.Print("\n")
		printSummery()
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
	rootCmd.PersistentFlags().BoolVarP(&relativePath, "relative-path", "f", false, "Print the relative path to the directory being searched.")
	rootCmd.PersistentFlags().BoolVarP(&directoryOnly, "directory-only", "d", false, "Only print directories, not files.")
	rootCmd.PersistentFlags().BoolVarP(&pemission, "permission", "p", false, "Print file permissions for all files.")
	rootCmd.PersistentFlags().IntVarP(&nestedLevel, "level", "L", 0, "Allow traversing specified nested levels only.")
}
