const USER_STATUS = {
  "normal": "正常",
  "ban": "封禁",
  "delete": "已删除",
}

const userStatus = (code) => {
  let status = USER_STATUS[code]
  return status ? status : code
}

export default {
  USER_STATUS, userStatus,
}