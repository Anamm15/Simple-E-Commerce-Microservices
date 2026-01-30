package constant

import "os"

func GetJWTSecret() string {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		return "BOIDSGBOABSDGIASBDIOBASDBOBIDSBOFUAB"
	}
	return secret
}
