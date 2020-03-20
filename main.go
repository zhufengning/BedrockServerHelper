package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func printLalala() {
	rand1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	sentences := []string{"石剑是检验真理的唯一标准!", "你，就是我的Operator吗？", "呐呐呐呐呐呐", "服务器即将爆炸！", "这里是Server酱，老二刺螈了", "林地府邸一根棒，预备，起！", "不定期崩溃不是bug，是本服务器的特性！"}
	for true {
		fmt.Println("me 酱:", sentences[rand1.Int()%len(sentences)])
		time.Sleep(10 * time.Minute)
	}
}

func main() {
	//fmt.Println(time.Now().Hour())
	go printLalala()
	//var tmp string
	cin := bufio.NewReader(os.Stdin)
	for true {
		tmp, err := cin.ReadString('\n')
		if err == nil {
			fmt.Print(tmp)
		}
	}
}
