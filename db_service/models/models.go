package models

type Task struct {
	ID			int			`json:"id"`
	Title		string		`json:"title"`
	Body		string		`json:"body"`
	Done		bool		`json:"done"`
}