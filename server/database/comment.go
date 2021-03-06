package database

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Text    string
	TopicID uint
	BuyerID uint
	Date    string
}

type MoreComment struct {
	Comment
	Buyer
}

func CreateComment(comment Comment) error {
	return gormDB.Create(&comment).Error
}

func GetComments(id uint) ([]MoreComment, error) {
	comments := []MoreComment{}
	err := gormDB.Raw(
		"SELECT * FROM comments c INNER JOIN buyers b ON c.buyer_id = b.id WHERE c.topic_id=? AND c.deleted_at IS NULL",
		id).Scan(&comments).Error
	return comments, err
}

func DeleteComment(id, buyerID uint) error {
	return gormDB.Where("id = ? AND buyer_id = ?", id, buyerID).Delete(&Comment{}).Error
}

func DeleteCommentByTopic(id uint) error {
	return gormDB.Where("topic_id = ?", id).Delete(&Comment{}).Error
}
