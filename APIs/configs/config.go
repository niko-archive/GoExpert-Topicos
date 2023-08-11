package configs

import (
	"log"
	"strconv"

	"github.dev/nicolasmmb/GoExpert-Topicos/internal/entity"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

var envs *ENVs

type ENVs struct {
	// DB Configs
	DB_HOST string `mapstructure:"DB_HOST"`
	DB_PORT string `mapstructure:"DB_PORT"`
	DB_USER string `mapstructure:"DB_USER"`
	DB_PASS string `mapstructure:"DB_PASS"`
	DB_NAME string `mapstructure:"DB_NAME"`
	// JWT Configs
	JWT_SECRET string `mapstructure:"JWT_SECRET"`
	JWT_ALG    string `mapstructure:"JWT_ALG"`
	JWT_EXP    int    `mapstructure:"JWT_EXP"`
	// Server Configs
	SERVER_HOST string `mapstructure:"SERVER_HOST"`
	SERVER_PORT int    `mapstructure:"SERVER_PORT"`
}

func LoadENVs(path string) *ENVs {
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&envs)
	if err != nil {
		panic(err)
	}
	return envs
}

func PrintSeparator() {
	log.SetFlags(0)
	log.SetPrefix("")
	log.Default().Println("=============================================")
}

func (envs *ENVs) Print() {

	log.SetFlags(0)
	PrintSeparator()
	log.SetPrefix("[ENVs] ")
	log.Default().Println("DB_HOST:", envs.DB_HOST)
	log.Default().Println("DB_PORT:", envs.DB_PORT)
	log.Default().Println("DB_USER:", envs.DB_USER)
	log.Default().Println("DB_PASS:", envs.DB_PASS)
	log.Default().Println("DB_NAME:", envs.DB_NAME)
	log.Default().Println("JWT_ALG:", envs.JWT_ALG)
	log.Default().Println("JWT_EXP:", envs.JWT_EXP)
	log.Default().Println("JWT_SECRET:", envs.JWT_SECRET)
	log.Default().Println("SERVER_HOST:", envs.SERVER_HOST)
	log.Default().Println("SERVER_PORT:", envs.SERVER_PORT)
	log.SetPrefix("")
	PrintSeparator()

}

func (envs *ENVs) GetServerAddress() string {
	return envs.SERVER_HOST + ":" + strconv.Itoa(envs.SERVER_PORT)
}

func CreateAdmin(db *gorm.DB) {
	u, err := entity.NewUser("admin", "admin@mail.com", "admin")
	if err != nil {
		log.Fatal(err)
	}
	db.Create(&u)
}
