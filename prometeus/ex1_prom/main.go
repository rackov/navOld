package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Создаем счетчик, который будет увеличиваться при каждом запросе
	httpRequestsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "The total number of HTTP requests",
	})

	// Создаем гистограмму для отслеживания времени обработки запросов
	httpRequestDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Duration of HTTP requests in seconds",
		Buckets: prometheus.DefBuckets, // Стандартные бакеты для гистограммы
	})
)

// Функция для обработки HTTP запросов
func handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	// Увеличиваем счетчик запросов
	httpRequestsTotal.Inc()

	// Имитируем некоторую работу
	time.Sleep(time.Duration(time.Millisecond * 200))
	// Записываем время обработки запроса
	httpRequestDuration.Observe(time.Since(start).Seconds())
	w.Write([]byte("Hello, world!\n"))
}

func main() {
	// Задаем порт для экспорта метрик
	addr := flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
	flag.Parse()

	// Регистрируем обработчик для метрик
	http.Handle("/metrics", promhttp.Handler())

	// Регистрируем обработчик для запросов
	http.HandleFunc("/", handler)

	// Запускаем HTTP сервер
	log.Fatal(http.ListenAndServe(*addr, nil))
}
