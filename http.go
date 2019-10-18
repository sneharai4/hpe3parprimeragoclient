package hpe3parprimeragoclient

import "github.com/go-resty/resty/v2"
import "crypto/tls"
import "fmt"

type AuthSuccess struct {
   Key string `json:"key"`
}

var session_key AuthSuccess

func HttpPost(URI string, postBody string) (*resty.Response, error) {

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

   fmt.Println("Response body is", resp.Body())
   json.Unmarshal(resp.Body(), &session_key)
   fmt.Println("Printing only session key out: ", session_key.Key)
   fmt.Println("Printing only session key struct out: ", session_key)

   return  resp, nil
