package main

import (
	"fmt"
	"net/http"
	"os"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/httpapi"
)

func main() {
	api := httpapi.NewAPI(http.NewServeMux(), domain.Meta{Version: "1.0.0"})
	spec, err := api.OpenAPI().YAML()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Stdout.Write(spec)
}
