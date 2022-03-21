package main

import (
	"fmt"
	"os"
)

// dlv debug --headless --listen=:2345

func main() {
	var cmdline string
	fmt.Println("WelCome to The Machine")
	fmt.Println("#*********# Go_MENU v0.1 #*********#")
	fmt.Println("help: User Manual || list: List Page")
	fmt.Println("setup: Setup Page || quit: QUIT")
	fmt.Println("And some hidden words")
	fmt.Println("####################################")
	for true {
		fmt.Println("\n#*********# Go_MENU v0.1 #*********#")
		fmt.Scan(&cmdline)

		switch cmdline {
		case "help":
			fmt.Println("How can I help you?")
			break
		case "list":
			fmt.Println("Hurry up to make your choise!")
			break
		case "setup":
			fmt.Println("Watch up where you going")
			break
		case "quit":
			fmt.Println("OK OK I'm done.")
			os.Exit(0)
		case "canyouhearme":
			fmt.Println("Hell yes")
			break
		case "nut":
			fmt.Println("I could be bounded in a nutshell and count myself a king of infinite space.")
			fmt.Println("https://github.com/phantomT/T-Shell")
			break
		default:
			fmt.Println("HAHA, VERY FUNNY, DO THAT AGAIN")
		}
	}
}
