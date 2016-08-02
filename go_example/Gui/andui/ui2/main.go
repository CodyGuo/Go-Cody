package main

import (
    "encoding/json"
    "fmt"
    "github.com/andlabs/ui"
    "io/ioutil"
    "net/http"
    "net/url"
    "os"
    "reflect"
    "regexp"
)

//从有道获取的JSON的解析结构
type basic struct {
    Us_Phonetic string
    Phonetic    string
    Uk_Phonetic string
    Explains    []string
}
type web struct {
    Value []string
    Key   string
}
type Youdao struct {
    Translation []string
    Basic       basic
    Query       string
    ErrorCode   int32
    Web         []web
}

//从百度获取的输入反馈中的部分
type str struct {
    WordsToSearch string
}

var (
    w     ui.Window
    si    ui.Label
    al    ui.Label
    table ui.Table
    input ui.TextField
    proxy string
)

func main() {
    if len(os.Args) != 1 {
        t := getJSON(getWords())
        fmt.Println(simple(t))
    } else {
        go ui.Do(gui)
        err := ui.Go()
        if err != nil {
            panic(err)
        }
    }
}

//gui界面，用的 github.com/andlabs/ui 的ui库
func gui() {
    input = ui.NewTextField()
    input.OnChanged(func() {
        go feedback() // 主要是这里在不同的刷新(table)内容
    })
    button := ui.NewButton("查词")
    button.OnClicked(guiTranslate)
    inputBox := ui.NewHorizontalStack(input, button)
    inputBox.SetStretchy(0)
    tab := ui.NewTab()
    table = ui.NewTable(reflect.TypeOf(str{}))
    table.OnSelected(tableSelected)
    tab.Append("单词", table)
    si = ui.NewLabel("Simple")
    tab.Append("简约", si)
    al = ui.NewLabel("All")
    tab.Append("全部", al)
    stack := ui.NewVerticalStack(inputBox, tab)
    stack.SetStretchy(1)
    w = ui.NewWindow("Window", 280, 350, stack)
    w.OnClosing(func() bool {
        ui.Stop()
        return true
    })
    w.Show()
}

//取词并在屏幕上显示
func guiTranslate() {
    t := getJSON(input.Text())
    si.SetText(simple(t))
    al.SetText(all(t))
}

//输入反馈功能，用的百度的输入反馈接口
func feedback() {
    table.Lock()
    d := table.Data().(*[]str)
    *d = []str{}
    table.Unlock()
    in := input.Text()
    conn := Proxy()
    address := "HTTP://nssug.baidu.com/su?prod=recon_dict&wd=" + in
    resp, err := conn.Get(address)
    if err != nil {
        fmt.Println("Connect error")
        os.Exit(3)
    }
    t, _ := ioutil.ReadAll(resp.Body)
    t = regexp.MustCompile(`\[.+`).Find(t)
    tmp := regexp.MustCompile(`\b\w+?\b`).FindAllString(string(t), -1)
    s := make([]str, len(tmp))
    for i, key := range tmp {
        s[i].WordsToSearch = key
    }
    table.Lock()
    d = table.Data().(*[]str)
    *d = s
    table.Unlock()
}

//输入反馈功能，反馈项被选中时进行查房
func tableSelected() {
    i := table.Selected()
    data := table.Data().(*[]str)
    d := *data
    t := getJSON(d[i].WordsToSearch)
    si.SetText(simple(t))
    al.SetText(all(t))
}

//以简洁的方式输出翻译的string
func simple(y *Youdao) string {
    text := y.Query
    if len(y.Basic.Explains) == 0 {
        text = text + "\nNo translation, please check your word"
    } else {
        for _, t := range y.Translation {
            text += "\n" + t
        }
        for i, t := range y.Basic.Explains {
            if (i == 0) && (t == y.Translation[0]) {
            } else if i == 0 {
                text += "\n\n" + t
            } else {
                text += "\n" + t
            }
        }
    }
    return text
}

//输出全部有道翻译反馈的string
func all(y *Youdao) string {
    text := y.Query
    if len(y.Basic.Explains) == 0 {
        text = text + "\nNo translation, please check your word"
    } else {
        for _, t := range y.Translation {
            text += "\n" + t
        }
        for i, t := range y.Basic.Explains {
            if (i == 0) && (t == y.Translation[0]) {
            } else if i == 0 {
                text += "\n\n" + t
            } else {
                text += "\n" + t
            }
        }
        for _, t := range y.Web {
            text += "\n\n" + t.Key
            for i, tmp := range t.Value {
                if i == 0 {
                    text += "\n" + tmp
                } else {
                    text += ";" + tmp
                }
            }
        }
    }
    return text
}

//根据输入单词取得翻译结果并解析为一个struct
func getJSON(words string) *Youdao {
    conn := Proxy()
    address := "http://fanyi.youdao.com/openapi.do?keyfrom=" +
        "GoldenDictPlugin" + "&key=" + "1683580050" +
        "&type=data&doctype=json&version=1.1&q=" + words
    resp, err := conn.Get(address)
    if err != nil {
        fmt.Println("Connect error")
        os.Exit(1)
    }
    j, _ := ioutil.ReadAll(resp.Body)
    var data *Youdao
    err = json.Unmarshal(j, &data)
    if err != nil {
        fmt.Println(err)
        os.Exit(5)
    }
    return data
}

//命令行模式下获取作为参数传入的单词
func getWords() string {
    var words string
    l := len(os.Args)
    if l == 1 {
        os.Exit(0)
    } else {
        words = url.QueryEscape(os.Args[1])
    }
    for i := 2; i < l; i++ {
        words = words + "%20" + url.QueryEscape(os.Args[i])
    }
    return words
}

//返回一个可设置代理的 http Client，代理在全局变量 proxy 中设定
func Proxy() *http.Client {
    var conn http.Client
    if proxy != "" {
        proxyUrl, err := url.Parse(proxy)
        if err != nil {
            fmt.Println(err)
            os.Exit(2)
        }
        conn.Transport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
    }
    return &conn
}
