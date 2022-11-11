package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sum := 0
	var maxNum int
	fmt.Println("请输入你要猜的范围")
	fmt.Scanf("%d", &maxNum)
	fmt.Scanln()
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)

	fmt.Println("可以输入你猜的数字了")
	for {
		var guess int
		_, err := fmt.Scanf("%d\n", &guess)
		if err != nil {
			fmt.Println("猜的是一个整数，不要乱输入别的东西")
			continue
		}
		fmt.Println("你猜的数字是", guess)
		if guess > secretNumber {
			fmt.Print("你的数字有点大了，")
			if guess-secretNumber <= 5 {
				fmt.Println("现在可不是放弃的时候，你已经非常接近正确答案了！")
				sum++
			} else if guess-secretNumber <= 20 {
				fmt.Println("但胜利近在眼前，时刻准备开香槟吧！")
				sum++
			} else if guess-secretNumber <= 35 {
				fmt.Println("离正确答案还有一段距离，加油吧！")
				sum++
			} else if guess-secretNumber <= 60 {
				fmt.Println("你在岸的这边，答案在海的那边。")
				sum++
			} else {
				fmt.Println("而且大的有点离谱！！！")
				sum++
			}
		} else if guess < secretNumber {
			fmt.Print("你的数字有点小了，")
			if secretNumber-guess <= 5 {
				fmt.Println("现在可不是放弃的时候，你已经非常接近正确答案了！")
				sum++
			} else if secretNumber-guess <= 20 {
				fmt.Println("但胜利近在眼前，时刻准备开香槟吧！")
				sum++
			} else if secretNumber-guess <= 35 {
				fmt.Println("离正确答案还有一段距离，加油吧！")
				sum++
			} else if secretNumber-guess <= 60 {
				fmt.Println("你在岸的这边，答案在海的那边。")
				sum++
			} else {
				fmt.Println("而且小的有点离谱！！！")
				sum++
			}
		} else {
			sum++
			if sum <= 5 {
				fmt.Println("你猜对辣！居然只用了", sum, "次就成功辣！！我愿称你为重邮猜数小能手，为你鼓掌！！！")
			} else if sum <= 10 {
				fmt.Println("恭喜你你猜对了！只花了", sum, "次，看来你有在猜数上的天赋，想不想要再来一把紧张刺激的猜数游戏呢。")
			} else if sum <= 15 {
				fmt.Println("正确！你可是经历了一场大战呢，", sum, "次可是个不小的数目啊，再去多练习练习吧。")
			} else if sum <= 20 {
				fmt.Println("终于对了，虽然花了", sum, "次，但能坚持这么久真是不错，恭喜你。")
			} else if sum <= 40 {
				fmt.Println("猜了", sum, "次，可算是对了，累了吧，赶快去休息吧！")
			} else {
				fmt.Println("nb！居然能猜到", sum, "次，想不到你玩一个猜数游戏能如此坚持，重邮强者，恐怖如斯！")
			}
			return
		}
	}
}
