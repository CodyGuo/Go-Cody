// Copyright 2012-2014 The GoSNMP Authors. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in the
// LICENSE file.

// This program demonstrates BulkWalk.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/soniah/gosnmp"
)

var (
	writeFile string
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

	flag.StringVar(&writeFile, "w", "", "the oid nodes written to the file.")

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
	ok := false
	var file *os.File
	if writeFile != "" {
		file, _ = os.OpenFile(writeFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
		fmt.Fprintf(file, "%s = ", pdu.Name)
		ok = true
	} else {
		fmt.Printf("%s = ", pdu.Name)
	}
	defer file.Close()
	switch pdu.Type {
	case gosnmp.OctetString:
		b := pdu.Value.([]byte)
		if ok {
			fmt.Fprintf(file, "STRING: %s\n", string(b))
		} else {
			fmt.Printf("STRING: %s\n", string(b))
		}
	default:
		if ok {
			fmt.Fprintf(file, "TYPE %d: %d\n", pdu.Type, gosnmp.ToBigInt(pdu.Value))
		} else {
			fmt.Printf("TYPE %d: %d\n", pdu.Type, gosnmp.ToBigInt(pdu.Value))
		}
	}
	return nil
}
