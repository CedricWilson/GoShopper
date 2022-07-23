package models

type User struct {
	Name   string `json:"name"`
	Age    int `json:"age"`
	Height float32 `json:"height"`
}

func StaticUsers() []User {
return []User{{Name: "Cedric", Age: 2, Height: 3.4}, {Name: "Raj", Age: 2, Height: 3.4}, {Name: "Sagar", Age: 2, Height: 3.4}}
} 

