package main

import (
	"log/slog"
	"os"
	"testing"
)

func TestLolo(t *testing.T) {

	//handlerJSON := slog.NewJSONHandler(os.Stdout, nil)
	handlerTXT := slog.NewTextHandler(os.Stdout, nil)
	log := slog.New(handlerTXT)
	log.Info("TestLolo", "post_id", "234234")

}
