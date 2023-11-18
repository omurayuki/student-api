/*----------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See LICENSE in the project root for license information.
 *---------------------------------------------------------------------------------------*/

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/omurayuki/student-api/backend/internal/handler"
)

func main() {
	router := gin.Default()
	setupRoutes(router)
	router.Run(":8000")
}

func setupRoutes(router *gin.Engine) {
	router.GET("/students", handler.GetStudents)
	router.POST("/students", handler.CreateStudent)
	router.PUT("/students", handler.UpdateStudent)
	router.DELETE("/students", handler.DeleteStudent)
}
