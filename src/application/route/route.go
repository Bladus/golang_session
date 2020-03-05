package route

import (
    "io/ioutil"
    "net/http"
    "log"
)

type Route struct {}

func (r *Route) GetRequest(url string) ([]byte, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    b := []byte(body)
    return b, nil
}

func check_error(err error) {
    if err != nil {
        log.Println(err)
    }
}