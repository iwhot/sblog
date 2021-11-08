import CryptoJS from "crypto-js"

// 加密函数
export function encrypt(data, KEY, IV) {
	var encrypted = CryptoJS.AES.encrypt(data, CryptoJS.enc.Utf8.parse(KEY), {
		iv: CryptoJS.enc.Utf8.parse(IV),
		mode: CryptoJS.mode.CBC,
		padding: CryptoJS.pad.Pkcs7
	});
	return encrypted.toString(); //返回的是base64格式的密文
}

//解密
export function decrypt(encrypted, KEY, IV) {
	var decrypted = CryptoJS.AES.decrypt(encrypted, CryptoJS.enc.Utf8.parse(KEY), {
		iv: CryptoJS.enc.Utf8.parse(IV),
		mode: CryptoJS.mode.CBC,
		padding: CryptoJS.pad.Pkcs7
	});
	return decrypted.toString(CryptoJS.enc.Utf8);
}
