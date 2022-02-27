package EmailService

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/smtp"
	"os"
	"strings"
)

func Check(c *gin.Context)  {
	c.JSON(200,gin.H{"detail":"SERVICE RUNNING"})
}

var GlobalConfig *Config

type Mail struct {
	To []string	`form:"To"`
	Subject string	`form:"Subject"`
	Body string `form:"Subject"`
}

type Response struct{
	M Mail	`json:"email"`
	Detail string	`json:"detail"`
}

type Config struct{
	Email string	`json:"email"`
	Password string `json:"password"`
	Host string	`json:"host"`
	Port string	`json:"port"`
}



func ReadConfig() (*Config, error) {
	jsonFile, err := os.Open("config.json")
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
		return &Config{}, err
	}
	fmt.Println("Successfully Opened config.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var mconfig Config
	err = json.Unmarshal(byteValue, &mconfig)
	if err != nil {
		return &Config{}, err
	}
	return &mconfig,nil
}

func sendMail(mail *Mail) error{
	toList := mail.To
	msg := BuildMessage(mail)
	body := []byte(msg)
	auth := smtp.PlainAuth("", GlobalConfig.Email, GlobalConfig.Password, GlobalConfig.Host)
	err := smtp.SendMail(GlobalConfig.Host+":"+GlobalConfig.Port, auth, GlobalConfig.Email, toList, body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func BuildMessage(mail *Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", "nexcyb.noreply@nexcyb.com")
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)
	return msg
}

func SendMailHandler(c *gin.Context)  {
	var mMail Mail
	err := c.BindJSON(&mMail)
	if err != nil {
		c.JSON(500, gin.H{"detail": "malformed URL data"})
	}
	if sendMail(&mMail)!=nil{
		response := Response{mMail,"EMAIL FAILED TO SEND"}
		c.JSON(500,response)
	}
	response := Response{mMail,"EMAIL SENT"}
	c.JSON(200,response)
}
