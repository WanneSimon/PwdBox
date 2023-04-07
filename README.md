## wails 中的哪些坑坑洼洼

### 前端调试
dev模式下，鼠标右键会弹出调试用的菜单（包括调试工具），正式打包会默认禁用

### 事件与数据
wails 基于事件通信，并没有数据双向绑定的能力。  
go函数返回给前端的对象只是一个复制对象，修改此对象并不会对 go 内部的原变量造成影响。  
同理，在启动时，暴露给前端的对象，在前端修改后，对象并不会被修改。
```
wails.Run(&options.App{
  ...
  Bind: []interface{}{
    // 假设 config 是配置数据，前端修改此config后，go中的config对象并不会变化
		config,
	},
  ...
})
```
注： `frontend/wailsjs/go` 下的文件只由在程序运行 `(wails dev)` 的时候才会更新

#### 解决方式
将功能和数据结构分离，就是传统的 `Service` 和 `PO` 思想。
暴露 `Service` 给前端，`PO` 所有的操作在 `Service` 中完成。  
例如：
```
type ConfigOp struct {
}
type AppConfig struct {
  name: string
}

var Config Appconfig

// 由于前端的修改并不会影响到后端，所以要求前端传入新的对象，
// 保存成功后再替换掉后端中的旧对象，这样就能保证后端和前端的数据一致
func (co *ConfigOp) Save(newAC AppConfig) bool {
  // TODO save ...
	Config = newAC
  return true
}
func (co *ConfigOp) Load() *AppConfig {
  // todo load from file
  Config = load....
  return Config
}
func (co *ConfigOp) Get() *AppConfig {
	return &Config
}

```
### 数据结构暴露给前端
前端如果使用 `js`，可以忽略。 `ts` 就要注意了！
在上面，我们定义了 `AppConfig` ，但这个数据结构不存在于前端中。如果使用 `ts`，使用过程中 `AppConfig`的类型是 `any`， 后续赋值和取值无法通过类型检查。
```
type AppConfig struct {
  name: string
}
```
#### 解决方式
为了让结构体 `AppConfig` 有对应的 `ts` 类型，需要给 `AppConfig` 的所有属性加上 `json` 的 `tag`。
```
type AppConfig struct {
  name: string  `json:"name"`
  //name: string  `json:"name" yaml:"name"`
}
```
在运行时就能在 `frontend/wailsjs/go/models.ts` 中看到 `AppConfig` 这个类了

### 暴露的对象太多，懒得写路径
官方示例中，前端调用事件是这样写的
```
import { Greet } from "/wailsjs/go/main/App";
//import { Load } from "/wailsjs/go/conf/AppConfig";
```
如果有其他对象，还要继续写声明，

在 `frontend/wailsjs` 下新建 `index.ts`
```
export * from "./go/main/App";
export * from "./runtime/runtime";

import * as ConfigOp from './go/conf/ConfigOps'

export {
  ConfigOp
}
```
调用方式就变成
```
import { Greet, ConfigOp } from "/wailsjs/index"
//import { Greet, ConfigOp } from "@/../wailsjs/index" 如果 ts 配置中使用了 @ 变量路径
```

### 事件是异步的
`wails` 暴露给前端的函数都是异步 `Promise` ，这就导致你想同步执行时就要写一堆 `async wait`。
```
  async mounted() { 
    await this.getConfigPath()
  },
  methods: {
    // 获取文件夹路径
    async getConfigPath() {
      console.log("ConfigOp", ConfigOp)
      let conf =  await ConfigOp.Get().then(res => res)
      console.log("ConfigOp-after", conf)
      ...
    },
  }
```

### 事件返回值
事件返回结构体时，在前端会自动转换成 `json` 对象。  
返回字节数组时，会自动 `base64` 编码，变成字符串。
```
// 读取文件 （返回的字节数组会自动base64编码成字符串）
// 如果是图片文件，在字符串前面加上 'data:image/png;base64,' 就可以直接挂载到 `<img>` 的 `src` 上
func (fo *FileOp) Open(path string) []byte {
  ...
}
```


