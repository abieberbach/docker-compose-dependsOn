package main

import (
	"fmt"
	"github.com/abieberbach/docker-compose-dependsOn/config"
	"github.com/abieberbach/docker-compose-dependsOn/waiting"
)

func main() {
	fmt.Println("docker-compose-dependsOn")
	fmt.Println("============================")
	env := config.CreateFromEnvVariables()
	fmt.Println("============================")
	fmt.Println("")
	dependencyManager := waiting.NewFromEnv(env)
	dependencyManager.WaitForDependencies()
}
