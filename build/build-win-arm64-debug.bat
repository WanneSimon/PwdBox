cd ..
wails build ^
::: https://wails.io/docs/reference/cli#platforms
-platform windows/arm64 ^
-upx -webview2 download -debug true -windowsconsole  ^
-o pwdbox-win-arm64-debug.exe
build/messy.bat
PAUSE

