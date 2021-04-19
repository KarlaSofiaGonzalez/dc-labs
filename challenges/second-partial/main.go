package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"github.com/gin-gonic/gin"
	"encoding/base64"
	"strings"
	"time"
	"math/rand"
)

const characters = "123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var actualuser = ""
//var actualpassword = "pw123"
var actualtoken = ""
var usernames []string
var passwords []string

func GenerateToken(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = characters[rand.Intn(len(characters))]
    }
    return string(b)
}

func ValidateUsername(actualU string) bool {
    for i := 0; i < len(usernames); i++ {
		if actualU == usernames[i] {
			return true
		}
	}
	return false
}

func ValidatePassword(actualU, actualP string) bool {
	pos := 0
    for i := 0; i < len(usernames); i++ {
		if actualU == usernames[i] {
			pos = i
		}
	}
	if actualP == passwords[pos]{
		return true
	} else {
		return false
	}
}

func main() {
	r := gin.Default()

	r.GET("/signin", func(c *gin.Context) {
		authorization := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
		todecod, _ := base64.StdEncoding.DecodeString(authorization[1])
		userdata := strings.SplitN(string(todecod), ":", 2)

		if !ValidateUsername(userdata[0]){
			usernames = append(usernames, userdata[0])	
			passwords = append(passwords, userdata[1])
			c.JSON(http.StatusOK, gin.H{
				"message": "Hi " + userdata[0] + " , your user has been created.",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Error, your username already exists.",
			})
		}

	})
	
	r.GET("/login", func(c *gin.Context) {
		authorization := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
		todecod, _ := base64.StdEncoding.DecodeString(authorization[1])
		userdata := strings.SplitN(string(todecod), ":", 2)
		
		//Verify the username and password

		if ValidateUsername(userdata[0]) {
			if ValidatePassword(userdata[0], userdata[1]) {
				actualuser = userdata[0]
				tokenrand := GenerateToken(8)
				actualtoken = tokenrand

				c.JSON(http.StatusOK, gin.H{
					"message": "Hi " + actualuser + " , welcome to the  DPIP System",
					"token": tokenrand,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "The password is incorrect",
				})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "The username is not registered",
			})
		}
	})

	r.GET("/logout", func(c *gin.Context) {
		authorization := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
		token := authorization[1]

		if token == actualtoken {
			
			actualtoken = ""
			c.JSON(http.StatusOK, gin.H{
				"message": "Bye " + actualuser + ", your token has been revoked",
			})
			actualuser = ""
		}
	})
	
	r.POST("/upload", func(c *gin.Context) {
		authorization := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
		token := authorization[1]

		if token == actualtoken {
			file, err := c.FormFile("data")
			if err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
				return
			}

			fileName := filepath.Base(file.Filename)
			f, err := os.Open(fileName)

			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"message": "Error uploading image",
					"filename": fileName,
				})
				//return
			}else{
				fi, _ := f.Stat()
				c.JSON(http.StatusOK, gin.H{
					"message": "An image has been successfully uploaded",
					"filename": fileName,
					"size": strconv.Itoa(int(fi.Size()/1000)) + " kb",
				})
			}

			f.Close()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "ERROR, you have to log in",
			})
		}

	})
	
	r.GET("/status", func(c *gin.Context) {
		authorization := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
		token := authorization[1]

		if token == actualtoken {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hi " + actualuser + " , the DPIP System is Up and Running",
				"time": time.Now().Format("2006-01-02T15:04:05+07:00"),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "ERROR, you have to log in",
			})
		}
		
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}