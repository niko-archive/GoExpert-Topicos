package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.dev/nicolasmmb/GoExpert-Topicos/internal/dto"
	"github.dev/nicolasmmb/GoExpert-Topicos/internal/entity"
	"github.dev/nicolasmmb/GoExpert-Topicos/internal/infra/database"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{ProductDB: db}
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.ProductDB.Create(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	productOutput := dto.ProductOutputId{
		ID: p.ID.String(),
	}

	json.NewEncoder(w).Encode(productOutput)
	w.WriteHeader(http.StatusCreated)

}

func (h *ProductHandler) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	p, err := h.ProductDB.FindById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product := dto.ProductOutput{
		ID:        p.ID.String(),
		Name:      p.Name,
		Price:     p.Price,
		CreatedAt: p.CreatedAt.Format(time.RFC3339),
	}

	json.NewEncoder(w).Encode(product)
	w.WriteHeader(http.StatusOK)

}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	var product dto.UpdateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p, err = h.ProductDB.FindById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.Name = product.Name
	p.Price = product.Price

	err = h.ProductDB.Update(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	productOutput := dto.ProductOutputId{
		ID: p.ID.String(),
	}

	json.NewEncoder(w).Encode(productOutput)
	w.WriteHeader(http.StatusOK)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	err := h.ProductDB.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	productOutput := dto.ProductOutputId{
		ID: id,
	}

	json.NewEncoder(w).Encode(productOutput)
	w.WriteHeader(http.StatusOK)
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if sort == "" {
		sort = "asc"
	}
	if sort != "asc" && sort != "desc" {
		http.Error(w, "sort must be asc or desc", http.StatusBadRequest)
		return
	}

	products, err := h.ProductDB.FindAll(intPage, intLimit, sort)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var productsOutput []dto.ProductOutput

	for _, p := range products {
		productsOutput = append(productsOutput, dto.ProductOutput{
			ID:        p.ID.String(),
			Name:      p.Name,
			Price:     p.Price,
			CreatedAt: p.CreatedAt.Format(time.RFC3339),
		})
	}

	json.NewEncoder(w).Encode(productsOutput)
	w.WriteHeader(http.StatusOK)
}
