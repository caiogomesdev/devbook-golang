package dto

type UserDto struct {
  ID uint `json:"id,omitempty"`
  Name string `json:"name"`
  Nick string `json:"nick"`
  Email string `json:"email"`
  Password string `json:"password,omitempty"`
}
