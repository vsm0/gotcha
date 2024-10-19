package test

import (
	"github.com/vsm0/gotcha/model"

	"testing"
)

func TestCreateAccount(t *testing.T) {
	db := getDb(t)
	migrate(t, db)

	pwd := "123"
	a := &model.Account{
		Username: "Anon",
		Password: hash(pwd),
	}
	
	if res := db.Create(a); res.Error != nil {
		t.Fatalf("%v", res.Error)
	}
}

func TestLogin(t *testing.T) {
	db := getDb(t)
	migrate(t, db)

	pwd := "123"
	a := &model.Account{
		Username: "Anon",
	}

	if res := db.First(a); res.Error != nil {
		t.Fatalf("%v", res.Error)
	}

	t.Logf("Found account: %s", a.Username)

	if hash(pwd) != a.Password {
		t.Fatalf("Password doesn't match")
	}
}

func TestChangePassword(t *testing.T) {
	db := getDb(t)
	migrate(t, db)

	a := &model.Account{
		Username: "Anon",
	}

	if res := db.First(a); res.Error != nil {
		t.Fatalf("%v", res.Error)
	}

	t.Logf("Found account: %s", a.Username)

	pwd := "123"
	match := hash(pwd) == a.Password
	t.Logf("Confirm password... %t", match)

	if !match {
		t.Fatalf("Password doesn't match")
	}

	t.Logf("Setting new password...")

	newPwd := "456"
	a.Password = hash(newPwd)
	if res := db.Save(a); res.Error != nil {
		t.Fatalf("%v", res.Error)
	}
}

func TestDeleteAccount(t *testing.T) {
	db := getDb(t)
	migrate(t, db)

	pwd := "456"
	a := &model.Account{
		Username: "Anon",
		Password: hash(pwd),
	}

	if res := db.First(a); res.Error != nil {
		t.Fatalf("%v", res.Error)
	}

	t.Logf("Found account: %s", a.Username)

	if res := db.Delete(a); res.Error != nil {
		t.Fatalf("%v", res.Error)
	}
}
