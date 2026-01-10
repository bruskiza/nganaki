/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (

	"github.com/bruskiza/nganaki/internal/utils"
	"github.com/spf13/cobra"
	
)

const gitIgnore = ".gitignore"

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets the file off github.com/github/gitignore. io for a specified language",
	Run: func(cmd *cobra.Command, args []string) {
		lang := cmd.Flags().Lookup("language").Value.String()
		
		cmd.Printf("🔎 Getting gitignore for language: %s\n", lang)

		d := utils.NewDownloader()
		d.Language = lang

		body, err := d.Download()

		if err != nil {
			cmd.Printf("⛔️ Error downloading gitignore: %v\n", err)
			return
		}

		if cmd.Flags().Lookup("save").Value.String() == "true" {
			if utils.IsGitRepository() == false {
				cmd.Printf("⛔️ Current directory is not a git repository. Not creating %s file.\n", gitIgnore)
				return
			}
			
			err = utils.WriteFileIfNotExists(gitIgnore, body)
			if err != nil {
				cmd.Printf("⛔️ Error writing gitignore to file: %v\n", err)
				return
			}
			
			cmd.Printf("✅ Saved gitignore to %s\n", gitIgnore)
		}

	},
}

func init () {
	getCmd.Flags().StringP("username", "u", "bruskiza", "GitHub username to list organizations for")
	getCmd.Flags().StringP("language", "l", "go", "Langauge to get. Will be title cased.")
	getCmd.Flags().Bool("save", false, "Save the gitignore to a file in the current directory")


}




