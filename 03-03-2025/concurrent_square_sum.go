package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

func squareWorker(a int, sharedCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Square worker is working")
	num := a * a
	fmt.Println("Value stored in the shared channel", num)
	sharedCh <- num
}

func aggregateSquare(sharedCh chan int, maxLength int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Aggregate Square is working")
	count := 0
	sum := 0

	for count < maxLength {
		val := <-sharedCh
		time.Sleep(time.Second)
		fmt.Println("Values recieved from the shared channel: ", val)
		sum += val
		count++
	}
	fmt.Println("=======================================")
	fmt.Println("The sum of the square is: ", sum)
	fmt.Println("=======================================")
}

func getUserInput() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	number, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalln("Error converting the string to int", err)
	}
	return number
}

func main() {
	fmt.Println("Enter the length of the list: ")
	length := getUserInput()
	squareList := make([]int, length)

	fmt.Println("Enter the values you want to add in the squareList: ")
	for i := 0; i < length; i++ {
		squareList[i] = getUserInput()
	}
	fmt.Println("=======================================")
	var wg sync.WaitGroup
	result := make(chan int, len(squareList))
	wg.Add(len(squareList))

	for _, num := range squareList {
		go squareWorker(num, result, &wg)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	wg.Add(1)
	go aggregateSquare(result, len(squareList), &wg)
	wg.Wait()
	fmt.Println("All function has executed!")
}
