package gosms

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	token := "test-token"
	client := NewClient(token)

	if client == nil {
		t.Error("NewClient вернул nil")
	}

	if client.token != token {
		t.Errorf("ожидался токен %s, получен %s", token, client.token)
	}
}

func TestSendSMS(t *testing.T) {
	client := NewClient("test-token")

	req := SendSMSRequest{
		Message:     "Тестовое сообщение",
		PhoneNumber: "79999999999",
		DeviceID:    "test-device",
		ToSim:       1,
		CallbackID:  "test-callback",
	}

	// Примечание: этот тест будет падать без реального токена
	// В реальном проекте здесь нужно использовать mock-сервер
	_, err := client.SendSMS(req)
	if err == nil {
		t.Error("ожидалась ошибка из-за неверного токена")
	}
}

func TestGetSMS(t *testing.T) {
	client := NewClient("test-token")

	req := GetSMSRequest{
		ID: "6654a4e8f1527149588c89f2",
	}

	// Примечание: этот тест будет падать без реального токена
	// В реальном проекте здесь нужно использовать mock-сервер
	_, err := client.GetSMS(req)
	if err == nil {
		t.Error("ожидалась ошибка из-за неверного токена")
	}
}

func TestDeleteSMS(t *testing.T) {
	client := NewClient("test-token")

	req := DeleteSMSRequest{
		ID: "6654a4e8f1527149588c89f2",
	}

	// Примечание: этот тест будет падать без реального токена
	// В реальном проекте здесь нужно использовать mock-сервер
	err := client.DeleteSMS(req)
	if err == nil {
		t.Error("ожидалась ошибка из-за неверного токена")
	}
}

func TestListSMS(t *testing.T) {
	client := NewClient("test-token")

	req := ListSMSRequest{
		Limit:  5,
		Offset: 1,
		Search: "79999999999",
	}

	// Примечание: этот тест будет падать без реального токена
	// В реальном проекте здесь нужно использовать mock-сервер
	_, err := client.ListSMS(req)
	if err == nil {
		t.Error("ожидалась ошибка из-за неверного токена")
	}

	// Тест на некорректный limit
	req.Limit = 0
	_, err = client.ListSMS(req)
	if err == nil {
		t.Error("ожидалась ошибка из-за некорректного limit")
	}

	req.Limit = 101
	_, err = client.ListSMS(req)
	if err == nil {
		t.Error("ожидалась ошибка из-за некорректного limit")
	}
}

func TestGetDeviceInfo(t *testing.T) {
	client := NewClient("test-token")

	req := GetDeviceInfoRequest{
		DeviceID: "b1277815-fb6f-45e4-b87b-8dfb86b8f0a2",
	}

	// Примечание: этот тест будет падать без реального токена
	// В реальном проекте здесь нужно использовать mock-сервер
	_, err := client.GetDeviceInfo(req)
	if err == nil {
		t.Error("ожидалась ошибка из-за неверного токена")
	}
}

func TestEditDevice(t *testing.T) {
	client := NewClient("test-token")

	req := EditDeviceRequest{
		DeviceID: "b1277815-fb6f-45e4-b87b-8dfb86b8f0a2",
		IsActive: true,
	}

	// Примечание: этот тест будет падать без реального токена
	// В реальном проекте здесь нужно использовать mock-сервер
	err := client.EditDevice(req)
	if err == nil {
		t.Error("ожидалась ошибка из-за неверного токена")
	}
}

func TestDeleteDevice(t *testing.T) {
	client := NewClient("test-token")

	req := DeleteDeviceRequest{
		DeviceID: "b1277815-fb6f-45e4-b87b-8dfb86b8f0a2",
	}

	// Примечание: этот тест будет падать без реального токена
	// В реальном проекте здесь нужно использовать mock-сервер
	err := client.DeleteDevice(req)
	if err == nil {
		t.Error("ожидалась ошибка из-за неверного токена")
	}
}
