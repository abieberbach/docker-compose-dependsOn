package config

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Environment struct {
	WaitHosts              []string
	WaitUrls               []string
	WaitHostsTimeout       time.Duration
	WaitHostConnectTimeout time.Duration
	WaitBeforeHosts        time.Duration
	WaitAfterHosts         time.Duration
	WaitSleepInterval      time.Duration
}

func CreateFromEnvVariables() *Environment {
	fmt.Println("Configuration from environment:")
	result := &Environment{
		WaitHosts:              []string{},
		WaitUrls:               []string{},
		WaitHostsTimeout:       30 * time.Second,
		WaitHostConnectTimeout: 5 * time.Second,
		WaitBeforeHosts:        0 * time.Second,
		WaitAfterHosts:         0 * time.Second,
		WaitSleepInterval:      1 * time.Second,
	}
	readStrings("WAIT_HOSTS", &result.WaitHosts)
	readStrings("WAIT_URLS", &result.WaitUrls)
	readDuration("WAIT_HOSTS_TIMEOUT", &result.WaitHostsTimeout)
	readDuration("WAIT_HOST_CONNECT_TIMEOUT", &result.WaitHostConnectTimeout)
	readDuration("WAIT_BEFORE_HOSTS", &result.WaitBeforeHosts)
	readDuration("WAIT_AFTER_HOSTS", &result.WaitAfterHosts)
	readDuration("WAIT_SLEEP_INTERVAL", &result.WaitSleepInterval)
	return result
}

func readStrings(envVariableName string, result *[]string) {
	value, found := os.LookupEnv(envVariableName)
	if found {
		fmt.Println(envVariableName, "=", value)
		*result = strings.Split(value, ",")
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
