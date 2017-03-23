package main

func init() {
	register = new(Backup)
	register.loginUser = svn.loginUser
	register.loginPasswd = svn.loginPasswd

	register.svnUrl = "http://10.10.2.116:8088/svn/hupunac2.0/web/licenseManager"

	registerPath := MkBakDir(bakPath + "/注册服务器web代码")
	register.svnName = "licenseManager"
	register.pathName = registerPath + "/" + register.svnName

	register.fileName = "licensemanager.sql"
	register.cmd = cmdPublic + "rm -rf licensemanager.sql; mysqldump -R -h 10.10.2.251 -uroot -proot licensemanager >licensemanager.sql"
}
