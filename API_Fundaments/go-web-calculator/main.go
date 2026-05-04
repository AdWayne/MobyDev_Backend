package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/calculate", calculateHandler)
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается. Используйте POST.", http.StatusMethodNotAllowed)
		return
	}

	var a, b float64
	var op string

	_, err := fmt.Fscan(r.Body, &a, &b, &op)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Ошибка! Отправьте данные в формате: 10 20 +")
		return
	}

	var result float64

	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Ошибка: Деление на ноль")
			return
		}
		result = a / b
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Неизвестная операция: %s", op)
		return
	}

	fmt.Fprintf(w, "Результат: %g", result)
}
