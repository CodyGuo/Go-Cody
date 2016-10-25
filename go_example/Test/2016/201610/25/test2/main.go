package main

import (
	"fmt"
	"github.com/egonelbre/telnet"
)

func main() {
	fmt.Printf("Server started on :8000\n")
	telnet.ListenAndServe(":23", handle)
}

func handle(conn *telnet.Conn) {
	var login []string
	login = append(login, "Username:", "Password:")
	var loginCmd map[string]string = map[string]string{
		"Username:": "h3c",
		"Password:": "h3c",
	}

	var cmd map[string]string = map[string]string{
		"dis arp": `Type: S-Static   D-Dynamic
IP Address       MAC Address     VLAN ID  Port Name / AL ID      Aging     Type
10.10.1.138      2cd0-5ae3-8614  1        Ethernet1/0/18         11        D
10.10.1.195      3c97-0efd-b361  1        Ethernet1/0/35         11        D
10.10.1.180      4437-e68a-d304  1        Ethernet1/0/2          13        D
10.10.1.102      507b-9d6d-e0c1  1        Ethernet1/0/18         14        D
10.10.1.246      0025-115a-b7f8  1        Ethernet1/0/3          16        D
10.10.1.223      b051-8e03-a103  1        Ethernet1/0/13         17        D
10.10.1.175      f0de-f1f5-e388  1        Ethernet1/0/18         18        D
10.10.1.176      74e5-0bf4-16bc  1        Ethernet1/0/18         18        D
10.10.1.119      8c89-a5ed-fdd7  1        Ethernet1/0/18         19        D
10.10.1.200      acb5-7d8e-d9c3  1        Ethernet1/0/18         19        D
10.10.1.170      acb5-7d88-efc8  1        Ethernet1/0/18         19        D
10.10.1.189      7423-44bd-e47e  1        Ethernet1/0/18         19        D
10.10.1.187      0025-115a-b939  1        Ethernet1/0/18         19        D
10.10.1.137      4016-9f7a-3449  1        Ethernet1/0/18         19        D
10.10.1.147      0025-1159-a73e  1        Ethernet1/0/18         19        D
10.10.1.133      8c89-a5ed-f272  1        Ethernet1/0/18         19        D
10.10.1.149      68f7-2888-3c9e  1        Ethernet1/0/39         20        D
10.10.1.110      10c3-7bba-07ca  1        Ethernet1/0/18         20        D
10.10.1.105      00e0-4c36-11fb  1        Ethernet1/0/18         20        D
10.10.1.123      001e-64c5-ce48  1        Ethernet1/0/18         20        D
10.10.1.148      2c33-7a5f-e6e1  1        Ethernet1/0/18         20        D
10.10.1.126      4437-e68f-31f5  1        Ethernet1/0/18         20        D
10.10.1.136      3c97-0efd-bd7c  1        Ethernet1/0/39         20        D
10.10.1.127      6c0b-843e-f9eb  1        Ethernet1/0/2          20        D
10.10.1.254      000f-e2d3-9d75  1        Ethernet1/0/2          20        D
10.10.1.1        0090-7f84-5093  1        Ethernet1/0/1          20        D
10.10.1.122      0021-cc6d-ccb1  1        Ethernet1/0/18         20        D
10.10.1.248      000c-2923-0b67  1        Ethernet1/0/2          20        D
10.10.1.129      68f7-2888-3380  1        Ethernet1/0/39         20        D
10.10.1.117      68f7-2887-5a85  1        Ethernet1/0/18         20        D
10.10.1.111      68f7-2888-42b8  1        Ethernet1/0/18         20        D

---   31 entries found   ---
`,
	}

	for _, l := range login {
		conn.Print(l)
		log_in := <-conn.Lines
		if log_in != loginCmd[l] {
			conn.Print("user password error\n")
			conn.Terminate()
			return
		}
	}

	conn.Print("<H3C>")
	for {
		line := <-conn.Lines
		fmt.Printf("%s\n", line)
		for {
			if _, ok := cmd[line]; ok {
				conn.Printf("%s\n", cmd[line])
				if line == "quit" {
					conn.Terminate()
				}
			}
			conn.Print("<H3C>")
			line = <-conn.Lines
		}

	}
}
