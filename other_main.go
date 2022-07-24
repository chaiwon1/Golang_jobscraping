package main

// urlchecker
import (
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

func main() {
	c := make(chan result)

	urls := []string{
		"https://www.naver.com/",
		"https://www.youtube.com/",
		"https://www.github.com/",
		"https://your-naduri.herokuapp.com/",
		"https://www.podo.com/",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}

func hitURL(url string, c chan result) {
	fmt.Println("Checking :", url)
	resp, err := http.Get(url)
	status := "Connected"

	if err != nil || resp.StatusCode >= 400 {
		status = "Failed"
	}
	c <- result{url: url, status: status}
}

// mydict
import (
	"fmt"
	"github.com/chaiwon/learngo/mydict"
)

func main() {
	dictionary := mydict.Sajeon{"Hungry": "Carrot", "Bored": "Take a walk"}
	fmt.Println(dictionary)
	dictionary.Delete("Hungry")
	fmt.Println(dictionary)
}

// account
import (
	"fmt"
	"github.com/chaiwon/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("podo")
	fmt.Println(account.Owner())
	account.ChangeOwner("namu")
	fmt.Println(account.Owner())
}
