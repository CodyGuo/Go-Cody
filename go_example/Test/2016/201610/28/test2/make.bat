windres -o %GOPATH%\src\uac\main_res.syso %GOPATH%\src\uac\uac.rc
go build
del %GOPATH%\src\uac\main_res.syso
