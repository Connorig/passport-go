/**
 * window.localStorage 浏览器永久缓存
 * @method set 设置永久缓存
 * @method get 获取永久缓存
 * @method remove 移除永久缓存
 * @method clear 移除全部永久缓存
 */
export const Local = {
	// 设置永久缓存
	set(key: string, val: any) {
		window.localStorage.setItem(key, JSON.stringify(val));
	},
	// 获取永久缓存
	get(key: string) {
		let json: any = window.localStorage.getItem(key);
		return JSON.parse(json);
	},
	// 移除永久缓存
	remove(key: string) {
		window.localStorage.removeItem(key);
	},
	// 移除全部永久缓存
	clear() {
		window.localStorage.clear();
	},
};

/**
 * window.sessionStorage 浏览器临时缓存
 * @method set 设置临时缓存
 * @method get 获取临时缓存
 * @method remove 移除临时缓存
 * @method clear 移除全部临时缓存
 */
export const Session = {
	// 设置临时缓存
	set(key: string, val: any) {
		window.sessionStorage.setItem(key, JSON.stringify(val));
	},
	// 获取临时缓存
	get(key: string) {
		let json: any = window.sessionStorage.getItem(key);
		return JSON.parse(json);
	},
	// 移除临时缓存
	remove(key: string) {
		window.sessionStorage.removeItem(key);
	},
	// 移除全部临时缓存
	clear() {
		window.sessionStorage.clear();
	},
};

// var operator = "="
// function getCookieValue(keyStr: string) {
// 	var value = null;
// 	var s = window.document.cookie;
// 	var arr = s.split("; ");
// 	for (var i = 0; i < arr.length; i++) {
// 		var str = arr[i];
// 		var k = str.split(operator)[0];
// 		var v = str.split(operator)[1];
// 		if (k == keyStr) {
// 			value = v;
// 			break;
// 		}
// 	}
// 	return value;
// }

// function setCookieValue(key: string, value: any) {
// 	document.cookie = key + operator + value;
// }