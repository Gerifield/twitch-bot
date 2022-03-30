package vods

import "github.com/gerifield/twitch-bot/model"

func Handle(_ *model.Message) (string, error) {
	return "VOD csatorna: https://youtube.com/playlist?list=PLOIpbzPG9JERa1-smdKCXC1r9a9X_fID2", nil
}
