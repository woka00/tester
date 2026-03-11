package main

func main() {
	// _, err := os.Create("out/newfile.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Запуск HTTP сервера")
	// fmt.Println("второй вывод")
	// err := http_server.StartHTTPServer()
	// if err != nil {
	// 	fmt.Println("Во время работы сервера произошла ошибка:", err)
	// } else {
	// 	fmt.Println("Сервер завершился успешно")
	// }docker run postgres:18-bookworm

	// ctx := context.Background()

	// conn, err := simple_connection.CreateConnection(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// if err := simple_sql.CreateTable(ctx, conn); err != nil {
	// 	panic(err)
	// }

	// tasks, err := simple_sql.SelectRows(ctx, conn)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, task := range tasks {
	// 	if task.ID == 3 {
	// 		task.Title = "Покормить кошку"
	// 		task.Description = "30 г корма"
	// 		task.Completed = true
	// 		now := time.Now()
	// 		task.CompletedAt = &now

	// 		if err := simple_sql.UpdateTask(ctx, conn, task); err != nil {
	// 			panic(err)
	// 		}

	// 		break
	// 	}

	// }
	// fmt.Println("succeed")
}
