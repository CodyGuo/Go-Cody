package main

func init() {
	linuxBak = new(Backup)
	linuxBak.loginUser = svn.loginUser
	linuxBak.loginPasswd = svn.loginPasswd

	linuxBak.svnUrl = "http://10.10.2.116:8088/svn/hupunac2.0/linux/aca"

	linuxPath := MkBakDir(bakPath + "/linux代码")
	linuxBak.svnName = "aca"
	linuxBak.pathName = linuxPath + "/" + linuxBak.svnName
}
