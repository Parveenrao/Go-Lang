package config

import "os"

// Db config holds all database config settings 
// we use a struct of instead of raw dsn file  so that:
    // each file  is named and obvious 
	// we can validate fiels if needed


type DBConfig struct{
	Host       string 
	Port       string 
	User       string 
	Password   string 
	DBName     string 
	SSLMode    string 
}

// ServerConfig holds HTTP server settings.
type ServerConfig struct {
	Port string
}


// AppConfig is the root config — holds all sub-configs.
// main.go loads this once and passes pieces down.
type AppConfig struct {
	DB     DBConfig            // strucut embeddngs 
	Server ServerConfig
}


// Load read all config value from enviorment variables 

func Load() AppConfig {

	return  AppConfig{

		DB : DBConfig {
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "userapi"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},

		Server : ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
		
	}
	}


// getEnv returns the value of an env var,
// or the fallback string if the var is not set.
func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback

}



