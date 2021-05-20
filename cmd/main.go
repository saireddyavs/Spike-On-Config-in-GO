package main

import (
	"fmt"
	"go-env-ways/pkg/envfile"
	"go-env-ways/pkg/vault"
	"go-env-ways/pkg/viperenv"
	"os"
)

func envVariable(key, value string) string {
	os.Setenv(key, value)
	return os.Getenv(key)
}

func main() {

	fmt.Println(envVariable("name", "Sai Reddy"))

	fmt.Println(envfile.GoDotEnvVariable("id"))

	fmt.Println(viperenv.ViperEnvVariable("id"))

	fmt.Println(vault.GetVaultVariable("name"))

	fmt.Println(vault.GetVaultVariable("id"))

}
