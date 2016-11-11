package main

func init() {
	erp = new(Backup)
	erp.loginUser = svn.loginUser
	erp.loginPasswd = svn.loginPasswd

	erp.svnUrl = "http://10.10.2.116:8088/svn/hupunac2.0/web/hupuerp"

	erpPath := MkBakDir(bakPath + "/商机系统Web代码")
	erp.svnName = "hupuerp"
	erp.pathName = erpPath + "/" + erp.svnName

	erp.fileName = "hupuerp.sql"
	erp.cmd = cmdPublic + "rm -rf hupuerp.sql; mysqldump -R -h 10.10.2.251 -uroot -proot --force hupuerp >hupuerp.sql"
}
