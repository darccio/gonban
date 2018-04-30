package main

import (
	"fmt"
	"net/url"

	"github.com/zserge/webview"
)

func main() {
	app, err := Asset("www/dist.js")
	if err != nil {
		panic(err)
	}
	body := fmt.Sprintf(`<html>
		<head>
		<link href="https://fonts.googleapis.com/css?family=PT+Sans" rel="stylesheet">
		<style>
			body {
				background: #1d1f21;
			}
		</style>
		</head>
		<body>
		<script>
			%s
		</script>
		</body>
	</html>`, app)
	settings := webview.Settings{
		Title:     "Gonban",
		Width:     700,
		Height:    400,
		Resizable: true,
		URL:       fmt.Sprintf("data:text/html,%s", url.PathEscape(body)),
	}
	view := webview.New(settings)
	view.Run()
}
