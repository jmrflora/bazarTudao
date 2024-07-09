package handler

import (
	"net/http"
	"strconv"

	"github.com/jmrflora/bazarTudao/crud"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	crud *crud.Crud
}

func NewHandler(crud *crud.Crud) *Handler {
	return &Handler{crud: crud}
}

func (h *Handler) HandleGetOrdens(c echo.Context) error {
	ordens, err := h.crud.GetAllOrdens()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ordens)
}

func (h *Handler) HandleGetOrdensCompletas(c echo.Context) error {
	ordens, err := h.crud.GetOrdensCompletas()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ordens)
}

func (h *Handler) HandleGetOrdensIntocadas(c echo.Context) error {
	ordens, err := h.crud.GetOrdensIntocadas()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ordens)
}

func (h *Handler) HandleGetOrdensParciais(c echo.Context) error {
	ordens, err := h.crud.GetOrdensParciais()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ordens)
}

func (h *Handler) HandleGetEnvios(c echo.Context) error {
	envios, err := h.crud.GetAllEnvios()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, envios)
}

func (h *Handler) HandleGetItensPorIdOrdem(c echo.Context) error {
	idst := c.Param("id")
	id, err := strconv.Atoi(idst)
	if err != nil {
		return err
	}
	itens, err := h.crud.GetItensPorIdOrdem(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, itens)
}

func (h *Handler) HandleGetProdutos(c echo.Context) error {
	produtos, err := h.crud.GetAllProdutos()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, produtos)
}

func (h *Handler) HandlerGetAllClientes(c echo.Context) error {
	clientes, err := h.crud.GetAllClientes()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, clientes)
}
