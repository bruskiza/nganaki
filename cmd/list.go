/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (

	"github.com/bruskiza/nganaki/internal/utils"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the things on gitignore",
	Run: func(cmd *cobra.Command, args []string) {
		d := utils.NewDownloader()
		languages, err := d.ListLanguages()
		if err != nil {
			cmd.Printf("⛔️ Error listing languages: %v\n", err)
			return
		}

		cmd.Printf("📝 Available gitignore languages:\n")
		for _, lang := range languages {
			cmd.Printf("- %s\n", lang)
		}
	
	
		
	},
}





