package model

import (
	"time"
	
)

func init (){
	govalidator.SetFieldsReqiredByDefault(value:true)
}

type Base struct {
	ID        string    `jason:"id" valid:"uuid"`
	CreatedAt time.Time `jason:"created_at" valid:"-"`
	UpdateAt  time.Time `jason:"updated_at" valid:"-"`
}
