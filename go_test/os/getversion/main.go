// golang 实现读取exe dll apk 版本号
package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "path/filepath"
)

import (
    "github.com/lunny/axmlParser"
)

var (
    file fileInfo
)

const (
    MZ       = "MZ"
    PE       = "PE"
    RSRC     = ".rsrc"
    TYPET    = 16
    PEOFFSET = 64
    MACHINE  = 332
    DEFAULT  = `C:\Windows\System32\cmd.exe`
)

type fileInfo struct {
    FilePath string
    Version  string
    Debug    bool
}

func (f *fileInfo) checkError(err error) {
    if err != nil {
        log.Fatalln(err)
    }
}

func (f *fileInfo) GetExeVersion() (err error) {
    file, err := os.Open(f.FilePath)
    f.checkError(err)

    // 第一次读取64 byte
    buffer := make([]byte, 64)
    _, err = file.Read(buffer)
    f.checkError(err)
    defer file.Close()

    str := string(buffer[0]) + string(buffer[1])
    if str != MZ {
        log.Fatalln("读取exe错误,找不到 MZ.", f.FilePath)
    }

    peOffset := f.unpack([]byte{buffer[60], buffer[61], buffer[62], buffer[63]})
    if peOffset < PEOFFSET {
        log.Fatalln("peOffset 读取错误.", f.FilePath)
    }

    // 读取从文件开头移位到 peOffset，第二次读取 24 byte
    _, err = file.Seek(int64(peOffset), 0)
    buffer = make([]byte, 24)
    _, err = file.Read(buffer)
    f.checkError(err)

    str = string(buffer[0]) + string(buffer[1])
    if str != PE {
        log.Fatalln("读取exe错误,找不到 PE.", f.FilePath)
    }

    machine := f.unpack([]byte{buffer[4], buffer[5]})
    if machine != MACHINE {
        log.Fatalln("machine 读取错误.", f.FilePath)
    }

    noSections := f.unpack([]byte{buffer[6], buffer[7]})
    optHdrSize := f.unpack([]byte{buffer[20], buffer[21]})

    // 读取从当前位置移位到 optHdrSize，第三次读取 40 byte
    file.Seek(int64(optHdrSize), 1)
    resFound := false
    for i := 0; i < int(noSections); i++ {
        buffer = make([]byte, 40)
        file.Read(buffer)
        str = string(buffer[0]) + string(buffer[1]) + string(buffer[2]) + string(buffer[3]) + string(buffer[4])
        if str == RSRC {
            resFound = true
            break
        }
    }
    if !resFound {
        log.Fatalln("读取exe错误,找不到 .rsrc.", f.FilePath)
    }

    infoVirt := f.unpack([]byte{buffer[12], buffer[13], buffer[14], buffer[15]})
    infoSize := f.unpack([]byte{buffer[16], buffer[17], buffer[18], buffer[19]})
    infoOff := f.unpack([]byte{buffer[20], buffer[21], buffer[22], buffer[23]})

    // 读取从文件开头位置移位到 infoOff，第三次读取 infoSize byte
    file.Seek(int64(infoOff), 0)
    buffer = make([]byte, infoSize)
    _, err = file.Read(buffer)
    f.checkError(err)

    nameEntries := f.unpack([]byte{buffer[12], buffer[13]})
    idEntries := f.unpack([]byte{buffer[14], buffer[15]})

    var infoFound bool
    var subOff, i int64
    for i = 0; i < (nameEntries + idEntries); i++ {
        typeT := f.unpack([]byte{buffer[i*8+16], buffer[i*8+17], buffer[i*8+18], buffer[i*8+19]})
        if typeT == TYPET {
            infoFound = true
            subOff = int64(f.unpack([]byte{buffer[i*8+20], buffer[i*8+21], buffer[i*8+22], buffer[i*8+23]}))
            break
        }
    }
    if !infoFound {
        log.Fatalln("读取exe错误,找不到 typeT == 16.", f.FilePath)
    }

    subOff = subOff & 0x7fffffff
    infoOff = f.unpack([]byte{buffer[subOff+20], buffer[subOff+21], buffer[subOff+22], buffer[subOff+23]}) //offset of first FILEINFO
    infoOff = infoOff & 0x7fffffff
    infoOff = f.unpack([]byte{buffer[infoOff+20], buffer[infoOff+21], buffer[infoOff+22], buffer[infoOff+23]}) //offset to data
    dataOff := f.unpack([]byte{buffer[infoOff], buffer[infoOff+1], buffer[infoOff+2], buffer[infoOff+3]})
    dataOff = dataOff - infoVirt

    version1 := f.unpack([]byte{buffer[dataOff+48], buffer[dataOff+48+1]})
    version2 := f.unpack([]byte{buffer[dataOff+48+2], buffer[dataOff+48+3]})
    version3 := f.unpack([]byte{buffer[dataOff+48+4], buffer[dataOff+48+5]})
    version4 := f.unpack([]byte{buffer[dataOff+48+6], buffer[dataOff+48+7]})

    version := fmt.Sprintf("%d.%d.%d.%d", version2, version1, version4, version3)
    f.Version = version

    return nil
}

func (f *fileInfo) unpack(b []byte) (num int64) {
    for i := 0; i < len(b); i++ {
        num = 256*num + int64((b[len(b)-1-i] & 0xff))
    }
    return
}

func (f *fileInfo) GetApkVersion() (err error) {
    listener := new(axmlParser.AppNameListener)
    _, err = axmlParser.ParseApk(f.FilePath, listener)
    f.checkError(err)

    f.Version = listener.VersionName
    return nil
}

func init() {
    flag.StringVar(&file.FilePath, "path", DEFAULT, "Get exe or dll or apk version information.")
    flag.BoolVar(&file.Debug, "d", false, "if true print exe or dll file name.")
}
func main() {
    flag.Parse()

    suffix := filepath.Ext(file.FilePath)
    switch suffix {
    case ".exe", ".dll":
        file.GetExeVersion()
    case ".apk":
        file.GetApkVersion()
    default:
        log.Fatalln("仅能获取exe、dll、apk版本号,请重新输入程序路径.", file.FilePath)
    }

    switch {
    case file.Debug:
        fmt.Printf("%s 的版本号为: ", file.FilePath)
    case file.FilePath == DEFAULT:
        flag.PrintDefaults()
        fmt.Printf("%s 的版本号为: ", file.FilePath)
    }

    fmt.Printf("%s", file.Version)
}
