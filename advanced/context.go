package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work cancelled:", ctx.Err())
			return
		default:
			fmt.Println("Working")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	rootCtx := context.Background()
	ctx, cancel := context.WithTimeout(rootCtx, 2*time.Second)
	defer cancel()

	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	ctx = context.WithValue(ctx, "requestID", "D231FAWE342WQ")

	go doWork(ctx)

	time.Sleep(3 * time.Second)

	requestID := ctx.Value("requestID")
	if requestID != nil {
		fmt.Println("Request ID:", requestID)
	} else {
		fmt.Println("Request ID not found")
	}

	logWithContext(ctx, "Text")
}

func logWithContext(ctx context.Context, msg string) {
	requestIDVal := ctx.Value("requestID")
	log.Printf("Request ID: %v  - %v", requestIDVal, msg)
}

// func checkEvenOdd(ctx context.Context, num int) string {
// 	select {
// 	case <-ctx.Done():
// 		return "Operation canceled"
// 	default:
// 		if num%2 == 0 {
// 			return fmt.Sprintf("%d is even", num)
// 		} else {
// 			return fmt.Sprintf("%d is odd", num)
// 		}
// 	}
// }

// func main() {
// 	ctx := context.TODO()
// 	result := checkEvenOdd(ctx, 5)
// 	fmt.Println(result)
// 	ctx = context.Background()
// 	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
// 	defer cancel()
// 	result = checkEvenOdd(ctx, 10)
// 	fmt.Println(result)
// 	time.Sleep(2 * time.Second)
// 	result = checkEvenOdd(ctx, 15)
// 	fmt.Println(result)
// }

// func main() {
// 	todoContext := context.TODO()
// 	contextBkg := context.Background()
// 	ctx := context.WithValue(todoContext, "name", "John")
// 	fmt.Println(ctx)
// 	fmt.Println(ctx.Value("name"))
// 	ctx1 := context.WithValue(contextBkg, "city", "New York")
// 	fmt.Println(ctx1)
// 	fmt.Println(ctx1.Value("city"))
// }
