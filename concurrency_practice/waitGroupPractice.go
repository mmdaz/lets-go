package concurrency_practice

import (
	"fmt"
	"lets_go/interface_practice/modules"
	"sync"
	"time"
)

type PracticeWaitGroups struct {
	interfacePractice modules.MainModule
}

func NewWaitGroupPractice(interfacePractice modules.MainModule) *PracticeWaitGroups {
	return &PracticeWaitGroups{interfacePractice: interfacePractice}
}

func (p *PracticeWaitGroups) RunPracticeWaitGroups(users []int, finished chan bool) {

	go func() {
		tokensChan := make(chan []string)
		usersChan := make(chan int)

		n := len(users)
		if n == 0 {
			n = 10
		}
		wg := sync.WaitGroup{}
		wg.Add(n)

		tokensWg := sync.WaitGroup{}
		tokensWg.Add(n)

		for i := 0; i < n; i++ {
			go func(tokensChan chan []string, tokensWg *sync.WaitGroup) {
				for tokens := range tokensChan {
					fmt.Println("Sending to firebase", tokens)
				}
				tokensWg.Done()
			}(tokensChan, &tokensWg)
		}

		for i := 0; i < n; i++ {
			go func(usersChan chan int, tokensChan chan []string, wg *sync.WaitGroup) {
				for u := range usersChan {
					start := time.Now()
					fmt.Print("u in userChan: ", u)
					firebaseApples := []string{"sdfsdfdfs", "fdsfsdfsdf", "fdsfsdfsdf"}
					p.interfacePractice.RunMainModule(i)
					fmt.Println("Get postgres auths took %s", time.Since(start))
					var firebaseTokens []string
					for _, firebase := range firebaseApples {
						firebaseTokens = append(firebaseTokens, firebase)
					}
					tokensChan <- firebaseTokens
				}
				wg.Done()
			}(usersChan, tokensChan, &wg)
		}

		for _, u := range users {
			usersChan <- u
		}
		close(usersChan)
		wg.Wait()
		close(tokensChan)
		tokensWg.Wait()
		finished <- true
	}()

}
