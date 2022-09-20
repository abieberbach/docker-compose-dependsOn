package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Environment struct {
	WaitHosts              []string
	WaitUrls               []string
	WaitRedisKeys          []string
	WaitDBKeys             []string
	WaitHostsTimeout       time.Duration
	WaitHostConnectTimeout time.Duration
	WaitBeforeHosts        time.Duration
	WaitAfterHosts         time.Duration
	WaitSleepInterval      time.Duration
	WaitRedisAddr          string
	WaitRedisPassword      string
	WaitRedisDB            int
	WaitDBUrl              string
	WaitDBDriver           string
	WaitDBUser             string
	WaitDBPassword         string
}

func CreateFromEnvVariables() *Environment {
	fmt.Println("Configuration from environment:")
	result := &Environment{
		WaitHosts:              []string{},
		WaitUrls:               []string{},
		WaitRedisKeys:          []string{},
		WaitDBKeys:             []string{},
		WaitHostsTimeout:       30 * time.Second,
		WaitHostConnectTimeout: 5 * time.Second,
		WaitBeforeHosts:        0 * time.Second,
		WaitAfterHosts:         0 * time.Second,
		WaitSleepInterval:      1 * time.Second,
		WaitRedisAddr:          "",
		WaitRedisPassword:      "",
		WaitRedisDB:            0,
		WaitDBUrl:              "",
		WaitDBDriver:           "",
		WaitDBUser:             "",
		WaitDBPassword:         "",
	}
	readStrings("WAIT_HOSTS", &result.WaitHosts)
	readStrings("WAIT_URLS", &result.WaitUrls)
	readStrings("WAIT_REDIS_KEYS", &result.WaitRedisKeys)
	readStrings("WAIT_DB_KEYS", &result.WaitDBKeys)
	readDuration("WAIT_HOSTS_TIMEOUT", &result.WaitHostsTimeout)
	readDuration("WAIT_HOST_CONNECT_TIMEOUT", &result.WaitHostConnectTimeout)
	readDuration("WAIT_BEFORE_HOSTS", &result.WaitBeforeHosts)
	readDuration("WAIT_AFTER_HOSTS", &result.WaitAfterHosts)
	readDuration("WAIT_SLEEP_INTERVAL", &result.WaitSleepInterval)
	readString("WAIT_REDIS_ADDR", &result.WaitRedisAddr)
	readString("WAIT_REDIS_PASSWORD", &result.WaitRedisPassword)
	readInt("WAIT_REDIS_DB", &result.WaitRedisDB)
	readString("WAIT_DB_URL", &result.WaitDBUrl)
	readString("WAIT_DB_DRIVER", &result.WaitDBDriver)
	readString("WAIT_DB_USER", &result.WaitDBUser)
	readString("WAIT_DB_PASSWORD", &result.WaitDBPassword)
	return result
}

func readStrings(envVariableName string, result *[]string) {
	value, found := os.LookupEnv(envVariableName)
	if found {
		fmt.Println(envVariableName, "=", value)
		*result = strings.Split(value, ",")
	}
}

func readString(envVariableName string, result *string) {
	value, found := os.LookupEnv(envVariableName)
	if found {
		fmt.Println(envVariableName, "=", value)
		*result = value
	}
}

func readInt(envVariableName string, result *int) {
	stringValue, found := os.LookupEnv(envVariableName)
	if found {
		fmt.Println(envVariableName, "=", stringValue)
		value, err := strconv.Atoi(stringValue)
		if err != nil {
			panic(err)
		}
		*result = value
	}
}

func readDuration(envVariableName string, result *time.Duration) {
	value, found := os.LookupEnv(envVariableName)
	if found {
		fmt.Println(envVariableName, "=", value)
		var err error
		*result, err = time.ParseDuration(value + "s")
		if err != nil {
			fmt.Println(envVariableName, "=", value)
			panic(err)
		}
	}
}
