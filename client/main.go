package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main(){

	ctx := context.Background()

	ctx,cancel := context.WithTimeout(ctx,30*time.Second)

	defer cancel()

	ctx = context.WithValue(ctx , "foo" , "bar") 

	req, err := http.NewRequest(http.MethodGet, "http://localhost:5000" , nil)

	if err != nil{
		log.Fatal(err)

	}

	req = req.WithContext(ctx) // used to create  context .  context is set in req object

	res , err := http.DefaultClient.Do(req)
	if err !=nil{	
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}

	io.Copy(os.Stdout , res.Body)
}