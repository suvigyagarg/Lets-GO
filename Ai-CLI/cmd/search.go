/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"
	"google.golang.org/genai"
	"github.com/spf13/cobra"
)


func getResponse(args [] string) {

	userArgs := strings.Join(args[0:]," ")
    ctx := context.Background()
  client, _ := genai.NewClient(ctx, &genai.ClientConfig{
      APIKey:  os.Getenv("GEMINI_API_KEY"),
      Backend: genai.BackendGeminiAPI,
  })

  stream := client.Models.GenerateContentStream(
      ctx,
      "gemini-2.0-flash",
      genai.Text(userArgs),
      nil,
  )

  for chunk := range stream {
      part := chunk.Candidates[0].Content.Parts[0]
      fmt.Print(part.Text)
  }
   
}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Args:cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		getResponse(args)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
