package main

var (
    sumBak     = 11
    bakNum     = 0
    backName   string
    srcProPath = "/var/www/"
    srcSqlPath = "/root/cody_develop/sql/sqlbackup/"

    bakPath, linuxPath, imanPath, pcHelperPath, andoridPath, upgradPath, registerPath string

    autoTestingPath, webSitePath, businessPath, productPath, bbsPath string
)
var (
    autoTestingSvn, webSiteSvn, businessSvn SvnUrl

    linuxSvn, javaSvn, androidSvn, pcHelperSvn, nacUpgradeSvn, registerSvn SvnUrl
)

var (
    cmdPublic = "mkdir -p /root/cody_develop/sql/sqlbackup; cd /root/cody_develop/sql/sqlbackup;"

    imanWebSql  = cmdPublic + "rm -rf hupunac.sql; mysqldump -R -h 10.10.2.230 -uroot -proot hupunac >hupunac.sql"
    webSiteSql  = cmdPublic + "rm -rf hupuwebsite.sql; mysqldump -R -h 10.10.2.230 -uroot -proot hupuwebsite >hupuwebsite.sql"
    businessSql = cmdPublic + "rm -rf hupuerp.sql; mysqldump -R -h 10.10.2.230 -uroot -proot hupuerp >hupuerp.sql"

    registerSql = cmdPublic + "rm -rf licensemanager.sql; mysqldump -R -h 10.10.2.251 -uroot -proot licensemanager >licensemanager.sql"

    productFile = "cd /var/www/; rm -rf zentaopms.tar.gz; tar zcvf zentaopms.tar.gz --exclude=zentaopms/backup --exclude=zentaopms/tmp zentaopms/"
    productSql  = cmdPublic + "rm -rf hupu.sql; mysqldump -R -h 10.10.2.222 -uroot -p123456 hupu >hupu.sql"
    bbsSql      = cmdPublic + "rm -rf hupubbs.sql; mysqldump -R -h 10.10.2.222 -uroot -p123456 hupubbs >hupubbs.sql"
)

const (
    LAYOUT = "2006-01-02"

    BAK         = "研发部备份"
    BBS         = "论坛"
    SQL         = "数据库"
    LINUX       = "linux代码"
    IMANWEB     = "web代码"
    PRODUCT     = "禅道"
    ANDROID     = "android代码"
    UPGRIDE     = "升级程序代码"
    WEBSITE     = "新官网代码"
    PCHELPER    = "小助手代码"
    REGISTER    = "注册服务器web代码"
    AUTOTESTING = "自动化测试代码"
    BUSINESS    = "商机系统代码"
)
