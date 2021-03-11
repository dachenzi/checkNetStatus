package service

import (
	"fmt"
	"os/exec"
	"syscall"
)

func RestartKubeletService() (retCode int, err error) {
	// 创建命令对象
	cmd := exec.Command("/bin/bash", "-c", "ls -l")

	// 创建输出管道
	//stdout, err := cmd.StdoutPipe()
	//if err != nil {
	//	fmt.Println("创建管道异常：", err)
	//}
	//defer stdout.Close()

	// 执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("执行命令异常：", err)
		return 127, err
	}

	// 读取命令输出结果
	//retBytes, err := ioutil.ReadAll(stdout)
	//fmt.Println(string(retBytes))

	// 获取命令执行返回码
	if err := cmd.Wait(); err != nil {
		return 127, err
	}
	retCode = cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
	return retCode, nil

}
