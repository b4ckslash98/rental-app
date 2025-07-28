package entity

type Car struct {
    ID         int     `gorm:"primaryKey"`
    CarName    string  `gorm:"size:50"`
    DayRate    float64
    MonthRate  float64
    Image      string  `gorm:"size:256"`
}