import showdown from "showdown"

export function createShowDownConverter() {
  // @see https://github.com/showdownjs/showdown/
  // 方式1
  // let converter = new showdown.Converter()
  // // 启用 table 表格
  // converter.setOption("tables", true)

  // 方式2
  // let converter = new showdown.Converter({
  //   tables: true,  // 支持表格
  //   headerLevelStart: 1, // # 将使用 h1 来表示
  //   parseImgDimensions: true, // 开启md语法中对图片尺寸的支持
  //   simplifiedAutoLink: false,  // 是否自动给链接文字添加超链接
  //   prefixHeaderId: false,  // 是否自动给标题的id添加 'section' 的前缀，可以用于目录定位
  //   openLinksInNewWindow: true, // 新窗口中打开超链接
  //   metadata: true, // 支持md的顶部元数据, 使用 <<< 和 >>>（或---和 ---）包裹的内容
  // });

  // 方式3
  let converter = new showdown.Converter()
  converter.setFlavor('github')  // 使用预设配置 'original', 'vanilla', 'github'
  converter.setOption("ghCodeBlocks", true)
  converter.setOption("tables", true)
  return converter
}
