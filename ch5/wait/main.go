package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail.

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s);retrying...", err)
		//time.Sleep(time.Second << unit(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func add1(r rune) rune {
	return r + 1
}

func main() {

	fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
	fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
	fmt.Println(strings.Map(add1, "Admix"))    // "Benjy

	fmt.Println(func(r rune) rune {
		return r + 1
	}, "HAL-9000")
}
