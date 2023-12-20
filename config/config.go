package config

type Config struct {
	App 		App 		`json:"app"`
	Database 	Database 	`json:"database"`
}

type App struct {
	Port string `json:"port"`
}

type Database struct {
	Engine      string  `json:"engine"`
	Server		string	`json:"server"`
	Port 		string	`json:"port"`
	User 		string 	`json:"user"`
	Password 	string 	`json:"password"`
	Name 		string 	`json:"name"`
}

