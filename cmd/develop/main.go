package main

import (
	"fmt"
	"log"

	nmap "github.com/Ullaakut/nmap/v2"
)

func scanSubnet(subnet string) {
	// Создаем новый сканер с пинг-сканированием
	scanner, err := nmap.NewScanner(
		nmap.WithTargets(subnet),
		nmap.WithPingScan(), // Используем только пинг-сканирование
	)
	if err != nil {
		log.Fatalf("Ошибка создания сканера: %v", err)
	}

	// Выполняем сканирование
	result, _, err := scanner.Run()
	if err != nil {
		log.Fatalf("Ошибка выполнения сканирования: %v", err)
	}

	// Обрабатываем результаты
	for _, host := range result.Hosts {
			fmt.Printf("Host %s is up\n", host.Addresses[0].Addr)
	}
}

func main() {
	subnet := "10.4.1.0/24" // Замените на нужную подсеть
	scanSubnet(subnet)
}
