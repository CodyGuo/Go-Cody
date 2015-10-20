package main

import (
    "fmt"
    "github.com/secsy/goftp"
)

func main() {

    config := goftp.Config{
        User:             "nac_ftp",
        Password:         "qaz!@#",
        ActiveListenAddr: "10.10.2.140",
        ActiveTransfers:  true,
    }
    ftpConn, err := goftp.DialConfig(config, "10.10.3.227")
    fmt.Println(err)
    getwd, _ := ftpConn.Getwd()

    ftpConn.Mkdir("codyguo")
    fmt.Println(getwd)
}
