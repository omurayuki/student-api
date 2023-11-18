/*----------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See LICENSE in the project root for license information.
 *---------------------------------------------------------------------------------------*/

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/omurayuki/student-api/backend/internal/hello"
)

func main() {
	portNumber := "8000"
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, hello.Hello())
	})

	router.Run(":" + portNumber)
}
