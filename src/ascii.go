package shell

import (
	"fmt"
	"os"
	"runtime"
	"strings"
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

func retrieveDistributionName() string {
	if runtime.GOOS == "windows" {
		return "Windows"
	} else if runtime.GOOS == "linux" {
		return "Linux"
	} else {
		return "MacOS"
	}
}

func Ascii() {
	username := getUsername()
	distro := retrieveDistributionName()

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
