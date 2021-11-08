
export function getToken() {
	var token = window.localStorage.getItem('token')
	if (!token) {
		return ""
	}

	return token
}

export function getUserInfo() {
	var user = window.localStorage.getItem('user')
	if (!user) {
		return []
	}
	return JSON.parse(user)
}