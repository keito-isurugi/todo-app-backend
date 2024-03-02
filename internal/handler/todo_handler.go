package handler

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/keito-isurugi/todo-app-backend/internal/domain/entity"
	domain "github.com/keito-isurugi/todo-app-backend/internal/domain/repository"
	usecase "github.com/keito-isurugi/todo-app-backend/internal/usecase/todo"
)

type TodoHandler interface {
	ListTodos(c echo.Context) error
	RegisterTodo(c echo.Context) error
	ChangeTodo(c echo.Context) error
}

type todoHandler struct {
	todoRepo  domain.Todo
	zapLogger *zap.Logger
}

func NewTodoHandler(todoRepo domain.Todo, zapLogger *zap.Logger) TodoHandler {
	return &todoHandler{
		todoRepo:  todoRepo,
		zapLogger: zapLogger,
	}
}

type todoMaster struct {
	ID    int    `json:"id" example:"1"`
	Title string `json:"title" example:"タイトル"`
}

type todoResponse struct {
	ID       int    `json:"id" example:"1"`
	Title    string `json:"title" example:"テストタイトル"`
	DoneFlag bool   `json:"done_flag" example:"false"`
}

type listTodosResponse []todoResponse

// ListTodos
// @Summary		Todo一覧
// @Description	Todo一覧
// @id ListTodos
// @tags		todo
// @Accept		json
// @Produce		json
// @Success		200	{object}	listTodosResponse
// @Failure		400	{object}	fieldError
// @Failure		401	{object}	errResponse
// @Failure		403	{object}	errResponse
// @Failure		500	{object}	errResponse
// @Router		/todos [get]
func (h *todoHandler) ListTodos(c echo.Context) error {
	traceID := c.Get("trace_id").(string)

	uc := usecase.NewListTodosUsecase(h.todoRepo)
	ms, err := uc.Exec(c.Request().Context())
	if err != nil {
		res := createErrResponse(err)
		res.outputErrorLog(h.zapLogger, err.Error(), traceID)
		return c.JSON(res.Status, res)
	}

	res := make(listTodosResponse, len(ms))
	for i := range ms {
		m := ms[i]
		if err = copier.Copy(&res[i], &m); err != nil {
			res := createErrResponse(err)
			res.outputErrorLog(h.zapLogger, err.Error(), traceID)
			return c.JSON(res.Status, res)
		}
	}

	return c.JSON(http.StatusOK, res)
}

type registerTodoRequest struct {
	Title string `json:"title" example:"サンプルタイトル" ja:"タイトル" validate:"required,max=255"`
}

// RegisterTodo
// @Summary		Todo登録
// @Description	Todoを登録
// @id			RegisterTodo
// @tags		todo
// @Accept		json
// @Produce		json
// @Success		201			{object}	createdResponse
// @Failure		400			{object}	fieldError
// @Failure		401			{object}	errResponse
// @Failure		403			{object}	errResponse
// @Failure		500			{object}	errResponse
// @Router		/todos [post]
// @Param request body registerTodoRequest true "registerTodoRequest"
func (m *todoHandler) RegisterTodo(c echo.Context) error {
	traceID := c.Get("trace_id").(string)

	var req registerTodoRequest
	if err := c.Bind(&req); err != nil {
		res := createErrResponse(err)
		res.outputErrorLog(m.zapLogger, err.Error(), traceID)
		return c.JSON(res.Status, res)
	}
	if err := validate.Struct(req); err != nil {
		m.zapLogger.Error(err.Error())
		return c.JSON(http.StatusBadRequest, fieldErrors(err))
	}

	uc := usecase.NewRegisterTodoUsecase(m.todoRepo)
	param := entity.NewRegisterTodo(req.Title)

	id, err := uc.Exec(c.Request().Context(), param)
	if err != nil {
		res := createErrResponse(err)
		res.outputErrorLog(m.zapLogger, err.Error(), traceID)
		return c.JSON(res.Status, res)
	}

	return c.JSON(http.StatusCreated, newCreatedResponse(id))
}

type changeTodoRequest struct {
	TodoID int    `json:"todo_id" example:"1" ja:"TodoID" validate:"required"`
	Title  string `json:"title" example:"サンプルタイトル" ja:"タイトル" validate:"required,max=255"`
}

// ChangeTodo
// @Summary		Todo編集
// @Description	Todo編集
// @id ChangeTodo
// @tags		todo
// @Accept		json
// @Produce		json
// @Param		request	body	changeTodoRequest	true	"changeTodoRequest"
// @Success		202	{object}	emptyResponse
// @Failure		400	{object}	fieldError
// @Failure		401	{object}	errResponse
// @Failure		403	{object}	errResponse
// @Failure		404	{object}	errResponse
// @Failure		500	{object}	errResponse
// @Router		/todos/{id} [put]
// @Param		id						        path	string	true	"1"
func (m *todoHandler) ChangeTodo(c echo.Context) error {
	traceID := c.Get("trace_id").(string)
	id, _ := strconv.Atoi(c.Param("id"))
	title := c.Param("title")

	var req changeTodoRequest
	if err := c.Bind(&req); err != nil {
		res := createErrResponse(err)
		res.outputErrorLog(m.zapLogger, err.Error(), traceID)
		return c.JSON(res.Status, res)
	}

	if err := validate.Struct(req); err != nil {
		m.zapLogger.Error(err.Error())
		return c.JSON(http.StatusBadRequest, fieldErrors(err))
	}

	param := entity.NewChangeTodo(
		id,
		title,
	)

	uc := usecase.NewChangeTodoUsecase(m.todoRepo)
	if err := uc.Exec(c.Request().Context(), param); err != nil {
		res := createErrResponse(err)
		res.outputErrorLog(m.zapLogger, err.Error(), traceID)
		return c.JSON(res.Status, res)
	}
	return c.JSON(http.StatusAccepted, emptyResponse{})
}
