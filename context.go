// this is a seperate file about context which is different from client and server files mentioned.
package main

import (
	"context"
	"fmt"
	"log"
	"time"
)


func main(){

ctx := context.Background()

ctx , cancel:= context.WithTimeout(ctx, time.Second)

 cancel()

mySleepAndTalk(ctx , 5*time.Second, "hello sandeep putta")

}


func mySleepAndTalk(ctx context.Context , d time.Duration, msg string){

	select{ 
	case <- time.After(d):
		        fmt.Println(msg)
	case <- ctx.Done():
		log.Print( "sandeep putta context cancelled", ctx.Err())
	}

}
