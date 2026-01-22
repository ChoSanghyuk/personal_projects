package main

import "macro/macro"

func main() {
	macro.ReenforceUntilTarget(17)
}

// // 매크로는 전체 화면 기준
// func main() {

// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go func(wg *sync.WaitGroup) {
// 		macro.HwaDamMac()
// 		wg.Done()
// 	}(&wg)

// 	openTime, err := time.ParseInLocation("2006-01-02 15:04", "2025-09-24 13:00", time.Local)
// 	if err != nil {
// 		return
// 	}
// 	var diff time.Duration

// 	for {
// 		now := time.Now()

// 		fmt.Println("====================================================================")
// 		fmt.Printf("Now time: %v\n", now)
// 		fmt.Printf("Open time: %v\n", openTime)
// 		fmt.Printf("Current Diff: %v\n", diff)
// 		fmt.Printf("Targer Refresh time: %v\n", openTime.Add(-1*diff))
// 		fmt.Println()

// 		if now.After(openTime.Add(-1 * diff)) { // todo. ms 단위로 빨리 새로고침 필요. 왕복 시간의 절반으로 해보자.
// 			fmt.Printf("Refresh time: %v\n", now)
// 			// r.KeyPress("f5") // 맥에서는 f5가 안 먹는다.
// 			r.KeyTap("r", "cmd")
// 			break
// 		} else if now.Second() >= 50 && now.Second() < (58-int(diff.Seconds())) {
// 			serverTime, err := getServerTime("https://content.yanolja.com/event/216")
// 			if err != nil {
// 				continue
// 			}
// 			diffTemp := serverTime.Sub(now)

// 			fmt.Printf("Server time: %v\n", serverTime)
// 			fmt.Printf("Local time:  %v\n", now)
// 			fmt.Printf("Diff:0  %v\n", serverTime.Sub(now))
// 			if diffTemp > 1*time.Second {
// 				diff = diffTemp
// 				println(diff) // 이정도 만큼 더 빨리 새로고침
// 			}

// 			if diffTemp < 2*time.Second { // 너무 자주 반복 방지
// 				time.Sleep(1 * time.Second)
// 			}
// 		} else if 45 <= now.Second() && now.Second() < 50 {
// 			time.Sleep(2 * time.Second)
// 		} else if now.Second() < 45 {
// 			time.Sleep(10 * time.Second)
// 		}
// 	}

// 	wg.Wait()
// }

// func getServerTime(url string) (time.Time, error) {
// 	resp, err := http.Head(url)
// 	if err != nil {
// 		return time.Time{}, err
// 	}
// 	defer resp.Body.Close()

// 	dateHeader := resp.Header.Get("Date")
// 	if dateHeader == "" {
// 		return time.Time{}, fmt.Errorf("no Date header found")
// 	}

// 	// Parse the HTTP date format
// 	serverTime, err := time.Parse(time.RFC1123, dateHeader)
// 	if err != nil {
// 		return time.Time{}, err
// 	}

// 	return serverTime, nil
// }
