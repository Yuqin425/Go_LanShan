package dao

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var (
	users = map[string]string{}
	tels  = map[string]string{}
)

func AddUser(username, password string) {
	users[username] = password
	filePath := "f:/demo/user.txt"

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_SYNC, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	write.WriteString(username + "\n")
	write.WriteString(password + "\n")
	write.Flush()
}

func AddTel(username, telephone string) {
	tels[username] = telephone
	filePath := "f:/demo/tel.txt"

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_SYNC, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	write.WriteString(username + "\n")
	write.WriteString(telephone + "\n")
	write.Flush()
}

func SelectUser(username string) bool {
	var (
		user string
		pwd  string
		tel  string
		line string
	)
	filePath1 := "f:/demo/user.txt"
	filePath2 := "f:/demo/tel.txt"
	file1, err := os.OpenFile(filePath1, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	file2, err := os.OpenFile(filePath2, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	reader1 := bufio.NewReader(file1)
	for i := 1; ; i++ {
		line, err = reader1.ReadString('\n')

		if err != nil && err != io.EOF {
			fmt.Println(err)
		}
		if i%2 == 1 {
			user = line
			user = strings.Replace(user, "\n", "", -1)
		} else if i%2 == 0 {
			pwd = line
			pwd = strings.Replace(pwd, "\n", "", -1)
			users[user] = pwd
		}

		if err == io.EOF {
			break
		}
	}
	reader2 := bufio.NewReader(file2)
	for i := 1; ; i++ {
		line, err = reader2.ReadString('\n')

		if err != nil && err != io.EOF {
			fmt.Println(err)
		}
		if i%2 == 1 {
			user = line
			user = strings.Replace(user, "\n", "", -1)
		} else if i%2 == 0 {
			tel = line
			tel = strings.Replace(tel, "\n", "", -1)
			tels[user] = tel
		}

		if err == io.EOF {
			break
		}
	}
	if users[username] == "" {
		return false
	}
	return true
}

func SelectPasswordFromUsername(username string) string {
	return users[username]
}

func SelectTelephoneFromUsername(username string) string {
	return tels[username]
}

func DeleteUser(username string) {
	var (
		user string
		pwd  string
	)

	filePath1 := "f:/demo/delete.txt"
	filePath2 := "f:/demo/user.txt"
	filePath3 := "f:/demo/null.txt"
	file1, err1 := os.OpenFile(filePath1, os.O_RDWR|os.O_APPEND|os.O_SYNC, 0666)
	file2, err2 := os.OpenFile(filePath2, os.O_RDWR|os.O_APPEND|os.O_SYNC, 0666)
	file3, err3 := os.OpenFile(filePath2, os.O_RDWR|os.O_APPEND|os.O_SYNC, 0666)

	if err1 != nil && err2 != nil && err3 != nil {
		fmt.Println("文件打开失败", err1, err2)
	}
	defer file1.Close()
	defer file2.Close()
	defer file3.Close()

	reader := bufio.NewReader(file2)
	write := bufio.NewWriter(file1)
	for i := 1; ; i++ {
		line, err := reader.ReadString('\n')

		if err != nil && err != io.EOF {
			fmt.Println(err)
		}
		if i%2 == 1 {
			user = line
			user = strings.Replace(user, "\n", "", -1)
			if strings.EqualFold(user, username) {
				continue
			}
			write.WriteString(user + "\n")
			write.Flush()
		} else if i%2 == 0 {
			pwd = line
			if strings.EqualFold(user, username) {
				continue
			}
			write.WriteString(pwd + "\n")
			write.Flush()
		}
		if err == io.EOF {
			break
		}
	}

	data1, err4 := ioutil.ReadFile(filePath1)
	if err4 != nil {
		fmt.Printf("文件打开失败=%v\n", err4)
		return
	}
	err4 = ioutil.WriteFile(filePath2, data1, 0666)
	if err4 != nil {
		fmt.Printf("文件打开失败=%v\n", err4)
	}

	data2, err5 := ioutil.ReadFile(filePath3)
	if err5 != nil {
		fmt.Printf("文件打开失败=%v\n", err5)
		return
	}
	err5 = ioutil.WriteFile(filePath1, data2, 0666)
	if err5 != nil {
		fmt.Printf("文件打开失败=%v\n", err5)
	}
}

func DeleteTel(username string) {
	var (
		user string
		pwd  string
	)

	filePath1 := "f:/demo/delete.txt"
	filePath2 := "f:/demo/tel.txt"
	filePath3 := "f:/demo/null.txt"
	file1, err1 := os.OpenFile(filePath1, os.O_RDWR|os.O_APPEND|os.O_SYNC, 0666)
	file2, err2 := os.OpenFile(filePath2, os.O_RDWR|os.O_APPEND|os.O_SYNC, 0666)
	file3, err3 := os.OpenFile(filePath2, os.O_RDWR|os.O_APPEND|os.O_SYNC, 0666)

	if err1 != nil && err2 != nil && err3 != nil {
		fmt.Println("文件打开失败", err1, err2)
	}
	defer file1.Close()
	defer file2.Close()
	defer file3.Close()

	reader := bufio.NewReader(file2)
	write := bufio.NewWriter(file1)
	for i := 1; ; i++ {
		line, err := reader.ReadString('\n')

		if err != nil && err != io.EOF {
			fmt.Println(err)
		}
		if i%2 == 1 {
			user = line
			user = strings.Replace(user, "\n", "", -1)
			if strings.EqualFold(user, username) {
				continue
			}
			write.WriteString(user + "\n")
			write.Flush()
		} else if i%2 == 0 {
			pwd = line
			if strings.EqualFold(user, username) {
				continue
			}
			write.WriteString(pwd + "\n")
			write.Flush()
		}
		if err == io.EOF {
			break
		}
	}

	data1, err4 := ioutil.ReadFile(filePath1)
	if err4 != nil {
		fmt.Printf("文件打开失败=%v\n", err4)
		return
	}
	err4 = ioutil.WriteFile(filePath2, data1, 0666)
	if err4 != nil {
		fmt.Printf("文件打开失败=%v\n", err4)
	}

	data2, err5 := ioutil.ReadFile(filePath3)
	if err5 != nil {
		fmt.Printf("文件打开失败=%v\n", err5)
		return
	}
	err5 = ioutil.WriteFile(filePath1, data2, 0666)
	if err5 != nil {
		fmt.Printf("文件打开失败=%v\n", err5)
	}
}

func AddComment(comment string) {
	filePath := "f:/demo/comment.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_SYNC, 0666)

	if err != nil {
		fmt.Println("文件打开失败", err)
	}

	defer file.Close()

	write := bufio.NewWriter(file)
	write.WriteString(comment + "\n")
	write.Flush()
}
