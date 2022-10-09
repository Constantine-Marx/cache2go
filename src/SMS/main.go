package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var (
	g_sceneFlag = 0
	g_Msg       = make([]studentMsg, 0, 10)
	g_cnt       = 0
)

type studentMsg struct {
	Name  string
	ID    int
	Class string
}

func getMessage(content string, temp interface{}) {
	fmt.Printf("%s", content)
	fmt.Scanln(temp)
}

func showMainscene() {
	fmt.Println(`
		欢迎进入学生管理系统！
		1、进入系统
		2、退出系统
	`)
}

func showFirscene() {
	fmt.Println(`
		1、增加学生信息
		2、删除学生信息
		3、修改学生信息
		4、查看学生信息
		0、返回上一级
	`)
}

func studentAdd() {

	file, err := os.OpenFile("D:\\Go\\work\\src\\SMS\\sms.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}

	defer file.Close()

	var tmp = studentMsg{}
	fmt.Println("请输入学生姓名:")
	fmt.Scanln(&tmp.Name)
	fmt.Println("请输入学生学号:")
	fmt.Scanln(&tmp.ID)
	fmt.Println("请输入学生班级:")
	fmt.Scanln(&tmp.Class)

	g_Msg = append(g_Msg, tmp)
	fmt.Println(len(g_Msg))
	// g_Msg[len(g_Msg)-1] = tmp
	studentAppend()
}

func studentUnappend(){
	file, err := os.OpenFile("D:\\Go\\work\\src\\SMS\\sms.txt", os.O_CREATE|os.O_RDONLY|os.O_APPEND, 0666)//将文本数据读入结构体数组

			if err != nil {
				fmt.Println("open file failed,err:", err)
				return
			}

			if err == io.EOF {
				fmt.Println("Read EOF")
				return
			}

			defer file.Close()

			rdr := bufio.NewReader(file)
			for {
				var tmp = make([]byte, 2048)
				tmp, _, err = rdr.ReadLine()
				if err == io.EOF {
					break
				}
				if err != nil {
					fmt.Println("ERROR:", err)
				}
				s2 := studentMsg{}
				err1 := json.Unmarshal(tmp, &s2)
				if err1 != nil {
					fmt.Println("ERROR:", err1)
				}
				g_Msg[g_cnt] = s2
				g_cnt++
			}
			g_cnt = 0
}

func studentAppend(){
	file,err := os.OpenFile("d:/Go/work/src/SMS/sms.txt",os.O_CREATE|os.O_WRONLY|os.O_TRUNC,0644)
	if err != nil {
		fmt.Println("student append func open file failed,Error:",err)
		return
	}

	defer file.Close()

	for _,v := range g_Msg{
		file_byte,err1 := json.Marshal(v)

		if err1 != nil {
			fmt.Println("student append func marshal failed,Error:",err1)
		}

		_,err1 = file.Write(file_byte)
		if err1 != nil {
			fmt.Println("student append func file write failed,Error:",err1)
		}
		file.WriteString("\n")
		fmt.Println("更新成功！")

	}
}

func studentRemove() {

}

func studentModify() {

}

func studentCheck() {
	for _, k := range g_Msg {
		fmt.Println(k)
	}
}



func main() {

	for {
		switch g_sceneFlag {
		case 0:
			showMainscene()
			temp := 0
			getMessage("请输入您的选项:", &temp)
			if temp == 1 { //显示学生名单
				g_sceneFlag = 1
			} else if temp == 2 { //退出系统
				return
			} else { //输入不符格式
				g_sceneFlag = 3
			}
		case 1: //show list

			studentUnappend()

			temp := 0
			showFirscene()
			getMessage("请输入您的选项:", &temp)//读取用户输入
			switch temp {
			case 1:
				studentAdd()
			case 2:
				studentRemove()
			case 3:
				studentModify()
			case 4:
				studentCheck()
			case 0:
				g_sceneFlag = 0
			}
		case 2:
			return
		case 3:
			temp := 0 //初始化temp
			fmt.Print("Your input is undeclared.Please input again:")
			fmt.Scanln(&temp)
			if temp == 1 { //显示学生名单
				g_sceneFlag = 1
			} else if temp == 2 { //退出系统
				return
			} else { //输入不符格式
				g_sceneFlag = 3
			}
		}
	}
}
