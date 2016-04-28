@echo off
@if exist "rsrc.syso" (
    @del "rsrc.syso"
)

rsrc -manifest main.manifest -o rsrc.syso -ico main.ico

go build  -ldflags "-s -w -H=windowsgui"
pause
exit