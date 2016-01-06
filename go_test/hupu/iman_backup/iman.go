package main

func init() {
	iManBak = new(Backup)
	iManBak.loginUser = svn.loginUser
	iManBak.loginPasswd = svn.loginPasswd

	iManBak.svnUrl = "http://10.10.2.116:8088/svn/hupunac2.0/web/hupunac"

	javaPath := MkBakDir(bakPath + "/iMan代码")
	iManBak.svnName = "hupunac"
	iManBak.pathName = javaPath + "/" + iManBak.svnName

	iManBak.fileName = "hupunac.sql"
	iManBak.cmd = cmdPublic + "rm -rf hupunac.sql; mysqldump -R -h 10.10.2.230 -uroot -proot hupunac >hupunac.sql"
}
