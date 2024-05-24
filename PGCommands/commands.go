package main

type Command struct {
	Cmd         string
	Description string
}

var commands = map[string]Command{
	"Run UR":                 {Cmd: "cd ~/Developer/platform/source/controlling/Sources/Platform.Web; pnpm run ur", Description: "Update the platfor GraphQL and Relay"},
	"Sync translations":      {Cmd: "cd ~/Developer/platform/source/controlling/Sources/Platform.Web; pnpm run sync-translations", Description: "Sync the translation files"},
	"Update controlling api": {Cmd: "cd ~/Developer/platform/source/common/Sources; pnpm run update-controlling-api", Description: "Run autorst to generate Api client functions for Platform.Web"},
}
