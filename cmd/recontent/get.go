package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/recontentapp/cli/pkg/client"
	"github.com/recontentapp/cli/pkg/io/stdout"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get <resource>",
	Short: "Display one or many resources",
	Long:  "Prints a table of the most important information about the specified resources.",
	Example: heredoc.Doc(`
		List all projects
		$ recontent get projects
		List workspace languages
		$ recontent get languages
		List project languages
		$ recontent get languages -p <project_id>
	`),
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a resource argument")
		}

		validResource := false

		for _, resource := range getPossibleResources {
			if args[0] == resource {
				validResource = true
			}
		}

		if validResource == false {
			return errors.New(fmt.Sprintf(`resource "%s" is invalid`, args[0]))
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		resource := args[0]
		rows := [][]string{}

		switch resource {
		case "projects":
			{
				res, err := httpClient.GetProjectsWithResponse(context.Background(), &client.GetProjectsParams{})
				if res == nil || res.JSON200 == nil {
					return err
				}

				if len(res.JSON200.Data) == 0 {
					fmt.Println("No projects found")
					return nil
				}

				for _, value := range res.JSON200.Data {
					rows = append(rows, []string{value.Id, value.Name, value.CreatedAt.Format(time.RFC1123)})
				}

				stdout.RenderTable([]string{"ID", "NAME", "CREATED AT"}, rows)
			}
		case "languages":
			{
				params := client.GetLanguagesParams{}
				if len(projectID) > 0 {
					params.ProjectId = &projectID
				}

				res, err := httpClient.GetLanguagesWithResponse(context.Background(), &params)
				if res == nil || res.JSON200 == nil {
					return err
				}

				if len(res.JSON200.Data) == 0 {
					fmt.Println("No languages found")
					return nil
				}

				for _, value := range res.JSON200.Data {
					rows = append(rows, []string{value.Id, value.Name, value.Locale, value.CreatedAt.Format(time.RFC1123)})
				}

				stdout.RenderTable([]string{"ID", "NAME", "LOCALE", "CREATED AT"}, rows)
			}
		}

		return nil
	},
}
