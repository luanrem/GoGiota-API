package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

const (
	migrationsPath = "./internal/store/pgstore/migrations"
	configPath     = "./internal/store/pgstore/migrations/tern.conf"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	args := []string{
		"migrate",
		"--migrations", migrationsPath,
		"--config", configPath,
	}

	args = append(args, os.Args[1:]...) // <-- repassa o que digitar depois de chamar o script

	cmd := exec.Command(
		"tern",
		args...,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Command execution failed:", err)
		fmt.Println("Output: ", string(output))
		panic(err)
	}

	fmt.Println("Command executed successfully", string(output))
}
