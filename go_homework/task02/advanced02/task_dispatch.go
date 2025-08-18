package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/

// Task 表示一个任务，包含名称和执行函数
type Task struct {
	Name string
	Fn   func() error // 返回 error 便于错误处理
}

// TaskResult 表示任务执行结果和耗时
type TaskResult struct {
	TaskName string
	Duration time.Duration
	Success  bool
	Error    error
}

// TaskScheduler 任务调度器
type TaskScheduler struct {
	tasks []Task
}

// NewTaskScheduler 创建一个新的调度器
func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		tasks: make([]Task, 0),
	}
}

// AddTask 添加一个任务
func (s *TaskScheduler) AddTask(name string, fn func() error) {
	s.tasks = append(s.tasks, Task{
		Name: name,
		Fn:   fn,
	})
}

// Run 并发执行所有任务，并返回结果
func (s *TaskScheduler) Run() []TaskResult {
	//TODO 等待一组协程执行完成
	var wg sync.WaitGroup
	resultChan := make(chan TaskResult, len(s.tasks))

	// 启动每个任务的协程
	for _, task := range s.tasks {
		//TODO 每启动一个任务增加 1 必须在启动任务的协程之前
		wg.Add(1)
		// 启动任务的协程
		go func(t Task) {
			//TODO 这个任务执行完成后，等待组减 1 defer 必须在任务执行完成后
			defer wg.Done()

			start := time.Now()
			err := t.Fn()
			duration := time.Since(start)

			result := TaskResult{
				TaskName: t.Name,
				Duration: duration,
				Success:  err == nil,
				Error:    err,
			}

			resultChan <- result
		}(task)
	}

	// 等待所有任务完成
	go func() {
		// wg.Wait() —— （主协程）要等你们全部完成 关闭通道
		wg.Wait()
		close(resultChan)
	}()

	// 收集结果返回
	var results []TaskResult
	for result := range resultChan {
		results = append(results, result)
	}

	return results
}

// 示例任务函数
func exampleTask(name string, workTime time.Duration, mayFail bool) func() error {
	return func() error {
		fmt.Printf("任务 [%s] 开始执行...\n", name)
		time.Sleep(workTime)

		// 模拟可能失败的任务
		if mayFail && rand.Intn(10) < 3 {
			return fmt.Errorf("任务 [%s] 执行失败", name)
		}

		fmt.Printf("任务 [%s] 执行完成，耗时 %v\n", name, workTime)
		return nil
	}
}

func main() {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 创建调度器
	scheduler := NewTaskScheduler()

	// 添加一些示例任务
	scheduler.AddTask("下载文件", exampleTask("下载文件", 2*time.Second, false))
	scheduler.AddTask("处理数据", exampleTask("处理数据", 1*time.Second, false))
	scheduler.AddTask("发送邮件", exampleTask("发送邮件", 1500*time.Millisecond, true))
	scheduler.AddTask("备份数据库", exampleTask("备份数据库", 800*time.Millisecond, false))
	scheduler.AddTask("生成报告", exampleTask("生成报告", 2500*time.Millisecond, false))

	// 并发执行所有任务
	fmt.Println("开始执行任务...")
	startTime := time.Now()

	results := scheduler.Run()

	totalDuration := time.Since(startTime)

	// 输出统计结果
	fmt.Println("\n=========== 执行结果统计 ===========")
	for _, r := range results {
		status := "成功"
		if !r.Success {
			status = "失败: " + r.Error.Error()
		}
		fmt.Printf("任务: %-12s | 状态: %-20s | 耗时: %v\n", r.TaskName, status, r.Duration)
	}
	fmt.Printf("\n所有任务执行完毕，总耗时: %v\n", totalDuration)
}
