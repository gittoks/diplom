package database

import (
	"github.com/jinzhu/gorm"
)

type Topic struct {
	gorm.Model
	Name        string
	Description string
	BuyerID     uint
}

func CreateTopic(topic Topic) error {
	return gormDB.Create(&topic).Error
}

func GetTopics() ([]Topic, error) {
	topics := []Topic{}
	err := gormDB.Model(&Topic{}).Find(&topics).Error
	return topics, err
}

func GetTopic(id uint) (Topic, error) {
	topic := Topic{}
	err := gormDB.Where("id = ?", id).First(&topic).Error
	return topic, err
}

func DeleteTopic(id uint) error {
	return gormDB.Where("id = ?", id).Delete(&Topic{}).Error
}
