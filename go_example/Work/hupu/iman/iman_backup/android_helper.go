package main

func init() {
	androidHelper = new(Backup)
	androidHelper.loginUser = svn.loginUser
	androidHelper.loginPasswd = svn.loginPasswd

	androidHelper.svnUrl = "http://10.10.2.116:8088/svn/hupunac2.0/Android/AndroidApp"

	andoridPath := MkBakDir(bakPath + "/android助手")
	androidHelper.svnName = "AndroidApp"
	androidHelper.pathName = andoridPath + "/" + androidHelper.svnName
}
