package agent

import "os/exec"

func logout(username string) {
	//First try for GNOME.
	if err := exec.Command("gnome-session-quit", "--force").Run(); err != nil {
		//Then native linux.
		if err := exec.Command("killall", "-9", "-u", username).Run(); err != nil {
			fmt.Println("Failed to initiate restart:", err)
		}
	}
}

func restart() {
	if err := exec.Command("reboot").Run(); err != nil {
		fmt.Println("Failed to initiate restart:", err)
	}
}

func shutdown() {
	if err := exec.Command("shutdown", "0").Run(); err != nil {
		fmt.Println("Failed to initiate shutdown:", err)
	}
}
