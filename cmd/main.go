package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println("Добро пожаловать в игру Угадай число!\nЗагадываю число от 1 до 100")
	fmt.Println("Выберите уровень сложности: 1 - легкий, 2 - средний, 3 - сложный")
	difficulty()
	gameLogic()

}

var try int
var input int

func difficulty() {
	var tryInput int
	_, err := fmt.Scan(&tryInput)
	if err != nil {
		fmt.Println("Ошибка ввода, введите 1,2 или 3")
		return
	}
	switch tryInput {
	case 1:
		fmt.Println("Уровень сложности: Легкий")
		try = 10

	case 2:
		fmt.Println("Уровень сложности: Средний")
		try = 5

	case 3:
		fmt.Println("Уровень сложности: Сложный")
		try = 3

	default:
		fmt.Println("Ошибка ввода, введите 1,2 или 3")
	}
}

func gameLogic() {
	target := rand.IntN(100) + 1
	fmt.Println("Введите число от 1 до 100")
	for i := 0; i < try; i++ {
		fmt.Scan(&input)
		if input == target {
			fmt.Println("Поздравляю! Вы угадали за", i+1, "попыток")
			return
		}
		if input > target {
			fmt.Println("Число меньше чем", input, "\nОсталось попыток:", try-i-1)
		}
		if input < target {
			fmt.Println("Число больше чем", input, "\nОсталось попыток:", try-i-1)
		}
	}
	fmt.Println("Вы проиграли. Было загадано число:", target)
}
