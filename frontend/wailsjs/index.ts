// 使用方式 import { Greet } from "../wailsjs";  import { Greet } from "/wailsjs";
export * from "./go/main/App";
export * from "./runtime/runtime";

import * as FileOp from './go/env/FileOp'
import * as ConfigOp from './go/conf/ConfigOps'
import * as DbOp from './go/pwdbox/DbOp'
import * as PlatformService from './go/pwdbox/PlatformService'
import * as AccountService from './go/pwdbox/AccountService'
import * as PwdTool from './go/pwdbox/PwdTool'
import * as DataOutOp from './go/dataop/DataOutOp'

export {
  FileOp, ConfigOp, DbOp, PlatformService, AccountService, PwdTool, DataOutOp,
}

