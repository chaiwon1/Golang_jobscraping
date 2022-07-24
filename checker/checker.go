package checker

import (
	"errors"
	"fmt"
	"net/http"
)

var errCantConnect = errors.New("Can't Connect")

// HitURL func
func HitURL(url string) error {
	fmt.Println("Checking : ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Print(err, resp.StatusCode)
	}
	fmt.Println("Connected!")
	return nil
}
