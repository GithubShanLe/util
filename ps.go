package util

import (
	"bytes"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

//减除脚本自身判断重复脚本执行 2018-10-29 le.shan
//start
//取代 "|" 不能在exec,Command()中使用的方法
func Pipline(cmds ...*exec.Cmd) ([]byte, []byte, error) {
	if len(cmds) < 1 {
		return nil, nil, nil
	}
	var output bytes.Buffer
	var stderr bytes.Buffer
	var err error
	maxindex := len(cmds) - 1
	cmds[maxindex].Stdout = &output
	cmds[maxindex].Stderr = &stderr

	for i, cmd := range cmds[:maxindex] {
		if i == maxindex {
			break
		}

		cmds[i+1].Stdin, err = cmd.StdoutPipe()
		if err != nil {
			return nil, nil, err
		}
	}

	// Start each command
	for _, cmd := range cmds {
		err := cmd.Start()
		if err != nil {
			return output.Bytes(), stderr.Bytes(), err
		}
	}

	// Wait for each command to complete
	for _, cmd := range cmds {
		err := cmd.Wait()
		if err != nil {
			return output.Bytes(), stderr.Bytes(), err
		}
	}

	return output.Bytes(), stderr.Bytes(), nil
}

//执行 ps 命令
func PsGrep(idIndex string) *int {
	ps := exec.Command("ps", "-ef")                  //查询所有的进程
	grep := exec.Command("grep", idIndex)            //查询grep字段
	deleteGrep := exec.Command("grep", "-v", "grep") //剔除grep自己
	count := exec.Command("wc", "-l")
	result, _, err := Pipline(ps, grep, deleteGrep, count)
	if err != nil {
		return nil
	}
	psCount := string(result)
	stringCount := strings.Split(psCount, "\n")
	intPsCOunt, err := strconv.Atoi(stringCount[0])
	if err != nil {
		log.Println(err)
		return nil
	}
	return &intPsCOunt

}

//这里面的代码其实是无用功
func KillProcess(key string) {
	pfs := PsGrep(key)
	if pfs != nil {
		log.Println(*pfs)
		if *pfs == 0 {
			//"do nothing"
		} else if *pfs > 0 {
			//"kill 脚本"
			cmd := exec.Command("sh", "kill.sh", key)
			cmd.Run()
		}

	} else {
		log.Println("ps grep error")
	}

}
