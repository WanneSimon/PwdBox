cd ..
wails build ^
::: https://wails.io/docs/reference/cli#platforms
-platform windows/arm64^
-upx -debug false ^
-o pwdbox-win-arm64.exe
build/messy.bat
PAUSE

