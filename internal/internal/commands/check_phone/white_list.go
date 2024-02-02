package check_phone

import (
	"github.com/gonozov0/weddingtgbot/pkg/phone_utils"
)

var phoneWhitelist = []string{
	"+7 915 979 6484",
	"+7 915 998 6573",
}

func isPhoneInvited(phone string) bool {
	for _, p := range phoneWhitelist {
		if phone_utils.NormalizePhoneNumber(p) == phone_utils.NormalizePhoneNumber(phone) {
			return true
		}
	}

	return false
}
