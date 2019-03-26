package agent

import (
	"fmt"
	"os/exec"
)

func logout(username string) {
	//TODO Write me
}

func restart() {
	if err := exec.Command("cmd", "/C", "shutdown", "/r", "/f", "/t", "0").Run(); err != nil {
		fmt.Println("Failed to initiate restart:", err)
	}
}

func shutdown() {
	if err := exec.Command("cmd", "/C", "shutdown", "/s", "/f", "/t", "0").Run(); err != nil {
		fmt.Println("Failed to initiate shutdown:", err)
	}
}
