package model
type Article struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Title string `json:"title"`
	Date string `json:"date"`
	Content string `json:"content"`
}
type ArticleList struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Name string `json:"name"`
}
type ArtCreateReq struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Name 	string  `json:"name"`
}
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Token struct {
    Token string `json:"token"`
}