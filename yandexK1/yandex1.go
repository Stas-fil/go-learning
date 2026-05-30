package main

import (
	"fmt"
	"os"
)

func main() {
	var troom int
	var tcond int
	var mode string

	_, err := fmt.Scan(&troom, &tcond)
	if err != nil || tcond < -50 || tcond > 50 || troom < -50 || troom > 50 {
		os.Exit(1)
	}
	fmt.Scan(&mode)

	switch mode {
	case "freeze":
		// Охлаждение - может только уменьшать температуру
		if troom > tcond {
			fmt.Println(tcond)
		} else {
			fmt.Println(troom)
		}

	case "heat":
		// Нагрев - может только увеличивать температуру
		if troom < tcond {
			fmt.Println(tcond)
		} else {
			fmt.Println(troom)
		}

	case "auto":
		// Автоматический режим - может как увеличивать, так и уменьшать
		fmt.Println(tcond)

	case "fan":
		// Вентиляция - не изменяет температуру
		fmt.Println(troom)
	}
}
