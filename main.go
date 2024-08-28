package main

import (
	// "github.com/go-git/go-git/v5"

	"fmt"
	"os"

	"github.com/andyroid0/incantation/app"
	"github.com/andyroid0/incantation/cache"
	IncantConfig "github.com/andyroid0/incantation/config"
	"github.com/andyroid0/incantation/env"
	"github.com/andyroid0/incantation/logger"
	"github.com/andyroid0/incantation/scaffold"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func main() {
	env.SetupEnv()
	config := &IncantConfig.Config{}
	config.Init()
	scaffold.Go(config)
	logger, file := logger.Init(config)
	defer file.Close()
	cache.Cache{
		Config: config,
		Logger: logger,
	}.Init()

	zone.NewGlobal()

	p := tea.NewProgram(app.New(logger), tea.WithMouseCellMotion(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
