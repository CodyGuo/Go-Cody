package main

import (
    "fmt"
    "log"
    "os"
    "time"
)

func isDirExists(path string) bool {
    file, err := os.Stat(path)

    if err != nil {
        return os.IsExist(err)
    } else {
        return file.IsDir()
    }
    panic("not reached.")
}

func init() {
    linuxSvn.Svn = "http://10.10.2.116:8088/svn/hupunac2.0/linux/aca"
    javaSvn.Svn = "http://10.10.2.116:8088/svn/hupunac2.0/web/hupunac"
    pcHelperSvn.Svn = "http://10.10.2.116:8088/svn/hupunac2.0/windows/hpidmnac"
    androidSvn.Svn = "http://10.10.2.116:8088/svn/hupunac2.0/Android/AndroidApp"
    nacUpgradeSvn.Svn = "http://10.10.2.116:8088/svn/hupunac2.0/linux/nac_upgrade"
    registerSvn.Svn = "http://10.10.2.116:8088/svn/hupunac2.0/web/licenseManager"

    autoTestingSvn.Svn = "http://10.10.2.116:8088/svn/hupunac2.0/test/AutomaTedtesting"
    webSiteSvn.Svn = "http://10.10.2.116:8088/svn/hupunac2.0/web/HupuWebsite"
    businessSvn.Svn = "http://10.10.2.116:8088/svn/hupunac2.0/web/hupuerp"

}

func main() {
    pathData := time.Now().Format(LAYOUT)
    bakPath = string(pathData) + string(BAK)
    linuxPath = "./" + bakPath + "/" + LINUX
    imanPath = "./" + bakPath + "/" + IMANWEB
    pcHelperPath = "./" + bakPath + "/" + PCHELPER
    andoridPath = "./" + bakPath + "/" + ANDROID
    upgradPath = "./" + bakPath + "/" + UPGRIDE
    registerPath = "./" + bakPath + "/" + REGISTER

    autoTestingPath = "./" + bakPath + "/" + AUTOTESTING
    webSitePath = "./" + bakPath + "/" + WEBSITE
    businessPath = "./" + bakPath + "/" + BUSINESS

    productPath = "./" + bakPath + "/" + PRODUCT
    bbsPath = "./" + bakPath + "/" + BBS

    os.Mkdir(bakPath, 0777)
    fmt.Printf("--------开始备份,请耐心等待完成,总计: %d----------\n", sumBak)
    // linux 代码备份
    LinuxBackup()

    // iman web 代码备份
    JavaBackup()

    // pc helper 代码备份
    PcHelperBackup()

    // android help 代码备份
    AndroidBackup()

    // upgrade 代码备份
    UpgradeBackup()

    // 注册服务器代码备份
    RegisterBackup()

    // // 自动化测试代码备份
    AutoTestingBackup()

    // 新官网代码备份
    WebSiteBackup()

    // 商机代码备份
    BusinessBackup()

    // 禅道备份
    ProcutBackup()

    // 论坛备份
    BbsBackup()

    fmt.Println("------------------已备份完成,请手动关闭窗口------------\n")
    var tmp string
    fmt.Scanln(&tmp)

}

func LinuxBackup() {
    os.MkdirAll(linuxPath, 0777)
    bakNum += 1
    fmt.Printf("--------正在进行第 %2d 个备份,总计: %d-------\n", bakNum, sumBak)
    log.Println("[INFO] 正在备份linux代码.")
    if isDirExists(linuxPath + "/aca") {
        linuxSvn.Update(linuxSvn.Svn, linuxPath+"/aca")
    } else {
        linuxSvn.Get(linuxSvn.Svn, linuxPath+"/aca")
    }
}

func JavaBackup() {
    os.MkdirAll(imanPath, 0777)
    bakNum += 1
    fmt.Printf("--------正在进行第 %2d 个备份,总计: %d-------\n", bakNum, sumBak)
    log.Println("[INFO] 正在备份iman web代码.")
    if isDirExists(imanPath + "/hupunac") {
        javaSvn.Update(javaSvn.Svn, imanPath+"/hupunac")
    } else {
        javaSvn.Get(javaSvn.Svn, imanPath+"/hupunac")
    }

    sqldir := imanPath + "/" + SQL + "/"
    os.MkdirAll(sqldir, 0777)
    sshCmd(imanWebSql, srcSqlPath, sqldir, true)
}

func PcHelperBackup() {
    os.MkdirAll(pcHelperPath, 0777)
    bakNum += 1
    fmt.Printf("--------正在进行第 %2d 个备份,总计: %d-------\n", bakNum, sumBak)
    log.Println("[INFO] 正在备份pc helper代码.")
    if isDirExists(pcHelperPath + "/hpidmnac") {
        pcHelperSvn.Update(pcHelperSvn.Svn, pcHelperPath+"/hpidmnac")
    } else {
        pcHelperSvn.Get(pcHelperSvn.Svn, pcHelperPath+"/hpidmnac")
    }
}

func AndroidBackup() {
    os.MkdirAll(andoridPath, 0777)
    bakNum += 1
    fmt.Printf("--------正在进行第 %2d 个备份,总计: %d-------\n", bakNum, sumBak)
    log.Println("[INFO] 正在备份android help代码.")
    if isDirExists(andoridPath + "/AndroidApp") {
        androidSvn.Update(androidSvn.Svn, andoridPath+"/AndroidApp")
    } else {
        androidSvn.Get(androidSvn.Svn, andoridPath+"/AndroidApp")
    }
}

func UpgradeBackup() {
    os.MkdirAll(upgradPath, 0777)
    bakNum += 1
    fmt.Printf("--------正在进行第 %2d 个备份,总计: %d-------\n", bakNum, sumBak)
    log.Println("[INFO] 正在备份upgrade代码.")
    if isDirExists(upgradPath + "/nac_upgrade") {
        nacUpgradeSvn.Update(nacUpgradeSvn.Svn, upgradPath+"/nac_upgrade")
    } else {
        nacUpgradeSvn.Get(nacUpgradeSvn.Svn, upgradPath+"/nac_upgrade")
    }
}

func RegisterBackup() {
    os.MkdirAll(registerPath, 0777)
    bakNum += 1
    fmt.Printf("--------正在进行第 %2d 个备份,总计: %d-------\n", bakNum, sumBak)
    log.Println("[INFO] 正在备份注册服务器代码.")
    if isDirExists(registerPath + "/licenseManager") {
        registerSvn.Update(registerSvn.Svn, registerPath+"/licenseManager")
    } else {
        registerSvn.Get(registerSvn.Svn, registerPath+"/licenseManager")
    }

    sqldir := registerPath + "/" + SQL + "/"
    os.Mkdir(sqldir, 0777)
    sshCmd(registerSql, srcSqlPath, sqldir, true)
}

func AutoTestingBackup() {
    os.MkdirAll(autoTestingPath, 0777)
    bakNum += 1
    fmt.Printf("--------正在进行第 %2d 个备份,总计: %d-------\n", bakNum, sumBak)
    log.Println("[INFO]正在备份自动化测试代码.")

    if isDirExists(autoTestingPath + "/AutomaTedtesting") {
        autoTestingSvn.Update(autoTestingSvn.Svn, autoTestingPath+"/AutomaTedtesting")
    } else {
        autoTestingSvn.Get(autoTestingSvn.Svn, autoTestingPath+"/AutomaTedtesting")
    }
}

func WebSiteBackup() {
    os.MkdirAll(webSitePath, 0777)
    bakNum += 1
    fmt.Printf("--------正在进行第 %2d 个备份,总计: %d-------\n", bakNum, sumBak)
    log.Println("[INFO] 正在备份新官网代码.")
    if isDirExists(webSitePath + "/HupuWebsite") {
        webSiteSvn.Update(webSiteSvn.Svn, webSitePath+"/HupuWebsite")
    } else {
        webSiteSvn.Get(webSiteSvn.Svn, webSitePath+"/HupuWebsite")
    }

    sqldir := webSitePath + "/" + SQL + "/"
    os.Mkdir(sqldir, 0777)
    sshCmd(webSiteSql, srcSqlPath, sqldir, true)
}

func BusinessBackup() {
    os.MkdirAll(businessPath, 0777)
    bakNum += 1
    fmt.Printf("--------正在进行第 %2d 个备份,总计: %d-------\n", bakNum, sumBak)
    log.Println("[INFO] 正在备份商机代码.")
    if isDirExists(businessPath + "/hupuerp") {
        businessSvn.Update(businessSvn.Svn, businessPath+"/hupuerp")
    } else {
        businessSvn.Get(businessSvn.Svn, businessPath+"/hupuerp")
    }

    sqldir := businessPath + "/" + SQL + "/"
    os.Mkdir(sqldir, 0777)
    sshCmd(businessSql, srcSqlPath, sqldir, true)
}

func ProcutBackup() {
    os.MkdirAll(productPath, 0777)
    bakNum += 1
    fmt.Printf("--------正在进行第 %2d 个备份,总计: %d-------\n", bakNum, sumBak)
    log.Println("[INFO] 正在备份禅道.")
    sqldir := productPath + "/" + SQL + "/"
    os.Mkdir(sqldir, 0777)
    sshCmd(productSql, srcSqlPath, sqldir, true)

}

func BbsBackup() {
    os.MkdirAll(bbsPath, 0777)
    bakNum += 1
    fmt.Printf("--------正在进行第 %2d 个备份,总计: %d-------\n", bakNum, sumBak)
    log.Println("[INFO] 正在备份论坛.")
    sqldir := bbsPath + "/" + SQL + "/"
    os.Mkdir(sqldir, 0777)
    sshCmd(bbsSql, srcSqlPath, sqldir, true)
}
