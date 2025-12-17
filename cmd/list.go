/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the things on gitignore",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		
	},
}





