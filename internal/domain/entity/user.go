package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	LEVEL_XP_INC = 100
)

type User struct {
	ID             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
	Name           string             `json:"name" bson:"name"`
	Email          string             `json:"email" bson:"email"`
	HashedPassword string             `json:"hashed_password" bson:"hashed_password"`
	XP             int                `json:"xp" bson:"xp"`
	XPToLevelUp    int                `json:"xp_to_level_up" bson:"xp_to_level_up"`
	Level          int                `json:"level" bson:"level"`
	Streak         int                `json:"streak" bson:"streak"`
	LastStreak     time.Time          `json:"last_streak" bson:"last_streak"`
}

func (user *User) SetDefault() *User {
	user.CreatedAt = time.Now()
	user.XP = 0
	user.XPToLevelUp = 100
	user.Level = 1
	user.Streak = 1
	user.LastStreak = user.CreatedAt
	return user
}

func (user *User) UpdateStreak() *User {
	t := time.Now()
	y, m, d := t.Date()
	ly, lm, ld := user.LastStreak.Date()
	if y == ly && m == lm && d == ld {
		return user
	}
	yy, ym, yd := t.AddDate(0, 0, -1).Date()
	if yy == ly && ym == lm && yd == ld {
		user.Streak++
	} else {
		user.Streak = 1
	}
	user.LastStreak = t
	return user
}

func (user *User) UpdateLevel() *User {
	if user.XP >= user.XPToLevelUp {
		user.Level++
		user.XP -= user.XPToLevelUp
		user.XPToLevelUp += LEVEL_XP_INC
	}
	return user
}
