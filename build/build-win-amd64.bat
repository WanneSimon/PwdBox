cd ..
wails build ^
::: https://wails.io/docs/reference/cli#platforms
-platform windows/amd64^
-upx -debug false ^
-o pwdbox-win-amd64.exe
build/messy.bat
PAUSE

