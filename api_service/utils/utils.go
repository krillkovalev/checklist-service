package utils

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io"
	"bytes"
)

func ProxyRequest(client *http.Client, method, url string, body interface{}) ([]byte, error) {
    // Преобразуем тело запроса в JSON
    jsonBytes, err := json.Marshal(body)
    if err != nil {
        return nil, fmt.Errorf("error marshaling request body: %v", err)
    }

    // Создаем HTTP-запрос
    req, err := http.NewRequest(method, url, bytes.NewReader(jsonBytes))
    if err != nil {
        return nil, fmt.Errorf("failed to create request: %v", err)
    }

    // Отправляем запрос
    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("failed to send request: %v", err)
    }
    defer resp.Body.Close()

    // Проверяем статус-код ответа
    if resp.StatusCode != http.StatusOK {
        body, _ := io.ReadAll(resp.Body)
        return nil, fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(body))
    }

    // Читаем тело ответа
    responseBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to read response body: %v", err)
    }

    return responseBody, nil
}