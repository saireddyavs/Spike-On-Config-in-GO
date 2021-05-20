package vault

import (
	"log"

	"github.com/hashicorp/vault/api"
	"github.com/spf13/viper"
)

type VaultConfig struct {
	vaultEnabled bool
	vaultAddress string
	vaultToken   string
}

func (v VaultConfig) GetVaultSecrets() {
	config := &api.Config{
		Address: v.vaultAddress,
	}
	client, err := api.NewClient(config)
	if err != nil {
		log.Println(err)
		return
	}
	client.SetToken(v.vaultToken)

	secret, err := client.Logical().Read("NameAndId/data/details")
	if err != nil {
		log.Println(err)
		return
	}
	m, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		log.Printf("Reading secrets failed")
		return
	}

	for k, v := range m {
		viper.Set(k, v)
	}
}


func NewVaultConfig() VaultConfig {
	return VaultConfig{
		vaultEnabled: getBool("VAULT_ENABLED"),
		vaultAddress: getString("VAULT_ADDRESS"),
		vaultToken:   getString("VAULT_TOKEN"),
	}
}
