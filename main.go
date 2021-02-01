package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	aws "ecr-auth/utils"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Running at " + time.Now().UTC().String())

	fmt.Print("Fetching auth data from AWS... ")

	username, password, server, err := aws.GetUserAndPass()

	checkErr(err)

	fmt.Println("Success.")

	fmt.Println("Username -> ", username)
	fmt.Println("Token    -> ", password)
	fmt.Println("Server   -> ", server)

	fmt.Println("Loggin in")

	cmd := exec.Command("docker", "login", "--username", username, "--password", password, server)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Job complete.")
}
