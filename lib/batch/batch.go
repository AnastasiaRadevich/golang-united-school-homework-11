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
	parallel := int(pool)

	sem := make(chan struct{}, parallel)
	res = make([]user, n)

	for i := range res {
		wg.Add(1)
		sem <- struct{}{}
		go func(j int) {
			res[j] = getOne(int64(j))
			<-sem
			wg.Done()
		}(i)
	}
	wg.Wait()
	return res
}
