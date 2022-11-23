package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/recontentapp/cli/pkg/client"
	"github.com/recontentapp/cli/pkg/io/fileformat"
	"github.com/recontentapp/cli/pkg/io/filename"
	"github.com/recontentapp/cli/pkg/io/json"
	"github.com/recontentapp/cli/pkg/io/raw"
	"github.com/recontentapp/cli/pkg/io/yaml"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export <resource>",
	Short: "Export translations",
	Long:  "Downloads translations for a specific resource and outputs them in a chosen format.",
	Example: heredoc.Doc(`
		Export phrases translations to a simple JSON file
		$ recontent export phrases -p <project_id> -o json
		Export phrases translations for a given language & revision
		$ recontent export phrases -p <project_id> -l <language_id> -r <revision_id>
		Export phrases using a custom filename format
		$ recontent export phrases -p <project_id> -f "i18n/{{.LanguageKey}}.{{.FormatExtension}}"
	`),
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a resource argument")
		}

		validResource := false
		for _, resource := range exportPossibleResources {
			if args[0] == resource {
				validResource = true
			}
		}
		if validResource == false {
			return errors.New(fmt.Sprintf(`Resource "%s" is invalid`, args[0]))
		}

		_, err := fileformat.New(outputFormat)
		if err != nil {
			return errors.New(fmt.Sprintf(`Output file format "%s" is invalid for resource "%s"`, outputFormat, args[0]))
		}

		filenameFormatValid := filename.IsValid(filenameFormat)
		if filenameFormatValid == false {
			return errors.New(fmt.Sprintf(`Filename format "%s" is invalid`, filenameFormat))
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		languagesRes, err := httpClient.GetLanguagesWithResponse(context.Background(), &client.GetLanguagesParams{
			ProjectId: &projectID,
		})

		if languagesRes == nil || languagesRes.JSON200 == nil {
			return err
		}

		for _, language := range languagesRes.JSON200.Data {
			if len(languageID) > 0 && language.Id != languageID {
				continue
			}

			res, err := httpClient.GetPhrasesExportWithResponse(context.Background(), &client.GetPhrasesExportParams{
				ProjectId:  projectID,
				LanguageId: language.Id,
			})
			if res == nil || res.JSON200 == nil {
				return err
			}

			fileFormat := fileformat.Fileformat(outputFormat)

			switch fileFormat {
			case fileformat.FileformatJSON,
				fileformat.FileformatNestedJSON:
				{
					filename, err := filename.Render(filenameFormat, filename.Variables{
						LanguageLocale:  language.Locale,
						LanguageName:    language.Name,
						FormatExtension: "json",
					})

					if err != nil {
						return err
					}

					var data []byte

					if fileFormat == fileformat.FileformatNestedJSON {
						output, err := json.BuildNested(res.JSON200.Data)
						if err != nil {
							return err
						}
						data = output
					} else {
						output, err := json.Build(res.JSON200.Data)
						if err != nil {
							return err
						}
						data = output
					}

					err = raw.Write(filename, data)
					if err != nil {
						return err
					}
				}

			case fileformat.FileformatYAML,
				fileformat.FileformatNestedYAML:
				{
					filename, err := filename.Render(filenameFormat, filename.Variables{
						LanguageLocale:  language.Locale,
						LanguageName:    language.Name,
						FormatExtension: "yaml",
					})

					if err != nil {
						return err
					}

					var data []byte

					if fileFormat == fileformat.FileformatNestedYAML {
						output, err := yaml.BuildNested(res.JSON200.Data)
						fmt.Println(err)
						if err != nil {
							return err
						}
						data = output
					} else {
						output, err := yaml.Build(res.JSON200.Data)
						if err != nil {
							return err
						}
						data = output
					}

					err = raw.Write(filename, data)
					if err != nil {
						return err
					}
				}
			}
		}

		return nil
	},
}
