package main

import (
	"context"
	"fmt"
	"hw/feature_sql/connection"
	sqlrequestsgo "hw/feature_sql/sql_requests.go"
)

func main() {
	ctx := context.Background()
	conn, err := connection.ConnectionDB(ctx)
	if err != nil {
		panic(err)
	}
	err = conn.Ping(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Успешное подключение!")

	if err := sqlrequestsgo.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	// book := []sqlrequestsgo.BookModel{
	// 	{
	// 		Name:          "Мастер и Маргарита",
	// 		Author:        "Михаил Булгаков",
	// 		Resolution:    "Отличная книга",
	// 		Reales_year:   1967,
	// 		Completed:     false,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Преступление и наказание",
	// 		Author:        "Федор Достоевский",
	// 		Resolution:    "Глубокий психологический роман",
	// 		Reales_year:   1866,
	// 		Completed:     true,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Война и мир",
	// 		Author:        "Лев Толстой",
	// 		Resolution:    "Эпическое произведение русской литературы",
	// 		Reales_year:   1869,
	// 		Completed:     false,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "1984",
	// 		Author:        "Джордж Оруэлл",
	// 		Resolution:    "Антиутопия о тоталитаризме",
	// 		Reales_year:   1949,
	// 		Completed:     true,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	// Дополнительные книги (26 штук)
	// 	{
	// 		Name:          "Анна Каренина",
	// 		Author:        "Лев Толстой",
	// 		Resolution:    "Трагическая история любви",
	// 		Reales_year:   1877,
	// 		Completed:     true,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Гарри Поттер и философский камень",
	// 		Author:        "Джоан Роулинг",
	// 		Resolution:    "Начало магической саги",
	// 		Reales_year:   1997,
	// 		Completed:     true,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Маленький принц",
	// 		Author:        "Антуан де Сент-Экзюпери",
	// 		Resolution:    "Философская сказка для взрослых",
	// 		Reales_year:   1943,
	// 		Completed:     true,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Три товарища",
	// 		Author:        "Эрих Мария Ремарк",
	// 		Resolution:    "Роман о дружбе и любви",
	// 		Reales_year:   1936,
	// 		Completed:     false,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Улисс",
	// 		Author:        "Джеймс Джойс",
	// 		Resolution:    "Сложный модернистский роман",
	// 		Reales_year:   1922,
	// 		Completed:     false,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Властелин колец",
	// 		Author:        "Дж. Р. Р. Толкин",
	// 		Resolution:    "Эпическая фэнтези-сага",
	// 		Reales_year:   1954,
	// 		Completed:     true,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Гордость и предубеждение",
	// 		Author:        "Джейн Остин",
	// 		Resolution:    "Классика английской литературы",
	// 		Reales_year:   1813,
	// 		Completed:     true,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Сто лет одиночества",
	// 		Author:        "Габриэль Гарсиа Маркес",
	// 		Resolution:    "Магический реализм",
	// 		Reales_year:   1967,
	// 		Completed:     false,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Процесс",
	// 		Author:        "Франц Кафка",
	// 		Resolution:    "Абсурдистский роман",
	// 		Reales_year:   1925,
	// 		Completed:     false,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Великий Гэтсби",
	// 		Author:        "Фрэнсис Скотт Фицджеральд",
	// 		Resolution:    "Роман о \"ревущих двадцатых\"",
	// 		Reales_year:   1925,
	// 		Completed:     true,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Над пропастью во ржи",
	// 		Author:        "Джером Д. Сэлинджер",
	// 		Resolution:    "Роман о подростковом бунте",
	// 		Reales_year:   1951,
	// 		Completed:     true,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Мертвые души",
	// 		Author:        "Николай Гоголь",
	// 		Resolution:    "Поэма в прозе",
	// 		Reales_year:   1842,
	// 		Completed:     false,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Братья Карамазовы",
	// 		Author:        "Федор Достоевский",
	// 		Resolution:    "Философский роман",
	// 		Reales_year:   1880,
	// 		Completed:     false,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Одиссея",
	// 		Author:        "Гомер",
	// 		Resolution:    "Античный эпос",
	// 		Reales_year:   -800,
	// 		Completed:     true,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Дон Кихот",
	// 		Author:        "Мигель де Сервантес",
	// 		Resolution:    "Первый современный роман",
	// 		Reales_year:   1605,
	// 		Completed:     true,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Превращение",
	// 		Author:        "Франц Кафка",
	// 		Resolution:    "Новелла о метаморфозах",
	// 		Reales_year:   1915,
	// 		Completed:     true,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Фауст",
	// 		Author:        "Иоганн Вольфганг Гёте",
	// 		Resolution:    "Трагедия о познании",
	// 		Reales_year:   1808,
	// 		Completed:     false,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Гамлет",
	// 		Author:        "Уильям Шекспир",
	// 		Resolution:    "Трагедия мести",
	// 		Reales_year:   1603,
	// 		Completed:     true,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Скотный двор",
	// 		Author:        "Джордж Оруэлл",
	// 		Resolution:    "Аллегория на революцию",
	// 		Reales_year:   1945,
	// 		Completed:     true,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Портрет Дориана Грея",
	// 		Author:        "Оскар Уайльд",
	// 		Resolution:    "Роман о красоте и морали",
	// 		Reales_year:   1890,
	// 		Completed:     true,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Поющие в терновнике",
	// 		Author:        "Колин Маккалоу",
	// 		Resolution:    "Семейная сага",
	// 		Reales_year:   1977,
	// 		Completed:     false,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Мартин Иден",
	// 		Author:        "Джек Лондон",
	// 		Resolution:    "Роман о самореализации",
	// 		Reales_year:   1909,
	// 		Completed:     true,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Тень горы",
	// 		Author:        "Грегори Дэвид Робертс",
	// 		Resolution:    "Продолжение \"Шантарама\"",
	// 		Reales_year:   2015,
	// 		Completed:     false,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Атлант расправил плечи",
	// 		Author:        "Айн Рэнд",
	// 		Resolution:    "Философский роман об объективизме",
	// 		Reales_year:   1957,
	// 		Completed:     false,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Код да Винчи",
	// 		Author:        "Дэн Браун",
	// 		Resolution:    "Детективный триллер",
	// 		Reales_year:   2003,
	// 		Completed:     true,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Алиса в Стране чудес",
	// 		Author:        "Льюис Кэрролл",
	// 		Resolution:    "Детская сказка с философским подтекстом",
	// 		Reales_year:   1865,
	// 		Completed:     true,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Шерлок Холмс",
	// 		Author:        "Артур Конан Дойл",
	// 		Resolution:    "Сборник детективных рассказов",
	// 		Reales_year:   1887,
	// 		Completed:     true,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Убить пересмешника",
	// 		Author:        "Харпер Ли",
	// 		Resolution:    "Роман о расовой несправедливости",
	// 		Reales_year:   1960,
	// 		Completed:     true,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "451 градус по Фаренгейту",
	// 		Author:        "Рэй Брэдбери",
	// 		Resolution:    "Антиутопия о цензуре",
	// 		Reales_year:   1953,
	// 		Completed:     true,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// 	{
	// 		Name:          "Зов Ктулху",
	// 		Author:        "Говард Лавкрафт",
	// 		Resolution:    "Рассказ в жанре ужасов",
	// 		Reales_year:   1928,
	// 		Completed:     true,
	// 		Time_on_shelf: time.Now(),
	// 	},
	// }

	// for _, book := range book {
	// 	if err := sqlrequestsgo.InsertRow(ctx, conn, book); err != nil {
	// 		panic(err)
	// 	}
	// }

	books, err := sqlrequestsgo.SelectRow(ctx, conn)
	if err != nil {
		panic(err)
	}

	for _, book := range books {
		if book.ID == 2 {
			book.Name = "Да"
			book.Author = "точно"

			if err := sqlrequestsgo.UpdateRow(ctx, conn, book); err != nil {
				panic(err)
			}
			break
		}

	}

	if err := sqlrequestsgo.DelRow(ctx, conn, []int{1, 3, 5}); err != nil {
		panic(err)
	}
	fmt.Println("Start pagnation")
	if err := sqlrequestsgo.ListPages(ctx, conn, 3); err != nil {
		panic(err)
	}
}
