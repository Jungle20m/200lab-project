package common

import "time"

type SQLModel struct {
	ID         int        `json:"-" db:"id"`
	FakeId     *UID       `json:"id" db:"-"`
	Status     int        `json:"status" db:"status"`
	CreateTime *time.Time `json:"create_time" gorm:"create_time"`
	UpdateTime *time.Time `json:"update_time" gorm:"update_time"`
}

func (m *SQLModel) GenUID(dbType int) {
	uid := NewUID(uint32(m.ID), dbType, 1)
	m.FakeId = &uid
}
