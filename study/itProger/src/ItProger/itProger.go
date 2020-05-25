package ItProger

import "fmt"

func ItProger() {
	// fmt.Println("Helo world")

	var age int = 20
	var num float64 = 2.3456
	var str = "Helo world"
	var res int

	fmt.Println(age, num, str)

	var a = 10
	var b = 5

	res = a % b // остаток от деления
	fmt.Println(res)

	fmt.Printf("%f \n", num)
	fmt.Printf("%.2f \n", num)
	fmt.Printf("%T \n", num) // get type

	var isDone bool = true

	fmt.Printf("%t \n", isDone)

	if age < 5 {
		fmt.Println("Вам пора в детский сад")
	} else if age == 5 {
		fmt.Println("Вам пора в школу")
	} else if (age > 5) && (age <= 18) {
		var grade = age - 5
		fmt.Println("Пора идтии в ", grade, "класс")
	} else {
		fmt.Println("Вам пора в университет")
	}

	fmt.Println()

	switch age {
	case 5:
		fmt.Printf("\n%n Вам 5 лет")
	case 15:
		fmt.Printf("\n%n Вам 15 лет")
	default:
		fmt.Printf("\n%n Вам неизвестно сколько лет", 1)
	}

	var i = 1
	fmt.Println()

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	var arr [3]int
	arr[0] = 45
	arr[1] = 97
	arr[2] = 76

	nums := [3]float64{4.23, 5.23, 98.1}

	for i, value := range nums {
		fmt.Println(value, i)
	}

	webSutes := make(map[string]float64)
	webSutes["itProger"] = 0.8
	webSutes["yandex"] = 0.99
	fmt.Println(webSutes["itProger"])

	var a2 = 29
	var b2 = 1
	r := summ(a2, b2)
	fmt.Println("Сумма чисел равна: ", r)

	//замыкание

	var nums1 = 3
	multiple := func() int {
		nums1 *= 3
		return nums1
	}
	fmt.Println(multiple())

	//откладывание defer

	defer two()
	one()

	// указатель

	var x = 0
	pointer(&x)
	fmt.Println(x)

	bob := Cats{"Bob", 7, 0.878}
	fmt.Println("Bob age is", bob.age)

	fmt.Println(bob.test)

}

func (cat *Cats) test() float64 {
	return float64(cat.age) * cat.happiness

}

type Cats struct {
	name      string
	age       int
	happiness float64
}

func pointer(x *int) {
	*x = 2
}

func one() {
	fmt.Println("1")
}

func two() {
	fmt.Println("2")
}

func summ(num_1 int, num_2 int) int {
	return num_1 + num_2
}
