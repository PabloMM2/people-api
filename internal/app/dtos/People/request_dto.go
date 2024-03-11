package peopleDto

type PeopleCreateRquest struct {
	Name     *string  `json:"name" validate:"required"`
	LastName *string  `json:"lastName" validate:"required"`
	Email    *string  `json:"email" validate:"required"`
	Age      *uint    `json:"age" validate:"required"`
	Password *string  `json:"password" validate:"required"`
	Amount   *float64 `json:"amount"`
}
