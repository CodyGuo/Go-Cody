package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

import (
	"github.com/axgle/mahonia"
	"github.com/soniah/gosnmp"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		fmt.Printf("   %s [-c=<community>] host [oid]\n", filepath.Base(os.Args[0]))
		fmt.Printf("     host      - the host to walk/scan\n")
		fmt.Printf("     oid       - the MIB/Oid defining a subtree of values\n\n")
		flag.PrintDefaults()
	}

	var community string
	flag.StringVar(&community, "c", "public", "the community string for device")

	flag.Parse()

	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	target := flag.Args()[0]
	var oid string
	if len(flag.Args()) > 1 {
		oid = flag.Args()[1]
	}

	gosnmp.Default.Target = target
	gosnmp.Default.Community = community
	gosnmp.Default.Timeout = time.Duration(10 * time.Second) // Timeout better suited to walking
	err := gosnmp.Default.Connect()
	if err != nil {
		fmt.Printf("Connect err: %v\n", err)
		os.Exit(1)
	}
	defer gosnmp.Default.Conn.Close()

	err = gosnmp.Default.BulkWalk(oid, printValue)
	if err != nil {
		fmt.Printf("Walk Error: %v\n", err)
		os.Exit(1)
	}
}

func printValue(pdu gosnmp.SnmpPDU) error {
	fmt.Printf("%s = ", pdu.Name)
	dec := mahonia.NewDecoder("gbk")
	switch pdu.Type {
	case gosnmp.OctetString:
		b := pdu.Value.([]byte)
		_, date, _ := dec.Translate(b, true)
		fmt.Printf("STRING: %s\n", string(date))
	case gosnmp.TimeTicks:
		b := gosnmp.ToBigInt(pdu.Value)
		t, _ := strconv.Atoi(fmt.Sprintf("%d", b))
		str_time := time.Unix(int64(t), 0).Format("Mon Jan 2 15:04:05.000 2006 +0800")
		fmt.Printf("TimeTicks: %d %s\n", b, str_time)
	case gosnmp.IPAddress:
		b := pdu.Value.(string)
		fmt.Printf("IpAddress: %s\n", string(b))
	default:
		fmt.Printf("TYPE %d: %d\n", pdu.Type, gosnmp.ToBigInt(pdu.Value))
	}
	return nil
}
