package main

func init() {
	nacupgrade = new(Backup)
	nacupgrade.loginUser = svn.loginUser
	nacupgrade.loginPasswd = svn.loginPasswd

	nacupgrade.svnUrl = "http://10.10.2.116:8088/svn/hupunac2.0/linux/nac_upgrade"

	nacupgradePath := MkBakDir(bakPath + "/升级程序代码")
	nacupgrade.svnName = "nac_upgrade"
	nacupgrade.pathName = nacupgradePath + "/" + nacupgrade.svnName
}
