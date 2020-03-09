package bases

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type BaseEntity struct {
	Validator *validator.Validate
}

func (b *BaseEntity) NewBaseEntity() *BaseEntity{
	b.Validator = validator.New()
	return &BaseEntity{}
}
func (b *BaseEntity) Validate(input interface{}) {
	if err := b.Validator.Struct(input); err != nil {
		fmt.Print(err)
	}
}