package models

type Task struct {
	Taskname string `JSON:"taskname" gorm:"not null"`
	Due      uint   `JSON:"due" gorm:"not null"`
}

type UpdateTask struct {
	Oldtask string `JSON:"oldtask" gorm:"not null"`
	Newtask string `JSON:"newtask" gorm:"not null"`
	Newdue  uint   `JSON:"newdue" gorm:"not null"`
}
