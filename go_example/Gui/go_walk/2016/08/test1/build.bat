@if exist "rsrc.syso" (
    @del "rsrc.syso"
)

rsrc -manifest main.manifest -o rsrc.syso
go build 
pause
exit
-ldflags="-H windowsgui"
-ldflags "-s -w"
