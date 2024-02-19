package check_phone

import (
	"github.com/gonozov0/weddingtgbot/pkg/phone_utils"
)

var phoneWhitelist = []string{
	phone_utils.Normalize("+7 903 691 9544"),
	phone_utils.Normalize("+7 915 998 6573"),
}

func isPhoneInvited(phone string) bool {
	for _, p := range phoneWhitelist {
		if p == phone_utils.Normalize(phone) {
			return true
		}
	}

	return false
}
