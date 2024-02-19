package shared

import (
	"github.com/gonozov0/weddingtgbot/pkg/phone_utils"
)

var (
	namesMap = map[string]string{
		phone_utils.Normalize("+7 903 691 9544"): "Мама Юля",
		phone_utils.Normalize("+7 910 826 9735"): "Папа Сергей",
		phone_utils.Normalize("+7 915 998 6573"): "Дмитрий",
		"gonozov0":                               "Дмитрий",
		"TaoGen":                                 "Ольга",
	}

	surnamesMap = map[string]string{
		phone_utils.Normalize("+7 903 691 9544"): "Гонозова",
		phone_utils.Normalize("+7 910 826 9735"): "Гонозов",
		phone_utils.Normalize("+7 915 998 6573"): "Гонозов",
		"gonozov0":                               "Гонозов",
		"TaoGen":                                 "почти Гонозова",
	}
)

type PersonInfo struct {
	Name    string
	Surname string
}

func (p PersonInfo) GetFullName() string {
	return p.Name + " " + p.Surname
}

func GetPersonInfo(login string) PersonInfo {
	return PersonInfo{
		Name:    namesMap[login],
		Surname: surnamesMap[login],
	}
}
