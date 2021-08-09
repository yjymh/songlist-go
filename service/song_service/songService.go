package song_service

import (
	"github.com/labstack/gommon/log"
	"github.com/yjymh/songlist-go/model"
	"github.com/yjymh/songlist-go/module/requests/music"
)

// AddSongInfo 添加新的歌曲
func AddSongInfo(title string) bool {
	var songInfo *model.SongInfo

	model.DB().Where("title=?", title).First(&songInfo)

	if songInfo.SongId != 0 {
		log.Infof("%s已经添加了", title)
		return true
	}

	songInfo, _ = music.GetMusicInfoByQQ(title)

	if songInfo == nil {
		log.Info("添加歌曲失败，没有查到该歌曲，请确认是否输入错误")
		return false
	}

	err := model.DB().Create(&songInfo).Error
	if err != nil {
		log.Error("数据插入失败")
		return false
	}
	return true
}

// QuerySongInfo 查询歌曲信息
// TODO 还不能指定用户
func QuerySongInfo(user string) ([]model.SongInfo, error) {
	var songs []model.SongInfo
	err := model.DB().Find(&songs).Error
	if err != nil {
		return nil, err
	}
	return songs, nil
}
