package util

import (
	"fmt"
	"os"
	"strings"
)

var (
	distro string
)

func version() string {
	data, err := os.ReadFile(".version")
	if err != nil {
		fmt.Println("Error: .version file not found.")
		return "Unknown"
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "VERSION=") {
			return strings.TrimPrefix(line, "VERSION=")
		}
	}
	fmt.Println("Error: Version not found in .version file.")
	return "Unknown"
}

func getUsername() string {
	username := os.Getenv("USER")
	if username == "" {
		username = os.Getenv("LOGNAME")
	}
	if username == "" {
		username = os.Getenv("USERNAME")
	}
	if username == "" {
		username = "Unknown"
	}
	return username
}

func retrieveDistributionName(distro *string) {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		*distro = "Unknown"
		return
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "NAME=") {
			nameStart := strings.Index(line, "=")
			if nameStart != -1 {
				nameStart++
				if line[nameStart] == '"' {
					nameStart++
					nameEnd := strings.Index(line[nameStart:], "\"")
					if nameEnd != -1 {
						*distro = line[nameStart : nameStart+nameEnd]
						return
					}
				}
			}
		}
	}
	*distro = "Unknown"
}

func Ascii() {
	username := getUsername()
	retrieveDistributionName(&distro)

	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣤⡀⠀⠀⠀⠀⠀")
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⣧⠃⡃⠀⠀⠀⠀⠀")
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⢿⠇⠏⡇⠀⠀⠀⠀")
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣟⠏⡜⠰⢣⠀⠀⠀⠀")
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡴⠟⣡⠊⡠⢁⣎⠀⠀⠀⠀")
	fmt.Println("⠀⠀⢀⣶⣰⡂⠀⠀⠀⠀⣠⠖⠋⡔⠊⡠⢏⡠⢃⡜⠀⠀⠀⠀")
	fmt.Println("⠀⠀⡧⣩⠾⢹⣓⣤⠀⠘⣯⡀⠀⣗⣋⣤⢃⠔⢩⠆⠀⠀⠀⠀")
	fmt.Println("⠀⣸⣅⠁⢁⡞⠨⡍⢷⢂⡾⠁⣰⠡⠤⣊⠥⠒⠁⠀⠀⠀⠀⠀")
	fmt.Println("⠘⠿⠽⠞⢫⠀⠀⠀⢻⠟⣡⡾⠱⡑⠦⠥⣀⣀⣀⠀⠀⠀⠀⠀")
	fmt.Println("⠀⠀⠀⠀⡎⢀⠐⠀⠁⠈⠺⠷⠣⠃⠁⡀⠈⢿⡇⠨⢢⠀⠀⠀")
	fmt.Println("⠀⠀⠀⢀⡃⠀⠐⠀⠂⠐⠂⢀⠂⠂⠐⢀⠀⡾⢔⢄⠡⣃⠀⠀")
	fmt.Printf("⠀⢸⡏⠃⠀⠙⣞⡆⠀⠀⠀⠀⠀⠀⣵⡚⠱⣎⣇⠀⠀   username: %s\n", username)
	fmt.Printf("⠀⣿⣷⠄⠀⠀⢹⣽⠀⠀⠀⠀⠀⣼⣳⠃⠀⠹⣾⣤⠀   distro: %s\n", distro)
	fmt.Printf("⠀⠉⠁⠀⠀⠰⠿⠟⠀⠀⠀⠀⠘⠛⠛⠀⠀⠀⠹⠋    version: %s\n", version())
}
