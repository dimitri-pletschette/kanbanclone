// This file is automatically generated. Do not modify it manually.

package main

import (
	"encoding/json"
	"strings"

	"github.com/mattermost/mattermost-server/v6/model"
)

var manifest *model.Manifest

const manifestStr = `
{
  "id": "focalboard",
  "name": "Mattermost Boards",
  "description": "The Mattermost Boards plugin",
  "homepage_url": "https://github.com/mattermost/focalboard",
  "support_url": "https://github.com/mattermost/focalboard/issues",
  "release_notes_url": "https://github.com/mattermost/focalboard/releases",
  "icon_path": "assets/starter-template-icon.svg",
  "version": "7.12.0",
  "min_server_version": "7.2.0",
  "server": {
    "executables": {
      "darwin-amd64": "server/dist/plugin-darwin-amd64",
      "darwin-arm64": "server/dist/plugin-darwin-arm64",
      "linux-amd64": "server/dist/plugin-linux-amd64",
      "linux-arm64": "server/dist/plugin-linux-arm64",
      "windows-amd64": "server/dist/plugin-windows-amd64.exe"
    },
    "executable": ""
  },
  "webapp": {
    "bundle_path": "webapp/dist/main.js"
  },
  "settings_schema": {
    "header": "",
    "footer": "",
    "settings": [
      {
        "key": "EnablePublicSharedBoards",
        "display_name": "Enable Publicly-Shared Boards:",
        "type": "bool",
        "help_text": "This allows board editors to share boards that can be accessed by anyone with the link.",
        "placeholder": "",
        "default": false,
        "hosting": ""
      }
    ]
  }
}
`

func init() {
	_ = json.NewDecoder(strings.NewReader(manifestStr)).Decode(&manifest)
}
