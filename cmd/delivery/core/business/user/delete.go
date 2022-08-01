package user

// type Model struct {
// 	Name       string             `json:"name" gorm:"column:name;"`
// 	UserId     int                `json:"-" gorm:"column:owner_id;"`
// 	Addr       string             `json:"address" gorm:"column:addr;"`
// 	Logo       *common.Image      `json:"logo" gorm:"column:logo;"`
// 	Cover      *common.Images     `json:"cover" gorm:"column:cover;"`
// 	User       *common.SimpleUser `json:"user" gorm:"preload:false;"`
// 	LikedCount int                `json:"liked_count" gorm:"column:liked_count;"` // computed field
// }

// type SQLModel struct {
// 	Id        int        `json:"-" gorm:"column:id;"`
// 	FakeId    *UID       `json:"id" gorm:"-"`
// 	Status    int        `json:"status" gorm:"column:status;default:1;"`
// 	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"created_at"`
// 	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"updated_at"`
// }
