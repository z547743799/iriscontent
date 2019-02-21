package service

import (
	"encoding/json"
	"fmt"
	"gitlab.com/z547743799/irismanager/db"
	"gitlab.com/z547743799/irismanager/models"

	"github.com/garyburd/redigo/redis"
	"github.com/go-xorm/xorm"
	"gitlab.com/z547743799/iriscontent/redisinit"
)

type ContentService interface {
	GetContentListByCid(id int64) []models.TbContent
}

type contentService struct {
	engine *xorm.Engine
}

func NewContentService() ContentService {
	return &contentService{
		engine: db.X,
	}
}

func (d *contentService) GetContentListByCid(id int64) []models.TbContent {

	pool := redisinit.Re.Get()
	defer pool.Close()
	enc, err := redis.Bytes(pool.Do("get", id))
	datalist := make([]models.TbContent, 0)
	if err != nil {
		return nil
	}
	if enc == nil {
		err := d.engine.Where("id=?", id).Desc("id").Find(&datalist)

		if err != nil {
			return nil
		}

		enc, _ := json.Marshal(datalist)
		_, err = pool.Do("set", id, enc)
		if err != nil {
			return nil
		}
		return datalist
	}
	err = json.Unmarshal(enc, datalist)
	if err != nil {
		fmt.Println(err)
	}
	return datalist
}
