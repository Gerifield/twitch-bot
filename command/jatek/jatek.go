package jatek

import (
	"fmt"

	"github.com/gerifield/twitch-bot/model"
)

var users []string

func init() {
	users = make([]string, 0)
}

func Handle(msg *model.Message) (string, error) {
	usr := msg.User()

	if userExistst(usr) {
		return fmt.Sprintf("%s mar regisztralva van!", usr), nil
	}
	users = append(users, usr)
	return fmt.Sprintf("%s regisztralva!", usr), nil
}

func userExistst(user string) bool {
	for _, u := range users {
		if u == user {
			return true
		}
	}

	return false
}
