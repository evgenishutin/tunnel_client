package app

import (
	config "ssh-proxy-app/config"
	"ssh-proxy-app/internal/usecase"

	proxy "ssh-proxy-app/pkg/proxy"

	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Run(conf config.Config) {
	myApp := fyneApp.New()
	newProxy := proxy.NewProxy(conf)
	proxyUseCase := usecase.NewProxyUseCase(newProxy)

	// pid := os.Getpid()
	// fmt.Printf("PID текущего процесса: %d\n", pid)

	myWindow := myApp.NewWindow("Proxy tunnel")
	myWindow.Resize(fyne.NewSize(400, 300))
	validationLabel := widget.NewLabel("")

	startButton := widget.NewButton("Connect", func() {
		proxyUseCase.StartProxy()
	})

	myWindow.SetContent(container.NewVBox(
		widget.NewLabel("Proxy tunnel"),
		validationLabel,
		startButton,
	))

	myWindow.ShowAndRun()
}
