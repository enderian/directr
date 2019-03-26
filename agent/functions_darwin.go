package agent

import "log"

func logout(username string) {
	log.Printf("Should be logging out user %s", username)
}

func restart() {
	log.Printf("Should be restarting")
}

func shutdown() {
	log.Printf("Should be shutting down")
}
