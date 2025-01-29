package models

type Task struct {
	ID			int64		`redis:"id" json:"id"`
	Title		string		`redis:"title" json:"title"`
	Body		string		`redis:"body" json:"body"`
	Done		bool		`redis:"done" json:"done"`
}
