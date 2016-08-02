package main

func init() {
	product = new(Backup)
	productPath := MkBakDir(bakPath + "/禅道备份")
	product.pathName = productPath

	product.fileName = "hupu.sql"
	product.cmd = cmdPublic + "rm -rf hupu.sql; mysqldump -R -h 10.10.2.222 -uroot -p123456 hupu >hupu.sql"
}
