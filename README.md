### 使用
1. 下载并安装 [webview2](https://developer.microsoft.com/n-us/microsoft-edge/webview2/)  
2. 从 [realease](/realease) 下载程序包，解压并运行 `pwdbox.exe`

### 相关技术
`go-1.18+` 
[wails](https://wails.io/docs/reference/cli) 
[webview2](https://developer.microsoft.com/n-us/microsoft-edge/webview2/) 
`node-v14.18.1`  `vue3`  `vite` 


### 构建
`build` 目录下有不同的构建脚本，运行即可。  

`wails` 对构建参数 `-webview2` 的说明和一般理解不太一样。  
`-webview2 download` 构建时会下载 `webview2` 并压缩进程序中。
构建时，使用 `-webview2 embed` 比不使用 `-webview2` 时的输出文件大，推荐不使用此选项。  

构建后，可执行文件同级目录下需要 `config/saya.yml` 和 `config/pwdbox.db3` 才能使用，对应仓库中 `config/saya.yml` 和 `config/pwdbox-init.db3` 。
