package main

import (
	"gopkg.in/mgo.v2"
	"database/sql"
	"fmt"
	"strconv"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"bytes"
	"bufio"
	"strings"
)

func main() {

	file, err := os.Open("/home/siva/finalnew_resultplatform.csv")
	defer file.Close()

	if err != nil {
		println(err)
	}

	reader := bufio.NewReader(file)

	session, err := mgo.Dial("10.15.0.145")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("userlist").C("newuserdata")


	for {
			var buffer bytes.Buffer
			var line []byte
			line, _, err = reader.ReadLine()
			buffer.Write(line)
			println(buffer.String())
			// If we're just at the EOF, break
			if err != nil {
				 break
				} else {
                words := strings.Split(string(line),",")
				var userInfoImport = UserInfo{}
				var userdataImport = UserData{}
				userInfoImport.Flag=false
				userdataImport.Msisdn=words[4]
				userdataImport.UID=words[1]
				userdataImport.Token=words[5]
				userdataImport.PlatformUID=words[2]
				userdataImport.PlatformToken=words[3]
				userInfoImport.UserData=userdataImport
				c.Insert(userInfoImport)

			}
		}
}


type UserInfo struct {
	UserData UserData `json:"UserData"`
	Flag bool `json:"flag"`
}

type UserData struct {
	Msisdn string `json:"msisdn"`
	Token  string `json:"token"`
	UID    string `json:"uid"`
	PlatformUID string `json:"platformuid"`
	PlatformToken string `json:"platformtoken"`
}


func getDBConnection() *sql.DB{

	db, err := sql.Open("mysql", "platform:p1@tf0rmD1st@tcp(10.15.8.4:3306)/usersdb")
	if(err!=nil){
		fmt.Println(err)
	}
	return db
}

func ToNullString(s string) sql.NullString {
	return sql.NullString{String : s, Valid : s != ""}
}

func ToIntegerVal(i int64) string {
	var valueInt string
	valueInt = strconv.FormatInt(int64(i), 10)
	return valueInt
}

func ToStringFromInt(i int) string {
	var valueInt string
	valueInt = strconv.Itoa(i)
	return valueInt
}

func ToString(s sql.NullString) string {
	var valInString string
	if(s.Valid) {
		valInString = s.String
		//fmt.Println(valInString)
	} else {
		valInString = "NULL"
		//fmt.Println(valInString)
	}
	return valInString
}

type platformUser struct {
	CreateTime    time.Time  `json:"create_time"`
	HikeToken     string `json:"hike_token"`
	HikeUID       string `json:"hike_uid"`
	ID            int64    `json:"id"`
	Msisdn        string `json:"msisdn"`
	PlatformToken string `json:"platform_token"`
	PlatformUID   string `json:"platform_uid"`
	Status        sql.NullString `json:"status"`
	UpdateTs      time.Time `json:"update_ts"`
}