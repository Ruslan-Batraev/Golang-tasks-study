package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

type ShutdownFlag struct {
	flag int32
}

func (s *ShutdownFlag) Set() {
	atomic.StoreInt32(&s.flag, 1)
}

func (s *ShutdownFlag) IsSet() bool {
	return atomic.LoadInt32(&s.flag) != 0
}

func main() {
	var wg sync.WaitGroup
	shutdown := &ShutdownFlag{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				if shutdown.IsSet() {
					fmt.Printf("Рабочий %d завершает работу\n", id)
					return
				}
				fmt.Printf("Рабочий %d работает...\n", id)
				time.Sleep(500 * time.Millisecond)
			}
		}(i)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	fmt.Println("Получен сигнал завершения, устанавливаем флаг...")
	shutdown.Set()

	wg.Wait()
	fmt.Println("Все рабочие остановлены, выход.")
}
