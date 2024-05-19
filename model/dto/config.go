package dto

type (
	ConfigData struct {
		DbConfig  DbConfig
		AppConfig AppConfig
	}

	DbConfig struct {
		Host        string
		DbPort      string
		User        string
		Pass        string
		Database    string
		MaxIdle     int
		MaxCounn    int
		MaxLifeTime string
		LogMode     int
	}

	AppConfig struct {
		Port string
	}
)
