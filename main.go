package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx =context.Background()

func main() {

	newCli:=NewRedisCli()

	_,err:=newCli.Set(ctx,"salom","hello",1*time.Minute).Result()

	if err!= nil {
		
		log.Println("err on set data ", err)
		return
	}
	res,err:=newCli.Get(ctx,"salom").Result()

	if err!= nil {


		log.Println("err on get data ", err)
		return
	}

	fmt.Println("result",res)




}



func NewRedisCli()*redis.Client{

	option:=&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	}

	redisCli:=redis.NewClient(option)

	msg,err:=redisCli.Ping(context.Background()).Result()

	if err!= nil{

		log.Panicln("eror on NewRedisCli ",err)
		return nil
	}
	fmt.Println(msg)
	return redisCli
}