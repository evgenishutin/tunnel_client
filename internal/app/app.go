package app

import (
	"fmt"
	"os"
	"ssh-proxy-app/internal/service"
	"ssh-proxy-app/internal/usecase"
	"ssh-proxy-app/pkg/helpers"

	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Run() {
	myApp := fyneApp.New()

	sshProxyService := service.NewSSHProxyService()
	sshProxyUseCase := usecase.NewProxySSHUseCase(*sshProxyService)

	pid := os.Getpid()
	fmt.Printf("PID текущего процесса: %d\n", pid)

	myWindow := myApp.NewWindow("SSH Proxy App")
	myWindow.Resize(fyne.NewSize(400, 300))

	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Username")

	hostEntry := widget.NewEntry()
	hostEntry.SetPlaceHolder("Ip address (e.g., 127.0.0.1)")

	validationLabel := widget.NewLabel("")

	sshButton := widget.NewCheck("Start SSH Proxy", func(checked bool) {
		if checked {
			username := usernameEntry.Text
			host := hostEntry.Text
			if username == "" {
				fmt.Println("Username cannot be empty")
				validationLabel.SetText("Username cannot be empty")
				return
			} else {
				validationLabel.SetText("")
			}
			if !helpers.IsValidIPv4(host) {
				validationLabel.SetText("is not a valid IPv4 address.")
				fmt.Println(host, "is not a valid IPv4 address.")
			} else {
				validationLabel.SetText("")
			}

			if username != "" && helpers.IsValidIPv4(host) {
				paramsSSH := sshProxyUseCase.SetParams(username, host)
				err := sshProxyUseCase.StartProxy(*paramsSSH)
				if err != nil {
					fmt.Println("Error starting SSH Proxy:", err)
				}
			}
			myWindow.SetTitle("turning on the tunnel")
		} else {
			err := sshProxyUseCase.StopProxy()
			if err != nil {
				fmt.Println("Error stoping SSH Proxy:", err)
			}
			myWindow.SetTitle("turning off the tunnel")
		}
	})

	myWindow.SetContent(container.NewVBox(
		widget.NewLabel("SSH Proxy App"),
		usernameEntry,
		hostEntry,
		validationLabel,
		sshButton,
	))

	myWindow.ShowAndRun()
}
