package config

import ("fmt"
        "log"
	     "os"
		 "github.com/joho/godotenv")


type Config struct {

	AppPort  string
	AppEnv   string
	DB       DBConfig
}


type DBConfig struct {

	Host      string
	Port      string 
	User      string 
	Password  string
	Name      string
	SSLMode   string
}


func (d DBConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		d.Host, d.User, d.Password, d.Name, d.Port, d.SSLMode,
	)
}

func Load() *Config  {

	// load env file locally 

	if err := godotenv.Load(); err != nil{
		log.Println(".env file not found , using system enviroment")
	}

	cfg := &Config{
		AppPort: mustEnv("APP_PORT"),
		AppEnv:  mustEnv("APP_ENV"),
		DB: DBConfig{
			Host:     mustEnv("DB_HOST"),
			Port:     mustEnv("DB_PORT"),
			User:     mustEnv("DB_USER"),
			Password: mustEnv("DB_PASSWORD"),
			Name:     mustEnv("DB_NAME"),
			SSLMode:  mustEnv("DB_SSLMODE"),
		},
	}

	return cfg
}




// mustEnv fails fast if env missing
func mustEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Missing required env: %s", key)
	}
	return val
}