package main

import "fmt"

type Movie struct {
	Name         string
	Director     string
	Screenwriter string
	Genre        string
	Actor        string
	Language     string
	Showtime     string
	Duration     string
	OtherName    string
}

func main() {
	m := Movie{
		Name:         "西线无战事",
		Director:     "爱德华·贝尔格",
		Screenwriter: "伊恩·斯托克尔 / 爱德华·贝尔格 / Lesley Paterson / 埃里希·玛利亚·雷马克",
		Actor:        "费利克斯·卡米勒 / 阿尔布雷希特·舒赫 / 亚伦·希尔默 / 丹尼尔·布鲁赫 / 塞巴斯蒂安·胡克 / 安东·范·卢克 / 大卫·史崔梭德 / 埃丁·哈萨诺维奇 / 乔·温特劳布 / 吕克·费特 / Andreas Döhler / 米夏埃尔·维滕博恩 / Peter Sikorski / 安德烈·马尔孔 / Dominikus Weileder / Tobias Langhoff / 阿德里安·格鲁内瓦尔德 / 菲利克斯·冯·布雷多 / 迈克尔·斯坦格 / Moritz Klaus",
		Genre:        "剧情 / 动作 / 战争",
		Language:     "英语 / 德语",
		Showtime:     "2022-09-12(多伦多电影节) / 2022-09-29(德国) / 2022-10-28(德国网络)",
		Duration:     "148分钟",
		OtherName:    "新西线无战事 / All Quiet on the Western Front",
	}
	fmt.Printf("请输入你想要了解的信息\n1.电影名\n2.导演\n3.编剧\n4.电影类型\n5.演员\n6.电影语言\n7.上映时间\n8.电影时长\n9.电影别名\n10.退出\n")
	var option int
	for {
		fmt.Scanf("%d", &option)
		fmt.Scanln()
		if option == 1 {
			fmt.Println(m.Name)
		} else if option == 2 {
			fmt.Println(m.Director)
		} else if option == 3 {
			fmt.Println(m.Screenwriter)
		} else if option == 4 {
			fmt.Println(m.Genre)
		} else if option == 5 {
			fmt.Println(m.Actor)
		} else if option == 6 {
			fmt.Println(m.Language)
		} else if option == 7 {
			fmt.Println(m.Showtime)
		} else if option == 8 {
			fmt.Println(m.Duration)
		} else if option == 9 {
			fmt.Println(m.OtherName)
		} else if option == 10 {
			return
		} else {
			fmt.Println("请输入正确序号")
		}
	}
}
