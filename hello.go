package client


import "net/http"
import "io/ioutil"

type Callback interface {
	OnSccess() string
	OnError() string
}

func Get(url string, callback Callback)  string {

    resp, _ := http.Get(url)
    defer resp.Body.Close()

    byteArray, _ := ioutil.ReadAll(resp.Body)
    return string(byteArray)
}

