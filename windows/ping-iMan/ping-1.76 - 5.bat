@echo off
:: ping�ĵ�ַ����־�ļ�
set pingServer=10.10.1.76
set logFile="ping-%pingServer%.log"
:: sleepTime=5 �Ƿ�����
set sleepTime=5
::����Ping���֮��д����־
set logWriteCount=999

::д����־ʱ��¼�Ĵ���
set /a logCount=1
set /a iteration=1
set /a sleepDo=%sleepTime% * 60
:Do
set /a num+=1
echo --------���ڽ��е� %iteration% ����, �� %num% �� ping [%pingServer%], sleepTime = %sleepTime%����---------
ping %pingServer% -n 1
sleep.exe %sleepDo%

if %num% == %logWriteCount% (
    set /a logCount+=1
    set /a iteration+=1
    echo -------log��¼�� %logCount% ��, sleepTime = %sleepTime%����, ��ǰʱ�� %Date% - %Time%------- >> %logFile%
    ping 10.10.1.76 -n 1 >> %logFile%
    set /a num=0
)

echo.
goto Do