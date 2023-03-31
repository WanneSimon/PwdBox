// 参考: https://blog.csdn.net/u013022210/article/details/55101558

/**
 * 字符串转字节数组
 * @param {*} str
 * @returns
 */
 function stringToByte(str) {
	var bytes = new Array();
	var len, c;
	len = str.length;
	for(var i = 0; i < len; i++) {
		c = str.charCodeAt(i);
		if(c >= 0x010000 && c <= 0x10FFFF) {
			bytes.push(((c >> 18) & 0x07) | 0xF0);
			bytes.push(((c >> 12) & 0x3F) | 0x80);
			bytes.push(((c >> 6) & 0x3F) | 0x80);
			bytes.push((c & 0x3F) | 0x80);
		} else if(c >= 0x000800 && c <= 0x00FFFF) {
			bytes.push(((c >> 12) & 0x0F) | 0xE0);
			bytes.push(((c >> 6) & 0x3F) | 0x80);
			bytes.push((c & 0x3F) | 0x80);
		} else if(c >= 0x000080 && c <= 0x0007FF) {
			bytes.push(((c >> 6) & 0x1F) | 0xC0);
			bytes.push((c & 0x3F) | 0x80);
		} else {
			bytes.push(c & 0xFF);
		}
	}
	return bytes;
}

/**
 * 字符串转字节数组
 * @param {*} str
 * @returns
 */
 function stringToUint8ByteArray(str) {
	let arr = stringToByte(str)
  return new Uint8Array(arr)
}

/**
 * 字节数组转字符串
 * @param {*} arr
 * @returns
 */
function byteToString(arr) {
	if(typeof arr === 'string') {
		return arr;
	}
	var str = '',
		_arr = arr;
	for(var i = 0; i < _arr.length; i++) {
		var one = _arr[i].toString(2),
			v = one.match(/^1+?(?=0)/);
		if(v && one.length == 8) {
			var bytesLength = v[0].length;
			var store = _arr[i].toString(2).slice(7 - bytesLength);
			for(var st = 1; st < bytesLength; st++) {
				store += _arr[st + i].toString(2).slice(2);
			}
			str += String.fromCharCode(parseInt(store, 2));
			i += bytesLength - 1;
		} else {
			str += String.fromCharCode(_arr[i]);
		}
	}
	return str;
}

/**
 * 字节数组每个字节取反
 * @param {*} bytesArr
 * @returns
 */
function reverseBytesArray(bytesArr) {
  let reverseBytes = []
  for(let i=0; i<bytesArr.length; i++) {
    reverseBytes[i] = ~bytesArr[i]
  }
  return reverseBytes
}

function reverseBytes(bytesArr, emptyContainer) {
  for(let i=0; i<bytesArr.length; i++) {
    emptyContainer[i] = ~bytesArr[i]
  }
  return emptyContainer
}

function testBytesAndString() {
  let str = "sfajl29348+-';`%(#}"
  let bs = stringToByte(str)
  console.log('bytes', bs)
  console.log('str', byteToString(bs))

  let reverseBs = []
  for(let i=0; i<bs.length; i++) {
    reverseBs[i] = ~bs[i]
  }
  console.log('rever', reverseBs)

  let originBs = []
  for(let i=0; i<reverseBs.length; i++) {
    originBs[i] = ~reverseBs[i]
  }
  console.log('origi', originBs)
  console.log('origin-str', byteToString(originBs))
}

function testBase64() {
  console.log("testBase64")

}

export default {
  stringToByte, byteToString, reverseBytesArray, reverseBytes, stringToUint8ByteArray,
}