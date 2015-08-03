package main

import (
    "fmt"

    "github.com/Unknwon/goconfig"
)

/*
# ini cfg configuration files such as read.
[Cody]
name : cody.guo
# sex注解top
sex : boy

# sex注解
[test]
测试 = 冒号
test = ok
No = no

[work]
work = IT
*/

func main() {
    read, err := goconfig.LoadConfigFile("read.ini")
    if err != nil {
        fmt.Println(err)
    }

    /*ini cfg配置文件读取*/
    fmt.Println(read.GetSectionList())              // 获取选项名称
    fmt.Println(read.GetKeyList("Cody"))            // 获取Cody选项下的key
    fmt.Println(read.GetSection("Cody"))            // 以map形式获取Cody选项下的key和value的值
    fmt.Println(read.GetSectionComments("work"))    // 获取选项work上的以“#”和“;”开头的注解
    fmt.Println(read.GetKeyComments("Cody", "sex")) // 获取选项Cody的sex键上的注解
    fmt.Println(read.GetValue("Cody", "name"))      // 获取选项Cody的name键的值

    /**/
    read.SetValue("Cody", "name", "cody.guo") // 设置选项Cody的name键的值为cody.guo
    goconfig.SaveConfigFile(read, "read.ini") // 保存

    read.DeleteKey("work", "work")                 // 删除选项work中的work键
    goconfig.SaveConfigFile(read, "readWrite.ini") // 另存为readWrite.ini

    read.DeleteSection("work")                // 删除选项work
    goconfig.SaveConfigFile(read, "read.ini") // 保存配置文件,会把 " : "的修改为" = "
}
