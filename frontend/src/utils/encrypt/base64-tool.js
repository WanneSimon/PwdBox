// https://www.npmjs.com/package/js-base64
import { Base64 } from 'js-base64'
import byteStringTool from './byte-string.js'
// 可以使用 window.atob 和 window.btoa 替换 base64 解码和编码部分
/*
 encode:
var text = 'foo © bar 𝌆 baz';
var bytes = utf8.encode(text);
var encoded = base64.encode(bytes);
console.log(encoded);

 decode:
var encoded = 'Zm9vIMKpIGJhciDwnYyGIGJheg==';
var bytes = base64.decode(encoded);
var text = utf8.decode(bytes);
console.log(text);
*/

/*
Base64-encode: MDEyMzQ1Njc4OTA=
Base64-decode: 01234567890
======================
Base64-2-encode: z87NzMvKycjHxs8=
Base64-2-decode: 01234567890
*/


function encode(raw) {
  let encodeStr = Base64.encode(raw)
  return encodeStr
}

function decode(encodeStr) {
  let decodedStr = Base64.decode(encodeStr)
  let originStr = decodedStr
  return originStr
} 

/** 后端返回的信息编码
 * 字符串字节取反，然后编码
 * @param {*} raw 字符串
 * @returns 
 */
function encodeResData(raw) {
  // 1. 获取字节
  let bytes = byteStringTool.stringToUint8ByteArray(raw)
  // console.log("en-origin: " + bytes)

  // 2. 字节取反
  let reversedArr = new Uint8Array(bytes.length)
  byteStringTool.reverseBytes(bytes, reversedArr)
  // console.log("en-reverse: " + reversedArr)

  // 3. 编码 
  let encodedStr = Base64.fromUint8Array(reversedArr)
  // console.log("encodedStr: " + encodedStr)

  return encodedStr
}
/** 对后端返回的信息解码
 * “字符串字节取反，然后编码” 的逆操作
 * @param {*} raw 
 * @returns 
 */
function decodeResData(encodeStr) {
  // 1. 解码
  let decodedUint8Arr = Base64.toUint8Array(encodeStr)
  // console.log('de-reverse: ' + decodedUint8Arr)

  // 2. 字节取反
  let reversedArr = new Uint8Array(decodedUint8Arr.length)
  byteStringTool.reverseBytes(decodedUint8Arr, reversedArr)
  // console.log("de-origin: " + reversedArr)

  // 3. 恢复原字符串
  let origin = byteStringTool.byteToString(reversedArr)
  return origin
}

/** 编码请求
 * @param {*} raw 原 json对象 或纯字符串
 * @returns
 */
function encodeReqData(raw) {
  let rawData = raw
  if(typeof(raw) != 'string') {
    rawData = JSON.stringify(raw)
  }
  return encodeResData(rawData)
}
/** 解码请求数据，返回 json 或 字符串
 * 
 * @param {*} encodeStr 
 * @param {*} notJson 不需要解析成 json 对象时，需要指定为 true
 * @returns 
 */
function decodeReqData(encodeStr, notJson) {
  let str = decodeResData(encodeStr)
  if(notJson) {
    return str
  }
  return JSON.parse(str)
}

function testBase64() {
  let str = "01234567890"
      // let str = "&898省圣诞节拉风地方"
      let encode = this.encode(str)
      let decode = this.decode(encode)
      console.log("encode", encode, ('MDEyMzQ1Njc4OTA='===encode))
      console.log("decode", decode)

      // let str2 = "&898省圣诞节拉风地方"
      let str2 = "01234567890"
      let encode2 = this.encodeResData(str2)
      let decode2 = this.decodeResData(encode2)
      console.log("encode2", encode2, ('z87NzMvKycjHxs8='===encode2))
      console.log("decode2", decode2)
}

export default {
  testBase64, encode, decode, encodeResData, decodeResData, encodeReqData, decodeReqData
}