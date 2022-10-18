package controller

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/velotio-tech/go-k8s-training/user/utils"
)

func (c *controller) getUserOrders(w http.ResponseWriter, r *http.Request) {
	redirectURL := utils.GetServiceConfig().OrderURL + r.URL.Path
	response, err := http.Get(redirectURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (c *controller) deleteUserOrders(w http.ResponseWriter, r *http.Request) {
	redirectURL := utils.GetServiceConfig().OrderURL + r.URL.Path
	req, err := http.NewRequest(http.MethodDelete, redirectURL, nil)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	client := http.Client{Timeout: 30 * time.Second}
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()
	w.WriteHeader(response.StatusCode)
}

func (c *controller) deleteOrder(w http.ResponseWriter, r *http.Request) {
	redirectURL := utils.GetServiceConfig().OrderURL + r.URL.Path
	req, err := http.NewRequest(http.MethodDelete, redirectURL, nil)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	client := http.Client{Timeout: 30 * time.Second}
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()
	w.WriteHeader(response.StatusCode)
}

func (c *controller) updateOrder(w http.ResponseWriter, r *http.Request) {
	redirectURL := utils.GetServiceConfig().OrderURL + r.URL.Path
	req, err := http.NewRequest(http.MethodPatch, redirectURL, r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	client := http.Client{Timeout: 30 * time.Second}
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(response.StatusCode)
	w.Header().Add("content-type", "application/json")
	w.Write(body)
}

func (c *controller) createUserOrder(w http.ResponseWriter, r *http.Request) {
	redirectURL := utils.GetServiceConfig().OrderURL + r.URL.Path
	req, err := http.NewRequest(http.MethodPost, redirectURL, r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	client := http.Client{Timeout: 30 * time.Second}
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(response.StatusCode)
	w.Header().Add("content-type", "application/json")
	w.Write(body)
}
