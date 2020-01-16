package config

import (
	"fmt"
	"github.com/ifaceless/go-starter/pkg/util/toolkit/resource"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type DBConfig struct {
	Master string
	Slaves []string
}

var (
	MySQLConfig *DBConfig
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Warnf("failed to load .env")
	}

	res, err := resource.DiscoverMySQL("go-starter")
	if err != nil {
		panic(err)
	}

	slaveAddrs := make([]string, 0)
	for _, ins := range res.SlaveInstances {
		slaveAddrs = append(slaveAddrs, makeDBAddr(ins.Addr()))
	}
	MySQLConfig = &DBConfig{
		Master: makeDBAddr(res.MasterInstances[0].Addr()),
		Slaves: slaveAddrs,
	}
}

func makeDBAddr(src string) string {
	return fmt.Sprintf("%s?charset=utf8&parseTime=true&loc=Local", src)
}