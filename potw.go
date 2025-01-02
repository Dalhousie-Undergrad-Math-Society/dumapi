package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)


type Body struct {
	Name string `json:"name"`
	Answer string `json:"answer"`
}

func submitAnswer(c *gin.Context) {
	body := Body{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Println(body)

	fi, err := os.OpenFile("answers.csv", os.O_APPEND | os.O_WRONLY | os.O_CREATE, 0600)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	defer func() {
		if err := fi.Close(); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		} else {
			c.JSON(http.StatusAccepted, &body)
		}
	}()

	output := fmt.Sprintf("%s, %s, %s,\n", body.Name, body.Answer, time.Now().UTC())

	if _, err = fi.WriteString(output); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}



}
