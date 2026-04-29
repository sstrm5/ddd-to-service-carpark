package main

const inventoryPath = "data/inventory.csv"

func main() {
	// // Создаём зависимости
	// partRepository := partRepo.NewRepository()
	// spaceshipRepository := spaceshipRepo.NewRepository()

	// // Создаём сервисы (каждый со своими зависимостями)
	// partSvc := partService.NewService(partRepository)
	// assemblySvc := assemblyService.NewService(partRepository, spaceshipRepository)

	// // Создаём handlers
	// partH := part.NewHandler(partSvc)
	// assemblyH := assemblyHandler.NewHandler(assemblySvc)

	// // Загружаем данные из CSV
	// if err := partRepository.LoadFromCSV(inventoryPath); err != nil {
	// 	log.Printf("Не удалось загрузить данные: %v", err)
	// }

	// log.Println("Сервер запущен на http://localhost:8080")

	// // Регистрируем маршруты
	// mux := http.NewServeMux()
	// handler.RegisterRoutes(mux, partH, assemblyH)

	// server := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: mux,
	// 	// ReadTimeout — максимальное время на чтение всего запроса (заголовки + тело)
	// 	ReadTimeout: 10 * time.Second,
	// 	// WriteTimeout — максимальное время на запись ответа клиенту
	// 	WriteTimeout: 10 * time.Second,
	// 	// IdleTimeout — максимальное время ожидания следующего запроса при keep-alive соединении
	// 	IdleTimeout: 60 * time.Second,
	// }
	// log.Fatal(server.ListenAndServe())
}
