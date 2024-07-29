package models

import (
	"fmt"
)

type Model interface {
	describe() string
}

type Article struct {
	Uuid      string
	Title     string
	Content   string
	CreatedOn string
	UdpatedOn string
	Category  Category
}

type Category struct {
	Uuid string
	Name string
}

func (a Article) describe() string {
	return fmt.Sprintf(
		"Title: %s \nCategory: %s\n",
		a.Title,
		a.Category.Name,
	)
}

func (c Category) describe() string {
	return fmt.Sprintf("The name: %s", c.Name)
}
