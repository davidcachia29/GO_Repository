package usecase

import "TestProject/infrastructure/entity"

type MessageUseCase struct {
	Message entity.Message
}

func (m *MessageUseCase) GetMessageText() map[string]string {
	return map[string]string{"message": m.Message.Text}
}
