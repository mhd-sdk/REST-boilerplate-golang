package models

type Todo struct {
	ID          int    `gorm:"primaryKey" json:"id" example:"1"`
	Title       string `binding:"required" json:"title" example:"My first todo"`
	Description string `binding:"required" json:"description" example:"This is my first todo"`
	IsDone      bool   `default:"false" json:"isDone" example:"false"`
}

type todoResponse struct {
	Message string `json:"message" example:"Todo created successfully"`
	Todo    Todo   `json:"todo"`
}

type TodoPostBody struct {
	Title       string `binding:"required" json:"title" example:"My first todo"`
	Description string `binding:"required" json:"description" example:"This is my first todo"`
	IsDone      bool   `default:"false" json:"isDone" example:"false"`
}
