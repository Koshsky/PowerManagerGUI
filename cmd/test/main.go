package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Koshsky/PowerManagerGUI/internal/netutils"
)

func main() {
	testPingFunc()
}

func testPingFunc() {
	popularIPs := []string{
		"8.8.8.8",         // Google DNS
		"1.1.1.1",         // Cloudflare DNS
		"208.67.222.222",  // OpenDNS
		"205.251.242.103", // Amazon AWS
		"40.112.72.205",   // Microsoft Azure
	}

	for _, ip := range popularIPs {
		fmt.Printf("Запуск Ping на %s...\n", ip)
		success, duration, err := netutils.Ping(ip, 80) // Используем порт 80 для TCP Ping
		if err != nil {
			log.Printf("Ошибка при выполнении Ping на %s: %v\n", ip, err)
			continue
		}

		if success {
			fmt.Printf("Ping на %s успешен! Время ответа: %v\n", ip, duration)
		} else {
			fmt.Printf("Ping на %s не удался.\n", ip)
		}

		// Задержка между запросами
		time.Sleep(1 * time.Second)
	}

	fmt.Println("\nТестирование Ping завершено.")

	// Вызов функции для тестирования пользовательского ввода
	testScan()
}

func testScan() {
	// Пример вызова других функций сканирования
	fmt.Println("\nНачинаем сканирование GERS...")
	gersIPs, err := netutils.ScanGERS()
	if err != nil {
		log.Fatalf("Ошибка при сканировании GERS: %v\n", err)
	}
	fmt.Println("Доступные IP-адреса GERS:", gersIPs)

	fmt.Println("\nНачинаем сканирование KUUFS...")
	kuufsIPs, err := netutils.ScanKUUFS()
	if err != nil {
		log.Fatalf("Ошибка при сканировании KUUFS: %v\n", err)
	}
	fmt.Println("Доступные IP-адреса KUUFS:", kuufsIPs)

	fmt.Println("\nНачинаем сканирование всей сети...")
	allIPs, err := netutils.ScanNetwork()
	if err != nil {
		log.Fatalf("Ошибка при сканировании сети: %v\n", err)
	}
	fmt.Println("Доступные IP-адреса в сети:", allIPs)

	fmt.Println("\nСканирование завершено.")
}
