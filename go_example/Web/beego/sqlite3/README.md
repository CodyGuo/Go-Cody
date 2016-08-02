# sqlite3安装
windows、linux在安装go-sqlite3时遇到的问题
# 目录
## 1. Windows上安装sqlite
* 教程内容来自 http://my.oschina.net/Obahua/blog/129689   

### 1.1. Win7上的安装
* 1> 安装过程中出现的错误。  
  ![](https://github.com/CodyGuo/Go-Cody/blob/master/beego/sqlite3/image/go-sqlite3-err.png)  
* 2> 经过无闻的帮助，找到了原因是cgo被禁用了。  
 ![](https://github.com/CodyGuo/Go-Cody/blob/master/beego/sqlite3/image/go%20env%20err.png)
* 3> 正确的cgo配置如下，set CGO_ENABLED=1，windows下修改环境变量即可。
 ![](https://github.com/CodyGuo/Go-Cody/blob/master/beego/sqlite3/image/go%20env%20ok.jpg)
