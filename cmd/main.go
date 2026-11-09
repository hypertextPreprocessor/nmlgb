package main
import (
	"time"
	"log"
	"sync"
	"os"
	"github.com/robfig/cron/v3"
)
//const Count = iota
var parllel chan int = make(chan int,20)

func main(){
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("[cron] :")
	log.SetOutput(os.Stdout)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	c := cron.New(cron.WithSeconds(),cron.WithLocation(loc))
	c.AddFunc("* * * * * *",func(){
		Task()
	})
	c.Start()
	select{}
}
func Task(){
	var wg sync.WaitGroup
	for range 3 {
		wg.Go(func(){
			Init()	
		})
		
		//go Init() //竞争抢占parllel
	}
	//往通道里发送100个任务
	for task := range 100 {
		parllel <- task
	}
	
	//close(parllel)
	//wg.Wait() //blocks until the WaitGroup task counter is zero.
}
func Init(){
	
	//loc, _ := time.LoadLocation("Asia/Shanghai")
	for r := range parllel {
		MyTask()
		log.Printf("[Worker %d] 收到并处理任务: %d", r, r)
	}
}
func test(){
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	
	parllel := make(chan int, 100) // 建议带缓冲，或者用无缓冲
	var wg sync.WaitGroup

	// 1. 启动固定数量（例如 3 个）的 worker 协程去消费任务
	workerCount := 3
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for task := range parllel {
				log.Printf("worker %d 正在处理: %d", workerID, task)
			}
		}(i)
	}

	// 2. 生产任务并写入 channel
	for p := 0; p < 100; p++ {
		parllel <- p
	}
	close(parllel) // 关闭 channel，通知 worker 任务已经发完

	// 3. 阻塞主协程，等待所有 worker 协程处理完毕
	wg.Wait()
	log.Println("所有任务处理完成")
}