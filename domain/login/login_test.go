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
print("the hashed password is: %v",hashedPassword)
	t.Run("login succeful", func(t *testing.T){
		t.Parallel()


		mockRepo := service.MockingFindByCpf(true, entity.Account{ID: "12345", Secret: string(hashedPassword)})
		mockAuth:= service.MockingGenerateToken("12345")
		
		manageLogin := login.ManageLogin{Repo: mockRepo, Auth: mockAuth}

		result := manageLogin.Login("12345", "thisIsMyPassword")
		expect := "token"

		if result != expect {
			t.Errorf("got %v want %v", result, expect)
		}
	})
 }

