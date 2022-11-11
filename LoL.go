package main

import (
	"fmt"
	"time"
)

type Fight interface {
	Skill()
	Attack()
	UseProp()
	Return()
}

type Legend struct {
	Name string
	Hp   int
	Mp   int
	Q    string
	W    string
	E    string
	R    string
}

var (
	y      = Legend{"亚索", 100, 100, "斩刚闪", "风之壁障", "踏前斩", "狂风绝息斩"}
	g      = Legend{"盖伦", 100, 100, "致命打击", "勇气", "审判", "德玛西亚正义"}
	emeny  = Legend{"敌方英雄", 100, 100, "Q技能", "W技能", "E技能", "R技能"}
	i      int
	j      int
	ticker = time.NewTicker(1 * time.Second)
	timer  = time.NewTimer(7 * time.Second)
)

func (l Legend) Skill() {
	if i == 1 {
		l = y
	} else {
		l = g
	}
	fmt.Println("你要释放的技能是\n1.Q技能\t2.W技能\t3.E技能\t4.R技能\t5.回城\t6.使用道具\t7.普攻")
	var x int
	for {
		if emeny.Hp <= 0 {
			fmt.Println("敌方死亡")
			fmt.Println("Victory")
			return
		}
		fmt.Scanf("%d", &x)
		fmt.Scanln()
		if l.Mp >= 0 {
			if x == 1 && l.Mp >= 5 {
				fmt.Println("释放Q技能", l.Q, "成功,Mp-5")
				fmt.Println("敌方Hp-5")
				emeny.Hp -= 5
				l.Mp -= 5
				fmt.Println("你现在的Mp为", l.Mp)
			} else if x == 2 && l.Mp >= 10 {
				fmt.Println("释放W技能", l.W, "成功,Mp-10")
				fmt.Println("敌方Hp-10")
				emeny.Hp -= 10
				l.Mp -= 10
				fmt.Println("你现在的Mp为", l.Mp)
			} else if x == 3 && l.Mp >= 10 {
				fmt.Println("释放E技能", l.E, "成功,Mp-10")
				fmt.Println("敌方Hp-10")
				emeny.Hp -= 10
				l.Mp -= 10
				fmt.Println("你现在的Mp为", l.Mp)
			} else if x == 4 && l.Mp >= 20 {
				fmt.Println("释放R技能", l.R, "成功,Mp-20")
				fmt.Println("敌方Hp-15")
				emeny.Hp -= 15
				l.Mp -= 20
				fmt.Println("你现在的Mp为", l.Mp)
			} else if x == 5 {
				l.Mp = 100
				Fight.Return(l)
			} else if x == 6 {
				Fight.UseProp(l)
				if j == 1 {
					l.Mp += 50
				} else if j == 2 {
					l.Hp += 20
				}
			} else if x == 7 {
				Fight.Attack(l)
			} else {
				fmt.Println("你的Mp不够，无法释放此技能，请回城或使用蓝瓶")
			}
		} else {
			fmt.Println("你的Mp不够，无法释放此技能，请回城或使用蓝瓶")
		}
	}
}

func (l Legend) Attack() {
	fmt.Println("你攻击了敌方英雄，敌方英雄血量-5")
	emeny.Hp -= 5
}

func (l Legend) UseProp() {
	fmt.Println("你要使用的道具是1.蓝瓶\t2.血瓶")
	fmt.Scanf("%d", &j)
	fmt.Scanln()
	if j == 1 {
		l.Mp += 50
	} else if j == 2 {
		l.Hp += 20
	}
}

func (l Legend) Return() {
	ticker = time.NewTicker(1 * time.Second)
	timer = time.NewTimer(7 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-timer.C:
			fmt.Println("回城成功")
			return
		case <-ticker.C:
			fmt.Println("回城中")
		}
	}
}

func Mylegend(m Legend) {
	fmt.Println("欢迎来到英雄联盟，你选择的英雄是" + m.Name)
	fmt.Println("你可以选择\n1.放技能\n2.攻击\n3.回城\n4.投降")
	for {
		var x int
		fmt.Scanf("%d", &x)
		fmt.Scanln()
		if x == 1 {
			Fight.Skill(m)
			return
		} else if x == 2 {
			Fight.Attack(m)
			if emeny.Hp <= 0 {
				fmt.Println("敌方死亡")
				fmt.Println("Victory")
				return
			}
		} else if x == 3 {
			Fight.Return(m)
		} else if x == 4 {
			fmt.Println("Defeat")
			return
		}
	}
}

func main() {
	var option int
	var x int
	fmt.Println("请输入你的命令\n1.获取英雄列表\n2.选择你的英雄\n3.退出游戏")
	for {
		fmt.Scanf("%d", &option)
		fmt.Scanln()
		if option == 1 {
			fmt.Println(y.Name, g.Name)
		} else if option == 2 {
			fmt.Println("请输入你想要选的英雄\n1.亚索\n2.盖伦")
			fmt.Scanf("%d", &i)
			fmt.Scanln()
			if i == 1 {
				Mylegend(y)
				fmt.Println("是否开启下一把1.是\t2.否")
				fmt.Scanf("%d", &x)
				fmt.Scanln()
				if x == 1 {
					continue
				} else {
					fmt.Println("退出游戏成功")
					return
				}
			} else if i == 2 {
				Mylegend(g)
				fmt.Println("是否开启下一把1.是\t2.否")
				fmt.Scanf("%d", &x)
				fmt.Scanln()
				if x == 1 {
					continue
				} else {
					fmt.Println("退出游戏成功")
					return
				}
			} else {
				fmt.Println("请输入正确的英雄序号")
			}
		} else if option == 3 {
			fmt.Println("退出游戏成功")
			return
		}
	}
}
