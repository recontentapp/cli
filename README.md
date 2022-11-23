# Recontent.app CLI

> Ship localized content faster without loosing engineers time

## Installation

Recontent CLI, or `recontent`, is a command-line interface to Recontent for use in your terminal or your scripts.

It can be installed through the following links:

TO DO

## Configuration

Define an environment variable named `RECONTENT_API_KEY` to authenticate with your Recontent workspace.

An API key can be generated from your workspace settings.

## Using the CLI

```sh
# List projects
recontent get projects

# List languages
recontent get languages

# List languages within a project
recontent get languages -p <project_id>
```

```sh
# Export phrases & translations in all languages as JSON files
# Possible outputs include json|json_nested|yaml|yaml_nested
recontent export phrases -p <project_id> -o json

# Export phrases & translations for a specific language within a revision
recontent export phrases -p <project_id> -l <language_id> -r <revision_id>

# Export phrases & translations in all languages with a custom file structure
# Possible format variables include LanguageKey|LanguageName|FormatExtension
recontent export phrases -p <project_id> -f "i18n/{{.LanguageKey}}.{{.FormatExtension}}"
```

## Contributing

When pulling updates, make sure to include ones from submodules using:

```sh
git pull --recurse-submodules

# Update submodules
git submodule update --remote
```

## Build the client

Prerequisites: 
- [`swagger-cli`](https://github.com/APIDevTools/swagger-cli)
- [`oapi-codegen`](https://github.com/deepmap/oapi-codegen)

```sh
make build-client
```
