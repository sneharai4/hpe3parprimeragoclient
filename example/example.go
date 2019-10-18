package main

import "fmt"
import "github.com/sneharai4/hpe3parprimeragoclient"

func main(){

  response,err := hpe3parprimeragoclient.HttpPost("https://15.212.192.252:8080/api/v1/credentials","{\"user\":\"3paradm\",\"password\":\"3pardata\"}")
  fmt.Println(" Response : ", response)
  fmt.Println(" Error : ", err)

}

