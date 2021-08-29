package xiaoyu1

/*
//#include <windows.h>
//#include <conio.h>
#include <ncurses.h>

//使用winapi来移动控制台光标
void gotoxy(int x,int y){
	COORD c;
	c.X=x,c.Y=y;
	SetConsoleCursorPosition(GetStdHandle(STD_OUTPUT_HANDLE),c);
}

int direct(){
	return _getch();
}

void hideCursor(){
	CONSOLE_CURSOR_INFO cci;
	cci.bVisible = FALSE;
	cci.dwSize = sizeof(cci);
	SetConsoleCursorInfo(GetStdHandle(STD_OUTPUT_HANDLE),&cci);
}
*/
import "C"

//设置控制台光标位置
func GotoPosition(X int,Y int){
	C.gotoxy(C.int(X),C.int(Y))
}
//无显获取键盘输入字符
func Direction()(key int){
	key = int(C.direct())
	return
}
//设置控制台光标隐藏
func HideCursor()  {
	C.hideCursor()
}

/*
#include <stdio.h>
#include <ncurses.h>
int demo(){
    WINDOW *snakeys_world;
    int offsetx, offsety;

    initscr();
    refresh();

    offsetx = (COLS - 20) / 2;
    offsety = (LINES - 20) / 2;

    snakeys_world = newwin(20,
                           20,
                           offsety,
                           offsetx);

    box(snakeys_world, 0 , 0);

    wrefresh(snakeys_world);

    getch();

    delwin(snakeys_world);

    endwin();

	printf("hello world");
	return 0;
}
*/
import "C"    // 该import与c代码直接不能有空格

func Test(){
	C.demo()
}