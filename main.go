package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// 创建HTTP/2.0客户端
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        10000,            // 最大空闲连接数
			MaxIdleConnsPerHost: 10000,            // 每个主机的最大空闲连接数
			IdleConnTimeout:     60 * time.Second, // 空闲连接的超时时间
		},
	}

	for i := 0; i < 2500; i++ {
		go func() {
			// 创建HTTP请求
			req, err := http.NewRequest("GET", "http://172.99.233.78/p5.jpg", nil)
			if err != nil {
				fmt.Printf("Error creating request: %v\n", err)
				return
			}

			// 发起并发请求
			for j := 0; j < 1000000; j++ {

				go func() {

					// 发送请求
					resp, err := client.Do(req)
					if err != nil {
						fmt.Printf("Error sending request: %v\n", err)
						return
					}
					//defer resp.Body.Close()

					// 处理响应（这里可以根据需要处理响应内容）
					// ...
					// 复制响应体到空设备（丢弃数据）
					_, err = io.Copy(io.Discard, resp.Body)
					if err != nil {
						fmt.Printf("Error copying response body: %v\n", err)
					}
					if j == 0 {
						//输出状态码
						fmt.Printf("Status Code: %d \n", resp.StatusCode)
					}

				}()
				time.Sleep(760 * time.Millisecond)
			}

		}()
		fmt.Printf("thread %d started\n", i)
		time.Sleep(20 * time.Millisecond)
	}

	// 等待所有请求完成

	time.Sleep(10 * time.Hour)
}
