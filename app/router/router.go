package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"technopartner/app/middlewares"
	_userData "technopartner/features/users/data"
	_userHandler "technopartner/features/users/handler"
	_userService "technopartner/features/users/service"

	_incomeData "technopartner/features/incomes/data"
	_incomeHandler "technopartner/features/incomes/handler"
	_incomeService "technopartner/features/incomes/service"

	_outcomeData "technopartner/features/outcomes/data"
	_outcomeHandler "technopartner/features/outcomes/handler"
	_outcomeService "technopartner/features/outcomes/service"
)

func InitRouter(db *gorm.DB, c *echo.Echo) {
	UserData := _userData.New(db)
	UserService := _userService.New(UserData)
	UserHandlerAPI := _userHandler.New(UserService)

	OutcomeData := _outcomeData.New(db)
	OutcomeService := _outcomeService.New(OutcomeData)
	OutcomeHandlerAPI := _outcomeHandler.New(OutcomeService)

	IncomeData := _incomeData.New(db)
	IncomeService := _incomeService.New(IncomeData)
	IncomeHandlerAPI := _incomeHandler.New(IncomeService)

	//Users
	c.POST("/login", UserHandlerAPI.Login)
	c.POST("/create", UserHandlerAPI.CreateUser)

	//Incomes
	c.POST("/income", IncomeHandlerAPI.CreateIncome, middlewares.JWTMiddleware())
	c.GET("/income", IncomeHandlerAPI.GetAllIncomes, middlewares.JWTMiddleware())
	c.PUT("/income/:id", IncomeHandlerAPI.UpdateIncome, middlewares.JWTMiddleware())
	c.DELETE("/income/:id", IncomeHandlerAPI.DeleteIncome, middlewares.JWTMiddleware())

	//outcomes
	c.POST("/outcome", OutcomeHandlerAPI.CreateOutcome, middlewares.JWTMiddleware())
	c.GET("/outcome", OutcomeHandlerAPI.GetAllOutcomes, middlewares.JWTMiddleware())
	c.PUT("/outcome/:id", OutcomeHandlerAPI.UpdateOutcome, middlewares.JWTMiddleware())
	c.DELETE("/outcome/:id", OutcomeHandlerAPI.DeleteOutcome, middlewares.JWTMiddleware())
}
