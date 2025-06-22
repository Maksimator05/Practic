package main

import "fmt"

// Двигатель - вложенная структура
type Engine struct {
	Power        int     // мощность в л.с.
	Volume       float32 // объем в литрах
	FuelType     string  // тип топлива
	IsTurbo      bool    // наличие турбины
	Manufacturer string  // производитель
}

// Автомобиль - основная структура
type Car struct {
	Brand      string
	Model      string
	Year       int
	Mileage    int    // пробег в км
	EngineInfo Engine // вложенная структура двигателя
}

func main() {
	// Создаем экземпляр автомобиля с вложенным двигателем
	myCar := Car{
		Brand:   "Toyota",
		Model:   "Camry",
		Year:    2022,
		Mileage: 15000,
		EngineInfo: Engine{
			Power:        249,
			Volume:       2.5,
			FuelType:     "бензин",
			IsTurbo:      true,
			Manufacturer: "Toyota Motor Corporation",
		},
	}

	// Выводим информацию об автомобиле
	fmt.Println("Информация об автомобиле:")
	fmt.Printf("Марка: %s\n", myCar.Brand)
	fmt.Printf("Модель: %s\n", myCar.Model)
	fmt.Printf("Год выпуска: %d\n", myCar.Year)
	fmt.Printf("Пробег: %d км\n", myCar.Mileage)

	// Выводим информацию о двигателе
	fmt.Println("\nХарактеристики двигателя:")
	fmt.Printf("Мощность: %d л.с.\n", myCar.EngineInfo.Power)
	fmt.Printf("Объем: %.1f л\n", myCar.EngineInfo.Volume)
	fmt.Printf("Тип топлива: %s\n", myCar.EngineInfo.FuelType)
	fmt.Printf("Турбина: %v\n", myCar.EngineInfo.IsTurbo)
	fmt.Printf("Производитель: %s\n", myCar.EngineInfo.Manufacturer)

	// Модифицируем некоторые параметры
	myCar.Mileage = 18000
	myCar.EngineInfo.IsTurbo = false

	// Выводим обновленную информацию
	fmt.Println("\nОбновленная информация:")
	fmt.Printf("Новый пробег: %d км\n", myCar.Mileage)
	fmt.Printf("Турбина после модификации: %v\n", myCar.EngineInfo.IsTurbo)

	// Создаем функцию для вывода полной информации
	printCarInfo(myCar)
}

// Функция для вывода полной информации об автомобиле
func printCarInfo(c Car) {
	fmt.Println("\nПолная информация об автомобиле (из функции):")
	fmt.Printf("%s %s %d года, пробег %d км\n", c.Brand, c.Model, c.Year, c.Mileage)
	fmt.Printf("Двигатель: %.1f л, %d л.с., %s, турбина: %v\n",
		c.EngineInfo.Volume, c.EngineInfo.Power, c.EngineInfo.FuelType, c.EngineInfo.IsTurbo)
}
