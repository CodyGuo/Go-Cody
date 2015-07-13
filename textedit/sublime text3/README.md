# sublime text3
轻量级的代码编辑器

# 目录
## 1. 插件
常用插件安装：
http://www.cnsecer.com/3976.html
http://www.daqianduan.com/4820.html
http://www.cnblogs.com/Rising/p/3741116.html
http://www.cnblogs.com/dudumao/p/4054086.html

----------------------------------------------------------
名称 | 说明 | 配置详解
------------- | ------------- | -------------
`SublimeGit` |  用于git提交代码到github | Ctrl + Shift + P --> git quick add，git quick commit，git push current branch`Anaconda` | 代码自动提示和补全 |
`ConvertToUTF8` | 支持UTF8编码的代码 |
`SublimeLinter` | 代码高亮显示 |
`CSS Format` | 修改user配置后，可以保存CSS时自动格式化 |
`Tag` | HTML格式化插件 |
`Alignment` | 主要用于代码对齐，最新版据说已经集成了这个插件 | http://my.oschina.net/shede333/blog/170536
`BracketHighlighter` | 插件高亮设置方法 | http://www.tuicool.com/articles/reQJBj
`Ctags` | 函数跳转 | https://www.zybuluo.com/lanxinyuchs/note/33551，http://my.oschina.net/miaowang/blog/197060
`DocBlockr` | 代码自动注释生成 | http://blog.asdasd.cn/?p=740
`Emmet` | (前身为 Zen Coding) 是一个能大幅度提高前端开发效率的一个工具 | http://www.w3cplus.com/tools/emmet-cheat-sheet.html
`SublimeGDB` | Go使用GDB断点调试 | http://www.cnblogs.com/yourihua/archive/2012/11/12/2766763.html
`Side Bar` | 美化左侧导航栏 | http://blog.csdn.net/zm2714/article/details/8064259
`GoSublime` | Go | http://my.oschina.net/mrcuix/blog/153249
`ColorPicker` | 调色板 |


## 2. 常见问题
### 2.1 Tag安装后不能使用Ctrl + Alt + F 进行格式化
* 解决方法来自: http://www.java123.net/v/934921.html
  解决方法：在Github上下载压缩包版手动安装https://github.com/SublimeText/Tag 将解压的Tag-st3放在   sublime text目录\Data\Packages下，重启
快捷键能用。

### 2.2 Sublime text3 插件ColorPicker(调色板)不能使用快捷键
* http://www.java123.net/v/946446.html

### 2.3 保存Go文件时格式化tab为spaces
* 打开Preferences-->Package Settings-->GoSublime-->Settings Default-->修改"fmt_tab_indent": false，为false。

### 2.4 GoSublime --> User Settings -->:
* {
*    "env": {
*     "GOROOT":"D:\\go", //go的安装路径
*     "GOBIN":"D:\\go\\bin",
*     "GOPATH": "D:\\go\\gopath;D:\\go\\cpath", //您go的工作路径
*     "GOARCH":"386", //系统变量里面的 GOHOSTARCH ，386为32位平台,amd64为 64位平台
*     "GOOS":"windows", //系统里面的GOOS
*     "PATH":"%GOBIN%;%PATH%"
*    },
*    "comp_lint_enabled": true,   //打开这个才有下面的 comp_lint_commands标签里面的内容
*    "comp_lint_commands": [
*        {"cmd": ["go", "install"]}
*    ],
*    "on_save": [
*        {"cmd":"gs_comp_lint"}   //当按保存时以cmd自动执行的命令
*    ]
*}

### 2.5 Preferences-->Settings-User:
*{
*    "color_scheme": "Packages/Monokai Extended/Monokai Extended Light.tmTheme",
*    "draw_white_space": "all",
*    "font_size": 12,
*    "ignored_packages":
*    [
*        "Vintage"
*    ],
*    "tab_size": 4,
*    "translate_tabs_to_spaces": true,
*    "highlight_line": true,
*    "trim_trailing_white_space_on_save": true,
*    "atomic_save": true,
*    "auto_find_in_selection": true,
*    "highlight_modified_tabs": true,
*    "always_prompt_for_file_reload": tr*ue,
*}

### 2.6 Preferences--> Key Binbdings-User:
*[
*    { "keys": ["ctrl+f"], "command": "show_panel", "args": {"panel": "find", "reverse": false} },
*    // 删除当前行
*    *{ "keys": ["ctrl+d"], "command":"run_macro_file", "args": {"file":"Packages/Default/Delete Line.sublime-macro"} },
*]
*
*//=======================系统自带快捷键=======================//
*//=============选择=============//
*// Ctrl+L
*// 选择整行(按住-继续选择下行)
*//Ctrl+Shift+L
*//鼠标选中多行，按下 同时编辑这些行
*//鼠标中键
*//拖动，选择多行
*//Ctrl+左键点击
*//同时选中多个节点进行编辑
*//Ctrl+M
*// 光标移动至括号内开始或结束的位置
*
*// Ctrl+Shift+M
*// 选择括号内的内容(按住-继续选择父括号)
*//=============窗口=============//
*// SHIFT+ALT+数字
*// 分割窗口
*
*//=============行处理=============//
*// CTRL+J
*// 合并行JOIN
*// Ctrl+KU
*// 改为大写
*
*// Ctrl+KL
*// 改为小写
*// Ctrl+KK
*// 从光标处删除至行尾
*
*// Ctrl+Shift+D
*// 复制光标所在整行，插入在该行之前
*
*// Ctrl+J
*// 合并行(已选择需要合并的多行时)
*
*// Ctrl+/
*// 注释整行(如已选择内容，同“Ctrl+Shift+/”效果)
*
*// Ctrl+Shift+/
*// 注释已选择内容
*
*// Ctrl+Shift+V
*// 粘贴并自动缩进(其它兄弟写的，实测win系统自动缩进无效)
*
*// Ctrl+M
*// 光标跳至对应的括号
*
*// Alt+.
*// 闭合当前标签
*
*// Ctrl+Shift+A
*// 选择光标位置父标签对儿
*
*// Ctrl+Shift+[
*// 折叠代码
*
*// Ctrl+Shift+]
*// 展开代码
*
*// Ctrl+KT
*// 折叠属性
*
*// Ctrl+K0
*// 展开所有
*
*// Ctrl+U
*// 软撤销
*
*// Ctrl+T
*// 词互换
*
*// Ctrl+Enter
*// 插入行后
*
*// Ctrl+Shift Enter
*// 插入行前
*
*// Ctrl+K Backspace
*// 从光标处删除至行首
*
*// Shift+Tab
*// 去除缩进
*
*// Tab
*// 缩进
*
*// F9
*// 行排序(按a-z)

### 2.7 Preferences-->Package Settings-->Css Format-->Settings-User:
*{
*    "indentation": "    ",
*    "format_on_save": true,
*}

### 2.7 Preferences-->Package Settings-->CTags-->Settings-User:
*{
*    // Enable debugging
*    "debug": false,

*    // Enable autocomplete
*    "autocomplete": false,

*    // Alter this value if your ctags command is not in the PATH, or if using
*    // a different version of ctags to that in the path (i.e. for OSX).
*    //
*    // NOTE: You *should not* place entire commands here. These commands are
*    // built automatically using the values below. For example:
*    //   GOOD: "command": "/usr/bin/ctags"
*    //   BAD:  "command": "ctags -R -f .tags --exclude=some/path"
*    "command": "C:/Users/Administrator/AppData/Roaming/Sublime Text 3/Packages/FuncPreview/ctags.exe",

*    // Set to false to disable recursive search for ctag generation. This
*    // translates the `-R` parameter
*    "recursive" : true,

*    // Default read/write location of the tags file. This translates to the
*    // `-f [FILENAME]` parameter
*    "tag_file" : ".tags",

*    // Additional options to pass to ctags, i.e.
*    // ["--exclude=some/path", "--exclude=some/other/path", ...]
*    "opts" : [],

*    //
*    "filters": {
*        "source.python": {"type":"^i$"}
*    },

*    //
*    "definition_filters": {
*        "source.php": {"type":"^v$"}
*    },

*    //
*    "definition_current_first": true,

*    // Show the ctags menu in the context menus
*    "show_context_menus": true,

*    // Paths to additional tag files to include in tag search. This is a list
*    // of items in the form [["language", "platform"], "path"]
*    "extra_tag_paths": [[["source.python", "windows"], "C:\\Python27\\Lib\\tags"]],

*    // Additional tag files to search
*    "extra_tag_files": [".gemtags", "tags"],

*    // Set to false so as not to select searched symbol (in Vintage mode)
*    "select_searched_symbol": true
*}
