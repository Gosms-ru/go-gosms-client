package gosms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	baseURL = "https://api.gosms.ru/v1"

	// Эндпоинты API
	endpointSMSSend = "/sms/send"
	endpointSMSGet  = "/sms/get"
	endpointSMSDel  = "/sms/del"
	endpointSMSList = "/sms"

	// Эндпоинты для работы с устройствами
	endpointDeviceInfo = "/devices/get/info"
	endpointDeviceEdit = "/devices/edit"
	endpointDeviceDel  = "/devices/del"
)

// Client представляет собой клиент для работы с API GoSMS
type Client struct {
	token      string
	httpClient *http.Client
}

// NewClient создает новый экземпляр клиента GoSMS
func NewClient(token string) *Client {
	return &Client{
		token:      token,
		httpClient: &http.Client{},
	}
}

// SendSMSRequest представляет собой структуру запроса на отправку SMS
type SendSMSRequest struct {
	Message     string `json:"message"`      // обязательное поле
	PhoneNumber string `json:"phone_number"` // обязательное поле
	DeviceID    string `json:"device_id,omitempty"`
	ToSim       int    `json:"to_sim,omitempty"`
	CallbackID  string `json:"callback_id,omitempty"`
}

// SendSMSResponse представляет собой структуру ответа на отправку SMS
type SendSMSResponse struct {
	ID string `json:"id"`
}

// GetSMSRequest представляет собой структуру запроса на получение информации об SMS
type GetSMSRequest struct {
	ID string `json:"id"` // обязательное поле
}

// GetSMSResponse представляет собой структуру ответа с информацией об SMS
type GetSMSResponse struct {
	ID            string `json:"id"`
	Message       string `json:"message"`
	Status        int    `json:"status"`
	CallbackID    string `json:"callback_id"`
	DeviceID      string `json:"device_id"`
	PhoneNumber   string `json:"phone_number"`
	MessageStatus string `json:"message_status"`
	TimeCreate    int64  `json:"time_create"`
	ToSim         *int   `json:"to_sim"`
}

// DeleteSMSRequest представляет собой структуру запроса на удаление SMS
type DeleteSMSRequest struct {
	ID string `json:"id"` // обязательное поле
}

// ListSMSRequest представляет собой структуру запроса на получение списка SMS
type ListSMSRequest struct {
	Limit  int    `json:"limit"`            // обязательное поле, минимум 1, максимум 100
	Offset int    `json:"offset"`           // опциональное поле, по умолчанию 1
	Search string `json:"search,omitempty"` // опциональное поле для поиска
}

// Pagination представляет собой структуру пагинации
type Pagination struct {
	TotalRecords int `json:"total_records"`
	Limit        int `json:"limit"`
	Offset       int `json:"offset"`
}

// ListSMSResponse представляет собой структуру ответа со списком SMS
type ListSMSResponse struct {
	Pagination Pagination       `json:"pagination"`
	SMSList    []GetSMSResponse `json:"sms_list"`
}

// SIMCard представляет информацию о SIM-карте
type SIMCard struct {
	SlotIndex   int    `json:"slot_index"`
	DisplayName string `json:"display_name"`
}

// GetDeviceInfoRequest представляет запрос на получение информации об устройстве
type GetDeviceInfoRequest struct {
	DeviceID string `json:"device_id"` // обязательное поле
}

// GetDeviceInfoResponse представляет ответ с информацией об устройстве
type GetDeviceInfoResponse struct {
	DeviceID           string    `json:"device_id"`
	DeviceBatteryState int       `json:"device_battery_state"`
	DeviceName         string    `json:"device_name"`
	IsActive           bool      `json:"is_active"`
	IsCharging         bool      `json:"is_charging"`
	LastOnlineDate     string    `json:"last_online_date"`
	DeviceNameType     string    `json:"device_name_type"`
	LowBatteryAlert    bool      `json:"low_battery_alert"`
	ToSim              int       `json:"to_sim"`
	SimList            []SIMCard `json:"sim_list"`
}

// EditDeviceRequest представляет запрос на редактирование устройства
type EditDeviceRequest struct {
	DeviceID string `json:"device_id"`           // обязательное поле
	IsActive bool   `json:"is_active,omitempty"` // опциональное поле
}

// DeleteDeviceRequest представляет запрос на удаление устройства
type DeleteDeviceRequest struct {
	DeviceID string `json:"device_id"` // обязательное поле
}

// SendSMS отправляет SMS сообщение
func (c *Client) SendSMS(req SendSMSRequest) (*SendSMSResponse, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка маршалинга запроса: %w", err)
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s%s", baseURL, endpointSMSSend), bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("ошибка создания запроса: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("неуспешный статус ответа: %d", response.StatusCode)
	}

	var smsResponse SendSMSResponse
	if err := json.NewDecoder(response.Body).Decode(&smsResponse); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	return &smsResponse, nil
}

// GetSMS получает информацию об SMS по его ID
func (c *Client) GetSMS(req GetSMSRequest) (*GetSMSResponse, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка маршалинга запроса: %w", err)
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s%s", baseURL, endpointSMSGet), bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("ошибка создания запроса: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("неуспешный статус ответа: %d", response.StatusCode)
	}

	var smsResponse GetSMSResponse
	if err := json.NewDecoder(response.Body).Decode(&smsResponse); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	return &smsResponse, nil
}

// DeleteSMS удаляет SMS по его ID
func (c *Client) DeleteSMS(req DeleteSMSRequest) error {
	payload, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("ошибка маршалинга запроса: %w", err)
	}

	request, err := http.NewRequest("DELETE", fmt.Sprintf("%s%s", baseURL, endpointSMSDel), bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("ошибка создания запроса: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	response, err := c.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusNoContent {
		return fmt.Errorf("неуспешный статус ответа: %d", response.StatusCode)
	}

	return nil
}

// ListSMS получает список SMS с пагинацией
func (c *Client) ListSMS(req ListSMSRequest) (*ListSMSResponse, error) {
	// Проверяем обязательные параметры
	if req.Limit < 1 || req.Limit > 100 {
		return nil, fmt.Errorf("limit должен быть от 1 до 100")
	}

	// Формируем URL с параметрами запроса
	url := fmt.Sprintf("%s/sms?limit=%d", baseURL, req.Limit)
	if req.Offset > 0 {
		url = fmt.Sprintf("%s&offset=%d", url, req.Offset)
	}
	if req.Search != "" {
		url = fmt.Sprintf("%s&search=%s", url, req.Search)
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания запроса: %w", err)
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("неуспешный статус ответа: %d", response.StatusCode)
	}

	var smsResponse ListSMSResponse
	if err := json.NewDecoder(response.Body).Decode(&smsResponse); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	return &smsResponse, nil
}

// GetDeviceInfo получает информацию об устройстве
func (c *Client) GetDeviceInfo(req GetDeviceInfoRequest) (*GetDeviceInfoResponse, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка маршалинга запроса: %w", err)
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s%s", baseURL, endpointDeviceInfo), bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("ошибка создания запроса: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("неуспешный статус ответа: %d", response.StatusCode)
	}

	var deviceResponse GetDeviceInfoResponse
	if err := json.NewDecoder(response.Body).Decode(&deviceResponse); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	return &deviceResponse, nil
}

// EditDevice редактирует настройки устройства
func (c *Client) EditDevice(req EditDeviceRequest) error {
	payload, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("ошибка маршалинга запроса: %w", err)
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s%s", baseURL, endpointDeviceEdit), bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("ошибка создания запроса: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	response, err := c.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusNoContent {
		return fmt.Errorf("неуспешный статус ответа: %d", response.StatusCode)
	}

	return nil
}

// DeleteDevice удаляет устройство
func (c *Client) DeleteDevice(req DeleteDeviceRequest) error {
	payload, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("ошибка маршалинга запроса: %w", err)
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s%s", baseURL, endpointDeviceDel), bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("ошибка создания запроса: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	response, err := c.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusNoContent {
		return fmt.Errorf("неуспешный статус ответа: %d", response.StatusCode)
	}

	return nil
}
