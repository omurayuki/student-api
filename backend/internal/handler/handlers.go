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
