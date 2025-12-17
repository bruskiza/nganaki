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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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




