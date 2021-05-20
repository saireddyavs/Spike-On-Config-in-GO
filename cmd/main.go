package main

import (
	"fmt"
	"go-config/pkg/envfile"
	"go-config/pkg/vault"
	"go-config/pkg/viperenv"
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
