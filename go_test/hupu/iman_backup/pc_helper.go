package main

func init() {
	pcHelper = new(Backup)
	pcHelper.loginUser = svn.loginUser
	pcHelper.loginPasswd = svn.loginPasswd

	pcHelper.svnUrl = "http://10.10.2.116:8088/svn/hupunac2.0/windows/hpidmnac"

	pcHelperPath := MkBakDir(bakPath + "/PC助手代码")
	pcHelper.svnName = "hpidmnac"
	pcHelper.pathName = pcHelperPath + "/" + pcHelper.svnName
}
