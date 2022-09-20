package waiting

import (
	"fmt"
	"github.com/abieberbach/docker-compose-dependsOn/config"
	"strings"
	"time"
)

type CheckTask func(timeout time.Duration) bool

type DependencyManager struct {
	env        *config.Environment
	checkTasks []CheckTask
}

func NewFromEnv(env *config.Environment) *DependencyManager {
	return &DependencyManager{
		env:        env,
		checkTasks: createCheckTasksFromEnv(env),
	}
}

func createCheckTasksFromEnv(env *config.Environment) []CheckTask {
	result := make([]CheckTask, 0)
	for _, host := range env.WaitHosts {
		result = append(result, NewPortCheck(strings.TrimSpace(host)))
	}
	for _, url := range env.WaitUrls {
		result = append(result, NewUrlCheck(strings.TrimSpace(url)))
	}
	for _, redisKey := range env.WaitRedisKeys {
		result = append(result, NewRedisCheck(strings.TrimSpace(redisKey), env.WaitRedisAddr, env.WaitRedisPassword, env.WaitRedisDB))
	}
	for _, dbKey := range env.WaitDBKeys {
		result = append(result, NewDatabaseCheck(strings.TrimSpace(dbKey), env.WaitDBDriver, env.WaitDBUrl, env.WaitDBUser, env.WaitDBPassword))
	}
	return result
}

func (depManager *DependencyManager) WaitForDependencies() {
	totalDuration := 0 * time.Second
	if depManager.env.WaitBeforeHosts > 0 {
		time.Sleep(depManager.env.WaitBeforeHosts)
	}
	for _, task := range depManager.checkTasks {
		for {
			if totalDuration > depManager.env.WaitHostsTimeout {
				panic(fmt.Sprintf("hosts timeout reached: %v > %v", totalDuration.Seconds(), depManager.env.WaitHostsTimeout))
			}
			result := task(depManager.env.WaitHostConnectTimeout)
			if result {
				break
			}
			time.Sleep(depManager.env.WaitSleepInterval)
		}
		fmt.Println("----------------------------")
	}
	if depManager.env.WaitAfterHosts > 0 {
		time.Sleep(depManager.env.WaitAfterHosts)
	}
}
