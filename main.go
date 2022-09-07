package main

import ("fmt"
"net/http"
"html/template")

type User struct{
  name string
  age uint16 // целое неотризательное число
  money int16
  avg_grades, happiness float64
}

func (u User) getAllInfo() string {
  return fmt.Sprintf("User name is: %s. He is %d and he has :%d dollars",
    u.name, u.age, u.money)
}

func (u *User) setNewName(newName string) {
  u.name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
jack := User{"Jack",27, -50, 4.3, 0.7}
/*jack.setNewName("Jessica")
  fmt.Fprintf(w, jack.getAllInfo())*/
  tmpl, _ := template.ParseFiles("templates/home_page.html")
  tmpl.Execute(w, jack)
}

func about_page(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Moscow, Russia.")
}

func handleRequest() {
  http.HandleFunc("/", home_page)
  http.HandleFunc("/about/", about_page)
  http.ListenAndServe(":8080", nil)
}
func main() {
  //jack := User{name: "Jack", age: 27, money: -50, avg_grades: 4.3, happiness: 0.7}

handleRequest()

}
