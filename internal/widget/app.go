package widget

import (
	"dolphinPeek/internal/config"
	"dolphinPeek/internal/screens"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	currentScreen int
	content       *fyne.Container
	label         *widget.Label
	screens       []func() string
}

func New() *App {
	label := widget.NewLabel("")
	label.TextStyle.Monospace = true

	return &App{
		currentScreen: 0,
		label:         label,
		content:       container.NewBorder(nil, nil, nil, nil, label),
		screens: []func() string{
			screens.GetNeofetchInfo,
			screens.GetWeatherInfo,
			// screens.GetAnimationFrame, // TODO: implement
			screens.GetRandomJoke, // Random dolphin wisdom
		},
	}
}

func (a *App) GetContent() *fyne.Container {
	return a.content
}

func (a *App) Start() {
	ticker := time.NewTicker(config.GetCycleInterval())
	defer ticker.Stop()

	// Show first screen immediately
	a.updateScreen()

	for range ticker.C {
		a.currentScreen = (a.currentScreen + 1) % len(a.screens)
		a.updateScreen()
	}
}

func (a *App) updateScreen() {
	content := a.screens[a.currentScreen]()
	fyne.DoAndWait(func() {
		a.label.SetText(content)
	})
}
