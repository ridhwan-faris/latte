package latte

type Configuration interface {
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	Initialize(configPath string)
}
