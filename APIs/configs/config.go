package configs

import (
	"log"
	"os"
	"reflect"
	"strconv"

	"github.dev/nicolasmmb/GoExpert-Topicos/internal/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

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
	// DEBUG
	IS_DEBUG bool `mapstructure:"IS_DEBUG"`
}

func LoadENVs(path string) *ENVs {
	v := viper.New()

	if envs != nil {
		return envs
	}

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Fatal("No ENV config file found")
	}

	v.SetConfigFile(".env")
	v.AddConfigPath(path)
	v.AutomaticEnv()

	err = v.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = v.Unmarshal(&envs)
	if err != nil {
		log.Fatal(err)
	}

	return envs
}

func LoadENV() *ENVs {
	envs := &ENVs{}
	val := reflect.ValueOf(envs).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i).Tag.Get("mapstructure")
		value := os.Getenv(field)
		switch val.Field(i).Kind() {
		case reflect.String:
			val.Field(i).SetString(value)
		case reflect.Int:
			intValue, err := strconv.Atoi(value)

			if err != nil {
				panic(err)
			}
			val.Field(i).SetInt(int64(intValue))
		case reflect.Bool:
			boolValue, err := strconv.ParseBool(value)
			if err != nil {
				panic(err)
			}
			val.Field(i).SetBool(boolValue)
		case reflect.Struct:
			panic("Structs not supported yet")
		default:
			panic("Type not supported yet")
		}
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
	log.Default().Println("IS_DEBUG:", envs.IS_DEBUG)
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
	err = db.Create(&u).Error
	if err != nil {
		log.Default().Println("Admin already exists")
	}
}

func (envs *ENVs) GetLoggerType() logger.LogLevel {
	logType := logger.Silent
	if envs.IS_DEBUG {
		log.Println("Debug Mode: ON")
		logType = logger.Info
	} else {
		log.Println("Debug Mode: OFF")
	}
	return logType
}
