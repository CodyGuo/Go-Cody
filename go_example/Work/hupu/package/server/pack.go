package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
)

const (
	HELPEREXE = "hpidmNacSetup.exe"
)

type Package struct {
	fullName string
}

func NewPack() (*Package, error) {
	if err := NewSVN().Get(); err != nil {
		return nil, err
	}

	if err := NewBuildConfig().SetValue(); err != nil {
		return nil, err
	}

	os.RemoveAll(helperLocalPath + "\\out\\" + HELPEREXE)

	return new(Package), nil
}

func (p *Package) Pack() error {
	fmt.Printf("%sPack helper, please wait a moment...\n", INFO)
	var err error
	cmd := fmt.Sprintf("%s\\%s %s", helperLocalPath, "packinstall.exe", `/s`)
	for range rand.Perm(2) {
		err = runCmd(cmd)
		if err == nil {
			break
		}
	}

	return err
}

func (p *Package) Upload() error {
	fmt.Printf("%sAre uploading helper, please wait a moment...\n", INFO)
	if err := p.getName(); err != nil {
		return err
	}

	ftp := NewFTP()
	if err := ftp.Upload(p.fullName, HELPEREXE); err != nil {
		return err
	}

	return nil
}

func (p *Package) getName() error {
	packPath := fmt.Sprintf("%s\\%s", helperLocalPath, "out")
	files, err := ioutil.ReadDir(packPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if file.Name() == HELPEREXE {
			p.fullName = fmt.Sprintf("%s\\%s", packPath, file.Name())
			return nil
		}
	}

	return fmt.Errorf("%s exe is not packaged.", ERROR)
}
