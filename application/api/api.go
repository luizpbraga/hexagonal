package api

import (
	"github.com/luizpbraga/hexagonal/ports"
)

type Arithmetic interface {
	Addition(int32, int32) (int32, error)
	Division(int32, int32) (int32, error)
	Subtraction(int32, int32) (int32, error)
	Multiplication(int32, int32) (int32, error)
}

type Application struct {
	db    ports.Db
	arith Arithmetic
}

func NewApplication(db ports.Db, arith Arithmetic) *Application {
	return &Application{db: db, arith: arith}
}

func (app *Application) GetAddition(a, b int32) (int32, error) {
	answer, err := app.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}
	err = app.db.ToHistory(answer, "addition")
	return answer, err
}

func (app *Application) GetSubtraction(a, b int32) (int32, error) {
	answer, err := app.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}
	err = app.db.ToHistory(answer, "subtraction")
	return answer, err
}

func (app *Application) GetMultiplication(a, b int32) (int32, error) {
	answer, err := app.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	err = app.db.ToHistory(answer, "multiplication")
	return answer, err
}

func (app *Application) GetDivision(a, b int32) (int32, error) {
	answer, err := app.arith.Division(a, b)
	if err != nil {
		return 0, err
	}
	err = app.db.ToHistory(answer, "division")
	return answer, err
}
