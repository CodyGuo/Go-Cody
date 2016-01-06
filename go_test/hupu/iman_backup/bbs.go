package main

func init() {
	bbs = new(Backup)
	bbsPath := bakPath + "/论坛BBS"
	bbs.pathName = bbsPath

	bbs.fileName = "hupubbs.sql"
	bbs.cmd = cmdPublic + "rm -rf hupubbs.sql; mysqldump -R -h 10.10.2.222 -uroot -p123456 hupubbs >hupubbs.sql"
}
