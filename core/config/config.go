package config

const (
	ReleaseMode string = "release"
	DebugMode   string = "debug"
	TestMode    string = "release"
)

type Config struct {
	PostgresHost        string `envconfig:"POSTGRES_HOST" default:"localhost"`
	PostgresPort        string `envconfig:"POSTGRES_PORT" default:"5432"`
	PostgresDB          string `envconfig:"POSTGRES_DB" default:"slyfox-tails"`
	PostgresUser        string `envconfig:"POSTGRES_USER" default:"20624880"`
	PostgresPassword    string `envconfig:"POSTGRES_PASSWORD" default:"admin"`
	PostgresAutoMigrate bool   `envconfig:"POSTGRES_AUTO_MIGRATE" default:"false"`

	RedisHost     string `envconfig:"REDIS_HOST" default:"localhost"`
	RedisPort     string `envconfig:"REDIS_PORT" default:"6379"`
	RedisDB       int    `envconfig:"REDIS_DB" default:"0"`
	RedisPassword string `envconfig:"REDIS_PASSWORD" default:""`

	Mode string `envconfig:"MODE" default:"release"`
	Port string `envconfig:"PORT" default:":8080"`
}
