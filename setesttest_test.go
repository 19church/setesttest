package setesttest

import (
	"gorm.io/gorm"

	"testing"

	"time"

	"github.com/asaskevich/govalidator"

	. "github.com/onsi/gomega"
)

type Abc struct {
	gorm.Model
	Name     string `valid:"required~no"`
	Number   int    `valid:"range(1|10)"`
	Code     string `gorm:"uniqueIndex" valid:"matches(^[BMD]\\d{7}$)"`
	Email    string `gorm:"uniqueIndex" valid:"email"`
	Url      string `gorm:"uniqueIndex" valid:"url"`
	Tel      string `gorm:"uniqueIndex" valid:"matches(^\\d{10}$)"`
	Password string `valid:"minstringlength(8)"`
}

func TestBest(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("check name not blank", func(t *testing.T) {
		a := Abc{
			Name:     "",
			Number:   1,
			Code:     "B6304249",
			Email:    "saint@gmail.com",
			Url:      "https://www.youtube.com/",
			Tel:      "0900800440",
			Password: "12345678",
		}

		ok, err := govalidator.ValidateStruct(a)

		g.Expect(ok).ToNot(BeTrue())

		g.Expect(err).ToNot(BeNil())

		g.Expect(err.Error()).To(Equal("no"))
	})
}

func TestNumber(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("check number", func(t *testing.T) {
		a := Abc{
			Name:     "abc",
			Number:   11,
			Code:     "B6304249",
			Email:    "saint@gmail.com",
			Url:      "https://www.youtube.com/",
			Tel:      "0900800440",
			Password: "12345678",
		}

		ok, err := govalidator.ValidateStruct(a)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).ToNot(BeNil())

		g.Expect(err.Error()).To(Equal("Number: 11 does not validate as range(1|10)"))
	})

	t.Run("check Code", func(t *testing.T) {
		a := Abc{
			Name:     "abc",
			Number:   1,
			Code:     "MB6304249",
			Email:    "saint@gmail.com",
			Url:      "https://www.youtube.com/",
			Tel:      "0900800440",
			Password: "12345678",
		}

		ok, err := govalidator.ValidateStruct(a)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Code: MB6304249 does not validate as matches(^[BMD]\\d{7}$)"))
	})
}

func TestEmail(t *testing.T) {
	g := NewGomegaWithT(t)

	a := Abc{
		Name:     "abc",
		Number:   1,
		Code:     "B6304249",
		Email:    "saint@gmail",
		Url:      "https://www.youtube.com/",
		Tel:      "0900800440",
		Password: "12345678",
	}

	ok, err := govalidator.ValidateStruct(a)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Email: saint@gmail does not validate as email"))
}

func TestAaaa(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("url", func(t *testing.T) {
		a := Abc{
			Name:     "abc",
			Number:   1,
			Code:     "B6304249",
			Email:    "saint@gmail.com",
			Url:      "www",
			Tel:      "0900800440",
			Password: "12345678",
		}
		ok, err := govalidator.ValidateStruct(a)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Url: www does not validate as url"))
	})
}

func TestTel(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Tel", func(t *testing.T) {
		a := Abc{
			Name:     "abc",
			Number:   1,
			Code:     "B6304249",
			Email:    "saint@gmail.com",
			Url:      "https://www.youtube.com",
			Tel:      "0",
			Password: "12345678",
		}

		ok, err := govalidator.ValidateStruct(a)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Tel: 0 does not validate as matches(^\\d{10}$)"))
	})
}

func TestPassword(t *testing.T) {
	g := NewGomegaWithT(t)
	
	t.Run("Password", func(t *testing.T){
		a := Abc{
			Name:     "abc",
			Number:   1,
			Code:     "B6304249",
			Email:    "saint@gmail.com",
			Url:      "https://www.youtube.com",
			Tel:      "0900800440",
			Password: "1",
		}

		ok, err := govalidator.ValidateStruct(a)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Password: 1 does not validate as minstringlength(8)"))
	})
}

func TestPassPass(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("PassPass", func(t *testing.T){
		a := Abc{
			Name:     "abc",
			Number:   1,
			Code:     "B6304249",
			Email:    "saint@gmail.com",
			Url:      "https://www.youtube.com",
			Tel:      "0900800440",
			Password: "12345678",
		}

		ok, err := govalidator.ValidateStruct(a)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}

func init() {
	govalidator.CustomTypeTagMap.Set("startdate", func(i interface{}, _ interface{}) bool {
		t := i.(time.Time)
		if t.Before(time.Now().Add(-60 * time.Minute)) {
			return false
		} else {
			return true
		}
	})

	govalidator.CustomTypeTagMap.Set("addedtime", func(i interface{}, _ interface{}) bool {
		t := i.(time.Time)
		if t.Before(time.Now().Add(-60*time.Minute)) || t.After(time.Now().Add(60*time.Minute)) {
			return false
		} else {
			return true
		}
	})

	govalidator.CustomTypeTagMap.Set("abc", func(i interface{}, _ interface{}) bool {
		t := i.(time.Time)
		if t.Before(time.Now().Add(-2*time.Minute)) || t.After(time.Now().Add(2*time.Minute)) {
			return false
		} else {
			return true
		}
	})

	govalidator.CustomTypeTagMap.Set("aasdfs", func(i interface{}, _ interface{}) bool {
		t := i.(time.Time)
		if t.Before(time.Now().Add(-2*time.Minute)) || t.After(time.Now().Add(2*time.Minute)) {
			return false
		} else {
			return true
		}
	})
}
