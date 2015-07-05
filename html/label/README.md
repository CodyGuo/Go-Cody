# label
常用的HTML标签

# 目录
## 1. 跑马灯
标签 | 功能
------------- | -------------
`<marquee>...</marquee>` | 普通卷动  
`<marquee behavior=slide>...</marquee>` | 滑动
`<marquee behavior=scroll>...</marquee>` | 预设卷动
`<marquee behavior=alternate>...</marquee>` | 来回卷动
`<marquee direction=down>...</marquee>` | 向下卷动
`<marquee direction=up>...</marquee>` | 向上卷动
`<marquee direction=right></marquee>` | 向右卷动
`<marquee direction=left></marquee>` | 向左卷动
`<marquee loop=2>...</marquee>` | 卷动次数
`<marquee width=180>...</marquee>` | 设定宽度
`<marquee height=30>...</marquee>` | 设定高度
`<marquee bgcolor=FF0000>...</marquee>` | 设定背景颜色
`<marquee scrollamount=30>...</marquee>` | 设定卷动距离
`<marquee scrolldelay=300>...</marquee>` | 设定卷动时间

## 2. 字体效果
标签 | 功能
------------- | -------------
`<h1>...</h1>` | 标题字(最大)
`<h6>...</h6>` | 标题字(最小)
`<b>...</b>` | 粗体字
`<strong>...</strong>` | 粗体字(强调)
`<i>...</i>` | 斜体字
`<em>...</em>` | 斜体字(强调)
`<dfn>...</dfn>` | 斜体字(表示定义)
`<u>...</u>` | 底线
`<ins>...</ins>` | 底线(表示插入文字)
`<strike>...</strike>` | 横线
`<s>...</s>` | 删除线
`<del>...</del>` | 删除线(表示删除)
`<kbd>...</kbd>` | 键盘文字
`<tt>...</tt>` | 打字体
`<xmp>...</xmp>` | 固定宽度字体(在文件中空白、换行、定位功能有效)
`<plaintext>...</plaintext>` | 固定宽度字体(不执行标记符号)
`<listing>...</listing>` | 固定宽度小字体
`<font color=00ff00>...</font>` | 字体颜色
`<font size=1>...</font>` | 最小字体
`<font style =font-size:100 px>...</font>` | 无限增大

## 3. 区断标记
标签 | 功能
------------- | -------------
`<hr>` | 水平线
`<hr size=9>` | 水平线(设定大小)
`<hr width=80%>` | 水平线(设定宽度)
`<hr color=ff0000>` | 水平线(设定颜色)
`<br>` | (换行)
`<nobr>...</nobr>` | 水域(不换行)
`<p>...</p>` | 水域(段落)
`<center>...</center>` | 置中

## 4. 链接
标签 | 功能
------------- | -------------
`<base href=地址>` | (预设好连结路径)
`<a href=地址></a>` | 外部连结
`<a href=地址 target=_blank></a>` | 外部连结(另开新窗口)
`<a href=地址 target=_top></a>` | 外部连结(全窗口连结)
`<a href=地址 target=页框名></a>` | 外部连结(在指定页框连结)

## 5. 图像/音乐
标签 | 功能
------------- | -------------
`<img src=图片地址>`  贴图
`<img src=图片地址 width=180>` | 设定图片宽度
`<img src=图片地址 height=30>` | 设定图片高度
`<img src=图片地址 alt=提示文字>` | 设定图片提示文字
`<img src=图片地址 border=1>` | 设定图片边框
`<bgsound src=MID音乐文件地址>` | 背景音乐设定

## 6. 表格
标签 | 功能
------------- | -------------
`<table aling=left>...</table>` | 表格位置,置左
`<table aling=center>...</table>` | 表格位置,置中
`<table background=图片路径>...</table>` | 背景图片的URL=就是路径网址
`<table border=边框大小>...</table>` | 设定表格边框大小(使用数字)
`<table bgcolor=颜色码>...</table>` | 设定表格的背景颜色
`<table borderclor=颜色码>...</table>` | 设定表格边框的颜色
`<table borderclordark=颜色码>...</table>` | 设定表格暗边框的颜色
`<table borderclorlight=颜色码>...</table>` | 设定表格亮边框的颜色
`<table cellpadding=参数>...</table>` | 指定内容与网格线之间的间距(使用数字)
`<table cellspacing=参数>...</table>` | 指定网格线与网格线之间的距离(使用数字)
`<table cols=参数>...</table>` | 指定表格的栏数
`<table frame=参数>...</table>` | 设定表格外框线的显示方式
`<table width=宽度>...</table>` | 指定表格的宽度大小(使用数字)
`<table height=高度>...</table>` | 指定表格的高度大小(使用数字)
`<td colspan=参数>...</td>` | 指定储存格合并栏的栏数(使用数字)
`<td rowspan=参数>...</td>` | 指定储存格合并列的列数(使用数字)

## 7. 分割窗口
标签 | 功能
------------- | -------------
`<frameset cols="20%,*">` | 左右分割,将左边框架分割大小为20%右边框架的大小浏览器会自动调整
`<frameset rows="20%,*">` | 上下分割,将上面框架分割大小为20%下面框架的大小浏览器会自动调整
`<frameset cols="20%,*">` | 分割左右两个框架
`<frameset cols="20%,*,20%">` | 分割左中右三个框架 
`<frameset rows="20%,*,20%">` | 分割上中下三个框架
`<! - - ... - ->` | 批注
`<A HREF TARGET>` | 指定超级链接的分割窗口
`<A HREF=#锚的名称>` | 指定锚名称的超级链接
`<A HREF>` | 指定超级链接
`<A NAME=锚的名称>` | 被连结点的名称
`<ADDRESS>....</ADDRESS>` | 用来显示电子邮箱地址
`<B>` | 粗体字
`<BASE TARGET>` | 指定超级链接的分割窗口
`<BASEFONT SIZE>` | 更改预设字形大小
`<BGSOUND SRC>` | 加入背景音乐
`<BIG>` | 显示大字体
`<BLINK>` | 闪烁的文字
`<BODY TEXT LINK VLINK>` | 设定文字颜色
`<BODY>` | 显示本文
`<BR>` | 换行
`<CAPTION ALIGN>` | 设定表格标题位置
`<CAPTION>...</CAPTION>` | 为表格加上标题
`<CENTER>` | 向中对齐
`<CITE>...<CITE>` | 用于引经据典的文字
`<CODE>...</CODE>` | 用于列出一段程序代码
`<COMMENT>...</COMMENT>` | 加上批注
`<DD>` | 设定定义列表的项目解说
`<DFN>...</DFN>` | 显示"定义"文字
`<DIR>...</DIR>` | 列表文字卷标
`<DL>...</DL>` | 设定定义列表的卷标
`<DT>` | 设定定义列表的项目
`<EM>` | 强调之用
`<FONT FACE>` | 任意指定所用的字形
`<FONT SIZE>` | 设定字体大小
`<FORM ACTION>` | 设定户动式窗体的处理方式
`<FORM METHOD>` | 设定户动式窗体之资料传送方式
`<FRAME MARGINHEIGHT>` | 设定窗口的上下边界
`<FRAME MARGINWIDTH>` | 设定窗口的左右边界
`<FRAME NAME>` | 为分割窗口命名
`<FRAME NORESIZE>` | 锁住分割窗口的大小
`<FRAME SCROLLING>` | 设定分割窗口的滚动条
`<FRAME SRC>` | 将HTML文件加入窗口
`<FRAMESET COLS>` | 将窗口分割成左右的子窗口
`<FRAMESET ROWS>` | 将窗口分割成上下的子窗口
`<FRAMESET>...</FRAMESET>` | 划分分割窗口
`<H1>~<H6>` | 设定文字大小
`<HEAD   >` | 标示文件信息
`<HR>` | 加上分网格线
`<HTML>` | 文件的开始与结束
`<I>` | 斜体字
`<IMG ALIGN>` | 调整图形影像的位置
`<IMG ALT>` | 为你的图形影像加注
`<IMG DYNSRC LOOP>` | 加入影片
`<IMG HEIGHT WIDTH>` | 插入图片并预设图形大小
`<IMG HSPACE>` | 插入图片并预设图形的左右边界
`<IMG LOWSRC>` | 预载图片功能
`<IMG SRC BORDER>` | 设定图片边界
`<IMG SRC>` | 插入图片
`<IMG VSPACE>` | 插入图片并预设图形的上下边界
`<INPUT TYPE NAME value>` | 在窗体中加入输入字段
`<ISINDEX>` | 定义查询用窗体
`<KBD>...</KBD>` | 表示使用者输入文字
`<LI TYPE>...</LI>` | 列表的项目 ( 可指定符号 )
`<MARQUEE>` | 跑马灯效果
`<MENU>...</MENU>` | 条列文字卷标
`<META NAME="REFRESH" CONTENT URL>` | 自动更新文件内容
`<MULTIPLE>` | 可同时选择多项的列表栏
`<NOFRAME>` | 定义不出现分割窗口的文字
`<OL>...</OL>` | 有序号的列表
`<OPTION>` | 定义窗体中列表栏的项目
`<P ALIGN>` | 设定对齐方向
`<P>` | 分段
`<PERSON>...</PERSON>` | 显示人名
`<PRE>` | 使用原有排列
`<SAMP>...</SAMP>` | 用于引用字
`<SELECT>...</SELECT>` | 在窗体中定义列表栏
`<SMALL>` | 显示小字体
`<STRIKE>` | 文字加横线
`<STRONG>` | 用于加强语气
`<SUB>` | 下标字
`<SUP>` | 上标字
`<TABLE BORDER=n>` | 调整表格的宽线高度
`<TABLE CELLPADDING>` | 调整数据域位之边界
`<TABLE CELLSPACING>` | 调整表格线的宽度
`<TABLE HEIGHT>` | 调整表格的高度
`<TABLE WIDTH>` | 调整表格的宽度
`<TABLE>...</TABLE>` | 产生表格的卷标
`<TD ALIGN>` | 调整表格字段之左右对齐
`<TD BGCOLOR>` | 设定表格字段之背景颜色
`<TD COLSPAN ROWSPAN>` | 表格字段的合并
`<TD NOWRAP>` | 设定表格字段不换行
`<TD VALIGN>` | 调整表格字段之上下对齐
`<TD WIDTH>` | 调整表格字段宽度
`<TD>...</TD>` | 定义表格的数据域位
`<TEXTAREA NAME ROWS COLS>` | 窗体中加入多少列的文字输入栏
`<TEXTAREA WRAP>` | 决定文字输入栏是自动否换行
`<TH>...</TH>` | 定义表格的标头字段
`<TITLE>` | 文件标题
`<TR>...</TR>` | 定义表格美一行
`<TT>` | 打字机字体
`<U>` | 文字加底线
`<UL TYPE>...</UL>` | 无序号的列表 ( 可指定符号 )
`<VAR>...</VAR>` | 用于显示变量
