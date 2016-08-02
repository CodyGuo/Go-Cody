package main

func init() {
	erpApp = new(Backup)
	erpApp.loginUser = svn.loginUser
	erpApp.loginPasswd = svn.loginPasswd

	erpApp.svnUrl = "http://10.10.2.116:8088/svn/hupunac2.0/Android/erpApp"

	erpAppPath := MkBakDir(bakPath + "/商机系统App代码")
	erpApp.svnName = "erpApp"
	erpApp.pathName = erpAppPath + "/" + erpApp.svnName
}
