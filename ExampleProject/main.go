package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const userFile = " "

type User struct {
	Username string
	Password string
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n1. Sign Up")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			signUp(reader)
		case "2":
			login(reader)
		case "3":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func signUp(reader *bufio.Reader) {
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	if userExists(username) {
		fmt.Println("Username already exists. Please choose a different username.")
		return
	}

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	err := saveUser(username, password)
	if err != nil {
		fmt.Println("Error saving user:", err)
		return
	}

	fmt.Println("Sign up successful!")
}

func login(reader *bufio.Reader) {
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	if authenticateUser(username, password) {
		fmt.Println("Login successful! Welcome,", username)
	} else {
		fmt.Println("Invalid username or password.")
	}
}

func userExists(username string) bool {
	file, err := os.Open(userFile)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		if parts[0] == username {
			return true
		}
	}

	return false
}

func saveUser(username, password string) error {
	file, err := os.OpenFile(userFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s:%s\n", username, password))
	return err
}

func authenticateUser(username, password string) bool {
	file, err := os.Open(userFile)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		if parts[0] == username && parts[1] == password {
			return true
		}
	}

	return false
}
