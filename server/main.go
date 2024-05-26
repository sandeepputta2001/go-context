package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/sandeepputta2001/go-context/server/log"
)


func main(){

	flag.Parse()

	http.HandleFunc("/" , log.Decorate(handler))  // decorate is wrapper function , first decorate gets executed and then handler.	

	panic(http.ListenAndServe("127.0.0.1:5000",nil))

}

func handler(w http.ResponseWriter , r *http.Request){

	ctx := r.Context() 

	// context set in req is retrieved here.
	// context is hidden in every http request
	// by implementing context , server will knows if the client cancels the request in the middle of the process , which means
	// whenever the client cancels the request server will call  ctx.Done() which stops the server processing.


	log.Println(ctx ,"Handler started")

	defer log.Println(ctx , "Handler stopped")

	fmt.Printf("value for foo is %v" , ctx.Value("foo"))

	select {
	case <- time.After(5*time.Second):

	   fmt.Fprintln(w,"hello sandeep putta")
	case <- ctx.Done():
		
		err := ctx.Err()
		log.Println(ctx ,err.Error())
		http.Error(w , err.Error() , http.StatusInternalServerError)

	}

	
}