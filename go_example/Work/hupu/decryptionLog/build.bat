@echo off
@if exist "rsrc.syso" (
    @del "rsrc.syso"
)

::go-bindata-assetfs.exe bin/...
rsrc -manifest main.manifest -o rsrc.syso -ico main.ico

go build  -ldflags "-s -w -H=windowsgui"
pause
exit
