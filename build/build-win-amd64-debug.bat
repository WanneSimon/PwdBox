cd ..
wails build ^
::: https://wails.io/docs/reference/cli#platforms
-platform windows/amd64 ^
-upx -debug true -windowsconsole  ^
-o pwdbox-win-amd64-debug.exe
build/messy.bat
PAUSE

