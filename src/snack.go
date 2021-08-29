package main

import (
	"fmt"
	"math/rand"
	"os"
	"snack/xiaoyu1"
	"time"
)

const WIDE int = 20
const HIGH int = 20

var food Food

type Position struct{
	X int
	Y int
}

type Snack struct{
	size int  //长度
	dir byte   //方向，上下左右'U''D''L''R'
	pos [WIDE * HIGH]Position   //坐标
}

type Food struct{
	Position
}
func RandomFood(){
	food.X = rand.Intn(WIDE) +5
	food.Y = rand.Intn(HIGH)
	ShowUI(food.X,food.Y,'s')
}
//初始化蛇
func(s *Snack) InitSnack(){
	s.size = 2
	s.dir = 'R'
	s.pos[0].X = WIDE/2
	s.pos[0].Y = HIGH/2
	s.pos[1].X = WIDE/2-1
	s.pos[1].Y = HIGH/2
	//绘制蛇
	for i:=0;i<s.size;i++{
		var ui byte
		if i == 0{
			ui = '@'
		}else{
			ui = '*'
		}
		ShowUI(s.pos[i].X,s.pos[i].Y,ui)
	}
	//接受键盘按键信息
	//go添加一个独立函数，非阻塞运行
	go func(){
		for {
			switch xiaoyu1.Direction() {
			case 87,119,72:
				s.dir = 'U'
			case 83,115,80:
				s.dir = 'D'
			case 65,97,75:
				s.dir = 'L'
			case 68,100,77:
				s.dir = 'R'
			}
		}
	}()
}
//游戏逻辑
func (s *Snack) PlayGame(){
	for {
		//延迟模型
		time.Sleep(time.Second/2)
		//设置一个蛇的坐标移动量
		var nx,ny int = 0,0
		//更新蛇位置
		switch s.dir {
		case 'U':
			nx = 0
			ny = -1
		case 'D':
			nx = 0
			ny = 1
		case 'L':
			nx = -1
			ny = 0
		case 'R':
			nx = 1
			ny = 0
		}

		//蛇和墙体判断
		if s.pos[0].X < 0+5 ||
			s.pos[0].X>= WIDE+1||
			s.pos[0].Y<0||
			s.pos[0].Y>=HIGH{
			return
		}
		//蛇和身体判断
		for i:=1;i<s.size;i++{
			if s.pos[0].X == s.pos[i].X &&s.pos[0].Y == s.pos[i].Y{
				return
			}
		}
		//蛇和食物判断
		if s.pos[0].X == food.X &&s.pos[0].Y == food.Y{
			//身体+1
			s.size++
			RandomFood()
			//分数变量
		}
		//获取蛇尾坐标
		wx = s.pos[s.size-1].X
		wy = s.pos[s.size-1].Y
		//从尾部开始更新身体坐标
		for i := s.size-1;i>0;i--{
			s.pos[i].X = s.pos[i-1].X
			s.pos[i].Y = s.pos[i-1].Y
		}
		//更新蛇头坐标
		s.pos[0].X += nx
		s.pos[0].Y += ny
		//绘制蛇
		for i:=0;i<s.size;i++{
			var ui byte
			if i == 0{
				ui = '@'
			}else{
				ui = '*'
			}
			ShowUI(s.pos[i].X,s.pos[i].Y,ui)
		}
		ShowUI(wx,wy,' ')
	}
}

//"
//	#--------------------#
//	|                    |
//	|                    |
//	|                    |
//	|                    |
//	|                    |
//	|                    |
//	|                    |
//	|                    |
//	|                    |
//	|                    |
//	|                    |
//	|                    |
//	|                    |
//	|                    |
//	|                    |
//	|                    |
//	|                    |
//	|                    |
//	|                    |
//	|                    |
//	#--------------------#
//"

func MapInit(){
	fmt.Fprintln(os.Stderr,"\n\t#--------------------#\n\t|                    |\n\t|                    |\n\t|                    |\n\t|                    |\n\t|                    |\n\t|                    |\n\t|                    |\n\t|                    |\n\t|                    |\n\t|                    |\n\t|                    |\n\t|                    |\n\t|                    |\n\t|                    |\n\t|                    |\n\t|                    |\n\t|                    |\n\t|                    |\n\t|                    |\n\t|                    |\n\t#--------------------#")
}

// 绘制图形
func ShowUI(X int,Y int,ui byte){
	//找到对应坐标点光标位置
	xiaoyu1.GotoPosition(X*2+2,Y+2)
	fmt.Fprintf(os.Stderr,"%c",ui)
}

func main(){
	rand.Seed(time.Now().UnixNano())
	xiaoyu1.HideCursor()
	MapInit()
	RandomFood()
	var s Snack
	s.InitSnack()
	//fmt.Println(s)
	s.PlayGame()
	fmt.Println(food)
	//xiaoyu1.Test()
}

