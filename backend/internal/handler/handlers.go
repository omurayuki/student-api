package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omurayuki/student-api/backend/internal/student"
)

func GetStudents(c *gin.Context) {
	student := []student.Student{
		{ID: "12345", Name: "John Doe", Age: 20, Major: "Computer Science"},
		{ID: "12346", Name: "Jane Doe", Age: 22, Major: "Mathematics"},
	}

	c.JSON(http.StatusOK, student)
}

func CreateStudent(c *gin.Context) {
	// 新しい学生の登録
}

func UpdateStudent(c *gin.Context) {
	// 学生情報の更新
}

func DeleteStudent(c *gin.Context) {
	// 学生情報の削除
}


// // User user type
// type User struct {
// 	ID   int64    `json:"id,omitempty"`
// 	Name string   `json:"name,omitempty"`
// 	Addr *Address `json:"addr,omitempty"`
// }

// // Address address type
// type Address struct {
// 	City   string
// 	ZIP    int
// 	LatLng [2]float64
// }

// var alex = User{
// 	ID:   0,
// 	Name: "Yuki Omura",
// 	Addr: &Address{
// 		City: "",
// 		ZIP:  0,
// 		LatLng: [2]float64{
// 			0.0,
// 			0.0,
// 		},
// 	},
// }

// // Hello writes a welcome string
// func Hello() string {
// 	return "Hello, " + alex.Name
// }
