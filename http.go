package hpe3parprimeragoclient

import "github.com/go-resty/resty/v2"
import "crypto/tls"
import "fmt"

type AuthSuccess struct {
        /* variables */
}


func HttpPost(URI string, postBody string) (*resty.Response,error) {

   client := resty.New()

   client.SetTLSClientConfig(&tls.Config{ InsecureSkipVerify: true })
   fmt.Println("Body : ", postBody)

   resp, err := client.R().
                SetHeader("Content-type", "application/json").
                SetBody([]byte(postBody)).
                SetResult(&AuthSuccess{}).
                Post(URI)

   if err != nil {
      return nil, err
   }
   return  resp, nil

}

