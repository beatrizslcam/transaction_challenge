package login_test

import (
	"testing"
	"transactions/domain/login"
	"transactions/entity"
	"transactions/service"

	"golang.org/x/crypto/bcrypt"
)


 func TestLogin(t *testing.T){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("#thisIsMyPassword"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
}
	t.Run("login successfully", func(t *testing.T){
		t.Parallel()


		mockRepo := service.MockingFindByCpf(true, entity.Account{ID: "12345", Secret: string(hashedPassword)})
		mockAuth:= service.MockingGenerateToken(func(accountID string) string {
            return "mockedToken"})
		
		manageLogin := login.ManageLogin{Repo: mockRepo, Auth: mockAuth}

		result := manageLogin.Login("12345", "#thisIsMyPassword")
		expect := "mockedToken"

		if result != expect {
			t.Errorf("got %v want %v", result, expect)
		}
	})
	t.Run("can't not login ", func(t *testing.T){
		t.Parallel()


		mockRepo := service.MockingFindByCpf(true, entity.Account{ID: "12345", Secret: string(hashedPassword)})
		
		manageLogin := login.ManageLogin{Repo: mockRepo}

		result := manageLogin.Login("12345", "#thisISN'TMyPassword")
		expect := ""

		if result != expect {
			t.Errorf("got %v want %v", result, expect)
		}
	})
 }

