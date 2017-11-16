package main

import (
  "fmt"
  "flag"
  "net/http"
  "os"
)

func main() {
	 endpoint := flag.String("endpoint", "", "The http endpoint to be monitored.")
	
	flag.Parse()
	
	fmt.Println("Endpoint to be monitored:", *endpoint)
	
        for {
	    resp, err := http.Get(*endpoint)

            if (err == nil) && (resp.StatusCode == 200) {
              os.Exit(0) 
            } else if (err == nil) {
		fmt.Println(resp.StatusCode)
	    } else {
		fmt.Println(err)
            }
        } 
}
