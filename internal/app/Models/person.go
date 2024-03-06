package models

type Person struct {
	Id       *uint   `json:"id"`
	Name     *string `json:"name"`
	LastName *string `json:"lastName"`
	Email    *string `json:"email"`
	Age      *uint   `json:"age"`
}

func (Person) TableName() string {
	return "public.person"
}
