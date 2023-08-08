// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID              int32  `gorm:"column:id;primaryKey" json:"id"`
	Username        string `gorm:"column:username" json:"username"`
	Password        string `gorm:"column:password" json:"password"`
	FollowCount     int64  `gorm:"column:follow_count" json:"follow_count"`
	FollowerCount   int64  `gorm:"column:follower_count" json:"follower_count"`
	Avatar          string `gorm:"column:avatar" json:"avatar"`
	BackgroundImage string `gorm:"column:background_image" json:"background_image"`
	Signature       string `gorm:"column:signature" json:"signature"`
	TotalFavorited  int64  `gorm:"column:total_favorited" json:"total_favorited"`
	WorkCount       int64  `gorm:"column:work_count" json:"work_count"`
	FavoriteCount   int64  `gorm:"column:favorite_count" json:"favorite_count"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
