package vault

func GetVaultVariable(key string) string {
	vaultConfig := NewVaultConfig()

	vaultConfig.GetVaultSecrets()

	return getString(key, "dummy")
}
