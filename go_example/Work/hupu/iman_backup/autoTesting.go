package main

func init() {
	autoTesting = new(Backup)
	autoTesting.loginUser = svn.loginUser
	autoTesting.loginPasswd = svn.loginPasswd

	autoTesting.svnUrl = "http://10.10.2.116:8088/svn/hupunac2.0/test/AutomaTedtesting"

	autoTestingPath := MkBakDir(bakPath + "/自动化测试代码")
	autoTesting.svnName = "AutomaTedtesting"
	autoTesting.pathName = autoTestingPath + "/" + autoTesting.svnName
}
