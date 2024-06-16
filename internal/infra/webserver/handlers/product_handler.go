package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/PCPedroso/pos-fc-apis/internal/dto"
	"github.com/PCPedroso/pos-fc-apis/internal/entity"
	"github.com/PCPedroso/pos-fc-apis/internal/infra/database"
	pkg_entity "github.com/PCPedroso/pos-fc-apis/pkg/entity"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{ProductDB: db}
}

// Create Product godoc
// @Summary 	Create Product
// @Description Create Product
// @Tags 		products
// @Accept		json
// @Produce		json
// @Param		request	body dto.CreateProductInput true "product request"
// @Success		201
// @Failure		500 {object} Error
// @Router		/products [post]
// @Security ApiKeyAuth
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Get Product godoc
// @Summary 	Get Product
// @Description Get Product
// @Tags 		products
// @Accept		json
// @Produce		json
// @Param		id		path 		string	true	"product ID"	Format(uuid)
// @Success		200		{object}	entity.Product
// @Failure		404
// @Failure		500 	{object} 	Error
// @Router		/products/{id} [get]
// @Security ApiKeyAuth
func (h *ProductHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// Update Product godoc
// @Summary 	Update Product
// @Description Update Product
// @Tags 		products
// @Accept		json
// @Produce		json
// @Param		id		path 		string					true	"product ID"	Format(uuid)
// @Param		param	body 		dto.CreateProductInput	true	"product request"
// @Success		200
// @Failure		404
// @Failure		500 	{object} 	Error
// @Router		/products/{id}		[put]
// @Security ApiKeyAuth
func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.ID, err = pkg_entity.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

// Delete Product godoc
// @Summary 	Delete Product
// @Description Delete Product
// @Tags 		products
// @Accept		json
// @Produce		json
// @Param		id		path 		string					true	"product ID"	Format(uuid)
// @Success		200
// @Failure		404
// @Failure		500 	{object} 	Error
// @Router		/products/{id}		[delete]
// @Security ApiKeyAuth
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// List Products godoc
// @Summary 	List Products
// @Description get all Products
// @Tags 		products
// @Accept		json
// @Produce		json
// @Param		page	query 		string	false	"page number"
// @Param		limit	query 		string	false	"limit"
// @Success		200		{array}		entity.Product
// @Failure		404
// @Failure		500 	{object} 	Error
// @Router		/products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	page, limit := 0, 0
	sort := r.URL.Query().Get("sort")

	if p, err := strconv.Atoi(r.URL.Query().Get("page")); err == nil {
		page = p
	}

	if l, err := strconv.Atoi(r.URL.Query().Get("limit")); err == nil {
		limit = l
	}

	products, err := h.ProductDB.FindAll(page, limit, sort)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
