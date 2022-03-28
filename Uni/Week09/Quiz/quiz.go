package main

import (
  "fmt"
  "sync"
  "io/ioutil"
  "net/http"
)

func main() {

	arrayOfPageCode := []string{}

	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"https://aiesec.org/",
	}
	
	fmt.Println("Let's start executing queries")

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			in_body, err := http.Get(url)

			if err != nil {
				fmt.Println("Error!", url)
			}
			
			body, err := ioutil.ReadAll(in_body.Body)
			if err != nil {
				fmt.Println("Error!", url)
			}

			//Converting the body to type string
			sb := string(body)

			arrayOfPageCode = append(arrayOfPageCode, sb)
			
			fmt.Println(len(arrayOfPageCode))	

		}(url)
	}
	
	wg.Wait()
	fmt.Println("All requests completed")
}