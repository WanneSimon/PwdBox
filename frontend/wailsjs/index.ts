// 使用方式 import { Greet } from "../wailsjs";  import { Greet } from "/wailsjs";
export * from "./go/main/App";
export * from "./runtime/runtime";

import * as FileOp from './go/env/FileOp'
import * as ConfigOp from './go/conf/ConfigOps'

export {
  FileOp, ConfigOp
}

