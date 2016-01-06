package main

import (
    // "bufio"
    "github.com/dynport/gossh"
    "github.com/pkg/sftp"
    "log"
    // "strings"
)

// returns a function of type gossh.Writer func(...interface{})
// MakeLogger just adds a prefix (DEBUG, INFO, ERROR)
func MakeLogger(prefix string) gossh.Writer {
    return func(args ...interface{}) {
        log.Println((append([]interface{}{prefix}, args...))...)
    }
}

func main() {
    client := gossh.New("10.10.2.222", "root")
    // my default agent authentication is used. use
    client.SetPassword("hpiman")
    // for password authentication
    client.DebugWriter = MakeLogger("DEBUG")
    client.InfoWriter = MakeLogger("INFO ")
    client.ErrorWriter = MakeLogger("ERROR")
    client.Execute("mysqldump -R -uroot -p123456 hupunac >/root/cody_develop/sql/hupunac.sql")

    conn, _ := client.Connection()
    c, err := sftp.NewClient(conn, sftp.MaxPacket(5e9))
    c.OpenFile("/root/cody_develop/sql", "hupunac.sql")

    defer client.Close()

}
