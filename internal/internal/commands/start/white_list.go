package start

var loginWhitelist = []string{
	"gonozov0",
	"TaoGen",
}

func isLoginInvited(login string) bool {
	for _, whitelistedLogin := range loginWhitelist {
		if login == whitelistedLogin {
			return true
		}
	}
	return false
}
