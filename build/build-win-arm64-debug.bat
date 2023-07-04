cd ..
wails build ^
::: https://wails.io/docs/reference/cli#platforms
-platform windows/arm64 ^
-upx -debug true -windowsconsole  ^
-o pwdbox-win-arm64-debug.exe
build/messy.bat
PAUSE

