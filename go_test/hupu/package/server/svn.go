package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	LOCALPATH = "D:/iMan-SVN/helper"
	SVNURL    = "http://10.10.2.116:8088/svn/hupunac2.0/windows"
)

var (
	svnUser   string
	svnPasswd string
	svnLib    string
	branche   bool

	helperLocalPath string
)

func init() {
	flag.StringVar(&svnUser, "svnUser", "cody.guo", "set svn user.")
	flag.StringVar(&svnPasswd, "svnPasswd", "iman2015", "set svn passwd.")
	flag.StringVar(&svnLib, "svnLib", "hpidmnac", "set svn URL.")
	flag.BoolVar(&branche, "branche", false, "set branche true or false.")

	if _, err := os.Stat(LOCALPATH); err != nil {
		os.MkdirAll(LOCALPATH, 0777)
	}
}

type SVN struct {
	user   string
	passwd string
	libURL string
}

func NewSVN() *SVN {
	svn := new(SVN)
	svn.user = svnUser
	svn.passwd = svnPasswd
	if branche {
		svn.libURL = fmt.Sprintf("%s/branches/%s", SVNURL, svnLib)
	} else {
		svn.libURL = fmt.Sprintf("%s/%s", SVNURL, svnLib)
	}

	return svn
}

func (s *SVN) Get() error {
	fmt.Printf("%sIs getting SVN code, branch-->[ %s ].\n", INFO, s.libURL)
	helperLocalPath = fmt.Sprintf("%s/%s", LOCALPATH, svnLib)
	if _, err := os.Stat(helperLocalPath); err != nil {
		return s.get()
	}

	return s.update()
}

func (s *SVN) get() error {
	cmd := "svn checkout " + s.libURL + " --username=" + s.user +
		" --password=" + s.passwd + " --no-auth-cache " + helperLocalPath

	return runCmd(cmd)
}

func (s *SVN) update() error {
	cmd := "svn update --username=" + s.user +
		" --password=" + s.passwd + " --no-auth-cache " + helperLocalPath

	if err := runCmd(cmd); err != nil {
		if err := os.RemoveAll(helperLocalPath); err != nil {
			return err
		}
		return s.get()
	}
	return nil
}
