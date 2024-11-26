package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func startGoReplay() (*exec.Cmd, error) {
	// 设置 gor 命令和参数
	cmd := exec.Command("gor/linux/gor", "--input-raw", ":8888", "--output-stdout")

	// 设置命令的标准输出和标准错误输出（这一步不是必须的）
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 启动命令
	err := cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("failed to start GoReplay: %w", err)
	}

	// 返回启动的 Cmd 进程，方便后续停止
	return cmd, nil
}

func stopGoReplay(cmd *exec.Cmd) error {
	// 停止 gor 进程
	err := cmd.Process.Kill() // 通过 Kill 方法停止进程
	if err != nil {
		return fmt.Errorf("failed to stop GoReplay: %w", err)
	}
	return nil
}

func main() {
	// 启动 GoReplay
	cmd, err := startGoReplay()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// 等待 10 秒，模拟某些操作
	fmt.Println("GoReplay is running...")

	// 等待 10 秒钟
	time.Sleep(100 * time.Second)

	// 停止 GoReplay
	err = stopGoReplay(cmd)
	if err != nil {
		log.Fatalf("Error stopping GoReplay: %v", err)
	}

	fmt.Println("GoReplay has been stopped.")
}
