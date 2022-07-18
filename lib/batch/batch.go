package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	ch := make(chan int64, pool)
	res = []user{}
	for i := int64(0); i < n; i++ {
		wg.Add(1)
		ch <- i
		go func() {
			user := getOne(<-ch)
			mu.Lock()
			res = append(res, user)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	time.Sleep(time.Duration(n*100/pool*time.Millisecond.Nanoseconds() - time.Millisecond.Nanoseconds()*100))

	return res
}
