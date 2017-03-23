package main

func init() {
	webSite = new(Backup)
	webSite.loginUser = svn.loginUser
	webSite.loginPasswd = svn.loginPasswd

	webSite.svnUrl = "http://10.10.2.116:8088/svn/hupunac2.0/web/HupuWebsite"

	webSitePath := MkBakDir(bakPath + "/新官网代码")
	webSite.svnName = "HupuWebsite"
	webSite.pathName = webSitePath + "/" + webSite.svnName

	webSite.fileName = "hupuwebsite.sql"
	webSite.cmd = cmdPublic + "rm -rf hupuwebsite.sql; mysqldump -R -h 10.10.2.230 -uroot -proot hupuwebsite >hupuwebsite.sql"
}
