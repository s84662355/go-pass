package service

import (
	_ "GoPass/lib/redis"
	"GoPass/logic/model"
	"fmt"
	"strconv"
)

type Article struct {
}

func (a Article) UpdateReadAmount(data interface{}) {
	val, err := redisClient.Get(fmt.Sprintf("ReadAmount_%v", data)).Result()
	if err == nil {
		fmt.Println("阅读数量id: ", data, "  ", val)
		ReadCount, err := strconv.Atoi(val)
		if err == nil && ReadCount > 20 {
			article := model.Article{}.Get(data)
			if article.Id != 0 {
				ReadCount, _ := strconv.Atoi(redisClient.GetSet(fmt.Sprintf("ReadAmount_%v", data), 0).Val())
				article.UpdateReadAmount(ReadCount)
			}
		}

	}

}
