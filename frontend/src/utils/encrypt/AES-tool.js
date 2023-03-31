// https://wenku.baidu.com/view/beb92a2db868a98271fe910ef12d2af90242a83c.html

import CryptoJS from "crypto-js"
// https://www.npmjs.com/package/base-64
// import { Base64 } from 'js-base64'

import globalStore from '/src/components/pinia/store.js'

// const key = () => "formeformeforme0";
// const iv = () => "0123456789abcdef";

const aesData = () => {
  let store = globalStore()
  const aes = store?.aes
  // const aes = {
  //   key : "formeformeforme0",
  //   iv  : "0123456789abcdef"
  // }

  return {
    key: CryptoJS.enc.Utf8.parse(aes?.key),
    iv : CryptoJS.enc.Utf8.parse(aes?.iv),
  }
}

/** 编码*/
function encrypt(raw) {
  let data = aesData()
  let aesKey = data.key
  let aesIv =  data.iv

  const encryptStr = CryptoJS.AES.encrypt(raw, aesKey, {
    iv: aesIv,
    mode: CryptoJS.mode.CBC,
  })

  return encryptStr.toString()
}

/** 解码*/
function decrypt(str) {
  let data = aesData()
  let aesKey = data.key
  let aesIv =  data.iv

  let encryptStr = CryptoJS.AES.decrypt(str, aesKey, {
    iv: aesIv,
    mode: CryptoJS.mode.CBC,
    padding: CryptoJS.pad.Pkcs7
  })
  return encryptStr.toString(CryptoJS.enc.Utf8)
}

function testAesInner() {
  console.log("AES-test")

  // const key = () => "formeformeforme0"; // 16位
  // const iv = () => "0123456789abcdef"; // 16位

  let str = "01234567890abcd"
  // let str = "2.	结果分析的颜色"

  const en = encrypt(str)
  const de = decrypt(en)
  console.log("en", en, ("R1lT07wIxuEhnWV2OVkd3w==" == en ))
  console.log("de", de, ("01234567890abcd" == de), de.length)

}

function testAes() {
  // testAesInner()
}

export default {
  encrypt, decrypt, testAes
}