package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	e "shop/internal/entity/product"
	server "shop/server"
	"strconv"

	"time"

	"github.com/gorilla/mux"
)

type productUsecase interface {
	GetProductById(ctx context.Context, id int64) (e.Product, error)
	GetProductAll(ctx context.Context) ([]e.Product, error)
	UpdateProduct(ctx context.Context, id int64, product e.Product) (e.Product, error)
	DeleteProduct(ctx context.Context, id int64) error
	AddProduct(ctx context.Context, product e.Product) (e.Product, error)
}

type Handler struct {
	productUc productUsecase
}

func New(productUc productUsecase) *Handler {
	return &Handler{
		productUc: productUc,
	}
}

type AddProductResponse struct {
	ID int64 `json:"id"`
}

func (p *Handler) RootHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Entering RootHandler")
	fmt.Fprintf(w, "Hello Devcamp-2022-snd!")
}

func (h *Handler) AddProductHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering AddProductHandler")
	timeStart := time.Now()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[ProductHandler][AddProduct] unable to read body, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	var data e.Product
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("[ProductHandler][AddProduct] unable to unmarshal json, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

}

func (p *Handler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering GetProduct Handler")
	timeStart := time.Now()

	vars := mux.Vars(r)
	queryID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Println("[ProductHandler][GetProduct] bad request, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	resp, err := p.productUc.GetProductById(context.Background(), queryID)
	if err != nil {
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	server.RenderResponse(w, http.StatusOK, resp, timeStart)
	return
}

func (p *Handler) GetProductAllHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering GetProductAll Handler")
	timeStart := time.Now()
	var err error

	resp, err := p.productUc.GetProductAll(context.Background())
	if err != nil {
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	server.RenderResponse(w, http.StatusOK, resp, timeStart)
	return
}

func (p *Handler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering UpdateProduct Handler")
	timeStart := time.Now()
	vars := mux.Vars(r)
	queryID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Println("[ProductHandler][UpdateProduct] bad request, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[ProductHandler][UpdateProduct] unable to read body, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	var data e.Product
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("[ProductHandler][UpdateProduct] unable to unmarshal json, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	resp, err := p.productUc.UpdateProduct(context.Background(), queryID, data)
	if err != nil {
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	server.RenderResponse(w, http.StatusCreated, resp, timeStart)
	return
}

func (p *Handler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering GetProduct Handler")
	timeStart := time.Now()

	vars := mux.Vars(r)
	queryID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Println("[ProductHandler][DeleteProduct] bad request, err: ", err.Error())
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	err = p.productUc.DeleteProduct(context.Background(), queryID)
	if err != nil {
		server.RenderError(w, http.StatusBadRequest, err, timeStart)
		return
	}

	server.RenderResponse(w, http.StatusOK, "deleted", timeStart)
	return
}
