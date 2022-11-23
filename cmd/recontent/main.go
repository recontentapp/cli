package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/recontentapp/cli/pkg/client"
	"github.com/recontentapp/cli/pkg/config"
	"github.com/spf13/cobra"
)

var options config.Config
var httpClient client.ClientWithResponses
var getPossibleResources = []string{"projects", "languages"}
var exportPossibleResources = []string{"phrases"}

var projectID string
var revisionID string
var languageID string
var outputFormat string
var filenameFormat string

var rootCmd = &cobra.Command{
	Use:     "recontent",
	Version: "0.1.0",
	Short:   "How product teams manage localized content",
}

func main() {
	opts, err := config.New()
	if err != nil {
		fmt.Println(heredoc.Doc(`
		Config is invalid
		Make sure to have RECONTENT_API_KEY in your path
		$ export RECONTENT_API_KEY=example
	`))
		os.Exit(1)
	}
	options = *opts

	client, err := client.NewClientWithResponses(options.APIURL, client.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+options.Token)
		return nil
	}))

	if err != nil {
		fmt.Println("Could not initialize HTTP client")
		os.Exit(1)
	}
	httpClient = *client

	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&projectID, "project", "p", "", "Project id")

	rootCmd.AddCommand(exportCmd)
	exportCmd.Flags().StringVarP(&projectID, "project", "p", "", "Project id")
	exportCmd.Flags().StringVarP(&revisionID, "revision", "r", "", "Revision id")
	exportCmd.Flags().StringVarP(&languageID, "language", "l", "", "Language id")
	exportCmd.Flags().StringVarP(&outputFormat, "output", "o", "json", "Output format")
	exportCmd.Flags().StringVarP(&filenameFormat, "filename", "f", `{{.LanguageKey}}.{{.FormatExtension}}`, "Filename format")
	exportCmd.MarkFlagRequired("project")

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
