package main

import (
    // "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    // "os"
)

type Kjdata struct {
    Name  string
    Order string
}

func main() {
    url := "http://www.6617.com/xml/kjdata.json"
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
    }

    rebots, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }
    resp.Body.Close()

    var result []Kjdata
    json.Unmarshal(rebots, &result)

    fmt.Println(result)
    // fmt.Println(string(rebots))

    // var kjdata []Kjdata
    // err = json.Unmarshal(rebots[:len(rebots)], &kjdata)
    // if err != nil {
    //     fmt.Println("json: " + err.Error())
    // }

    // fmt.Println(kjdata)

}
