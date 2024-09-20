package tasks

import "github.com/trqvel/rest-to-do/db"

// Task представляет задачу
type Task struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	UserID      int    `json:"user_id"`
	Description string `json:"description"`
}

func GetTasksByUser(userID int) ([]Task, error) {
	var tasks []Task
	err := db.Conn.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func CreateTask(userID int, description string) error {
	task := Task{UserID: userID, Description: description}
	return db.Conn.Create(&task).Error
}
