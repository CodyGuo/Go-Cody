# Windows
使用windowns过程中遇到的问题

# 目录
## 1. Newsid
* win7使用newsid后蓝屏解决方法
* 1> 使用U盘等类似工具，进入PE系统
* 2> 打开C:\Windows\SysWOW64\config\Newsid Backup
* 3> 将里边所有的文件拷贝到C:\Windows\System32下，重启电脑
* 那么WIN7、WIN8下修改SID的方法如下：找到并运行 C:\Windows\System32\Sysprep\sysprep.exe 默认选择了"进入系统全新体验(OOBE)",勾选"通用",点击"确定"或者在cmd下执行如下命令: sysprep /oobe /generalize 。