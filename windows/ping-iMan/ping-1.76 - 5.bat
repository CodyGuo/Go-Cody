@echo off
:: ping的地址和日志文件
set pingServer=10.10.1.76
set logFile="ping-%pingServer%.log"
:: sleepTime=5 是分钟数
set sleepTime=5
::连续Ping多次之后写入日志
set logWriteCount=999

::写入日志时记录的次数
set /a logCount=1
set /a iteration=1
set /a sleepDo=%sleepTime% * 60
:Do
set /a num+=1
echo --------正在进行第 %iteration% 迭代, 第 %num% 次 ping [%pingServer%], sleepTime = %sleepTime%分钟---------
ping %pingServer% -n 1
sleep.exe %sleepDo%

if %num% == %logWriteCount% (
    set /a logCount+=1
    set /a iteration+=1
    echo -------log记录第 %logCount% 次, sleepTime = %sleepTime%分钟, 当前时间 %Date% - %Time%------- >> %logFile%
    ping 10.10.1.76 -n 1 >> %logFile%
    set /a num=0
)

echo.
goto Do