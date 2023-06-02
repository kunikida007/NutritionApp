package handler

import (
	"net/http"
	"strconv"

	"kunikida123456/NutritionApp/usecase"
	"kunikida123456/NutritionApp/util"

	"github.com/labstack/echo/v4"
)

// MealHandler meal handlerのinterface
type MealHandler interface {
	Post(c echo.Context) error
	Get(c echo.Context) error
	GetAll(c echo.Context) error
	Put(c echo.Context) error
	Delete(c echo.Context) error
}

type mealHandler struct {
	mealUsecase usecase.MealUsecase
}

func NewMealHandler(mealUsecase usecase.MealUsecase) MealHandler {
	return &mealHandler{mealUsecase: mealUsecase}
}

type postMealRequest struct {
	Memo     string  `json:"memo" `
	MealType string  `json:"mealtype" `
	Carbs    float64 `json:"carbs"`
	Fat      float64 `json:"fat"`
	Protein  float64 `json:"protein"`
	Calories float64 `json:"calories"`
}

type postMealResponse struct {
	ID       int     `json:"id"`
	Memo     string  `json:"memo" `
	MealType string  `json:"mealtype" `
	Carbs    float64 `json:"carbs"`
	Fat      float64 `json:"fat"`
	Protein  float64 `json:"protein"`
	Calories float64 `json:"calories"`
}

// Postmeal godoc
// @ID postMeals
// @Description 食事を追加する
// @Accept   json
// @Produce  json
// @Router   /meals [post]
// @Param    body body postMealRequest true "食事情報"
// @Success  201 {object} postMealResponse
// @Failure  400
// @Failure  401
// @Failure  500
func (mh *mealHandler) Post(c echo.Context) error {
	req := new(postMealRequest)

	if err := c.Bind(&req); err != nil {
		return err
	}

	userID := util.GetUserIDFromToken(c)
	createdMeal, err := mh.mealUsecase.Create(userID, req.Memo, req.MealType, req.Carbs, req.Fat, req.Protein, req.Calories)
	if err != nil {
		return err
	}

	res := postMealResponse{
		ID:       createdMeal.ID,
		Memo:     createdMeal.Memo,
		MealType: createdMeal.MealType,
		Carbs:    createdMeal.Carbs,
		Fat:      createdMeal.Fat,
		Protein:  createdMeal.Protein,
		Calories: createdMeal.Calories,
	}

	return c.JSON(http.StatusCreated, res)
}

// Getmeals godoc
// @ID getMealsId
// @Description 食事を追加する
// @Accept   json
// @Produce  json
// @Router   /meals/{id} [get]
// @Success  200 {object} getMealResponse
// @Failure  400
// @Failure  401
// @Failure  404
// @Failure 500
type getMealsResponse struct {
	ID       int     `json:"id"`
	Memo     string  `json:"memo"`
	MealType string  `json:"mealtype"`
	Carbs    float64 `json:"carbs"`
	Fat      float64 `json:"fat"`
	Protein  float64 `json:"protein"`
	Calories float64 `json:"calories"`
}

// Getmeal godoc
// @ID getMealsId
// @Description 食事を取得する
// @Accept   json
// @Produce  json
// @Router   /meals/{id} [get]
// @Success  200 {object} getMealsResponse
// @Failure  400
// @Failure  401
// @Failure  404
// @Failure 500
func (mh *mealHandler) Get(c echo.Context) error {
	id, err := strconv.Atoi((c.Param("id")))
	if err != nil {
		return err
	}

	foundMeal, err := mh.mealUsecase.FindByID(c, id)
	if err != nil {
		return err
	}
	res := getMealsResponse{
		ID:       foundMeal.ID,
		Memo:     foundMeal.Memo,
		MealType: foundMeal.MealType,
		Carbs:    foundMeal.Carbs,
		Fat:      foundMeal.Fat,
		Protein:  foundMeal.Protein,
		Calories: foundMeal.Calories,
	}
	return c.JSON(http.StatusOK, res)

}

type putMealsRequest struct {
	Memo     string  `json:"memo" `
	MealType string  `json:"mealtype" `
	Carbs    float64 `json:"carbs"`
	Fat      float64 `json:"fat"`
	Protein  float64 `json:"protein"`
	Calories float64 `json:"calories"`
}

type putMealsResponse struct {
	ID       int     `json:"id"`
	Memo     string  `json:"memo"`
	MealType string  `json:"mealtype"`
	Carbs    float64 `json:"carbs"`
	Fat      float64 `json:"fat"`
	Protein  float64 `json:"protein"`
	Calories float64 `json:"calories"`
}

// Putmeal godoc
// @ID putMealsId
// @Description 食事を更新する
// @Accept   json
// @Produce  json
// @Router   /meals/{id} [put]
// @Param    body body putMealsRequest true "食事情報"
// @Success  200 {object} putMealsResponse
// @Failure 400
// @Failure 401
// @Failure  404
// @Failure 500
func (mh *mealHandler) Put(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	req := new(putMealsRequest)

	if err := c.Bind(&req); err != nil {
		return err
	}

	updatedMeal, err := mh.mealUsecase.Update(c, id, req.Memo, req.MealType, req.Carbs, req.Fat, req.Protein, req.Calories)
	if err != nil {
		return err
	}

	res := putMealsResponse{
		ID:       updatedMeal.ID,
		Memo:     updatedMeal.Memo,
		MealType: updatedMeal.MealType,
		Carbs:    updatedMeal.Carbs,
		Fat:      updatedMeal.Fat,
		Protein:  updatedMeal.Protein,
		Calories: updatedMeal.Calories,
	}

	return c.JSON(http.StatusOK, res)
}

// Deletemeal godoc
// @ID deleteMealsId
// @Description 食事を消去する
// @Accept   json
// @Produce  json
// @Router   /meals/{id} [delete]
// @Success  204
// @Failure  400
// @Failure  401
// @Failure  404
// @Failure  500
func (mh *mealHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	err = mh.mealUsecase.Delete(c, id)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

// GetAllmeals godoc
// @ID getMeals
// @Description すべての食事を取得
// @Accept   json
// @Produce  json
// @Router   /meals [get]
// @Success  200 {array} []getMealsResponse
// @Failure  400
// @Failure  401
// @Failure  404
// @Failure 500
func (mh *mealHandler) GetAll(c echo.Context) error {
	foundMeals, err := mh.mealUsecase.FindAll(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, foundMeals)

}
