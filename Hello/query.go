package main

import(
	"fmt"
    "net/http"
    "encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

//返回结果结构体
type ResultUser struct {
	UserId int `json:"userId"`
	UserName string `json:"userName"`
	UserPassword string `json:"userPassword`
}

// 请求URL
const (
	COUNT_URL_QUERYBYID ="http://localhost:9090/queryById"
)


func main() {
	// 根据id查询
	uResult := queryById(4)
	if uResult.UserId == 4{
		fmt.Println("根据id查询用户==>",uResult)
	}
}

//根据id查询
func queryById(id int) ResultUser {
	contentType := "application/json;charset=utf-8"
    userParam, errs := json.Marshal(id) //转换成JSON返回的是byte[]
    if errs != nil {
        fmt.Println(errs.Error())
    }
    //fmt.Println(string(userParam))

     //发送请求
    req, err := http.NewRequest("POST", COUNT_URL_QUERYBYID, strings.NewReader(string(userParam)))
    if err != nil {
        panic(err)
    }
    req.Header.Set("Content-Type", contentType)
    client := &http.Client{}
    resp, err := client.Do(req)
	if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    //响应
    response, err := ioutil.ReadAll(resp.Body)
    //fmt.Println("response:", string(response))
    if err != nil {
        log.Println("Read failed:", err)
        return ResultUser{}
    }
    //log.Println("response:", string(response))
 
    //返回结果
    resultUser := ResultUser{}
    json.Unmarshal([]byte(string(response)), &resultUser) //json解析到结构体里面
 
    return resultUser
 
}