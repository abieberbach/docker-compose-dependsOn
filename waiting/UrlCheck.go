package waiting

import (
	"fmt"
	"net/http"
	"time"
)

func NewUrlCheck(url string) CheckTask {
	return func(timeout time.Duration) bool {
		client := http.Client{
			Timeout: timeout,
		}
		fmt.Print("Checking url: ", url)
		response, _ := client.Get(url)
		var result bool
		if response != nil {
			if response.Body != nil {
				response.Body.Close()
			}
			result = response.StatusCode == http.StatusOK
		}
		fmt.Println("", result)
		return result
	}
}
