package person

type Person struct {
	Id    uint64 `gorm:"primaryKey;autoIncrement;not null"`
	Name  string `gorm:"not null"`
	Age   int    `gorm:"not null"`
	Email string `gorm:"not null;unique"`
	Phone string `gorm:"not null"`
	Grade string `gorm:"not null"`
}
