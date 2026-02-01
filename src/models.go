package src

type MakeTarget struct {
	Name        string
	Description string
	IsPhony     bool
}

type Config struct {
	MakefilePath string
	DefaultArgs  []string
}

func GetDefaultConfig() *Config {
	return &Config{
		MakefilePath: "Makefile",
		DefaultArgs:  []string{},
	}
}
