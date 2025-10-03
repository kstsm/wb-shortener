package handler

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/kstsm/wb-shortener/internal/apperrors"
	"github.com/kstsm/wb-shortener/internal/dto"
	"github.com/kstsm/wb-shortener/internal/models"
	"github.com/kstsm/wb-shortener/internal/utils"
	"github.com/kstsm/wb-shortener/internal/validation"
	"net/http"
)

func (h Handler) shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.ShortenRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if err := validation.IsValidURL(req.URL); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if req.CustomAlias != "" {
		if err := validation.IsValidShortURL(req.CustomAlias); err != nil {
			utils.WriteError(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	resp, err := h.service.ShortenURL(r.Context(), &req)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrAliasAlreadyExists):
			utils.WriteError(w, http.StatusConflict, "Alias already exists")
		default:
			utils.WriteError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	utils.SendJSON(w, http.StatusCreated, resp)
}

func (h Handler) redirect(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "shortURL")
	if err := validation.IsValidShortURL(shortURL); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	reqInfo := models.RequestInfo{
		UserAgent: r.Header.Get("User-Agent"),
		IP:        utils.GetClientIP(r),
		Referer:   r.Header.Get("Referer"),
	}

	originalURL, err := h.service.Redirect(r.Context(), shortURL, reqInfo)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrNotFound):
			utils.WriteError(w, http.StatusNotFound, "Short URL not found")
		default:
			utils.WriteError(w, http.StatusInternalServerError, "Failed to redirect")
		}
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}

func (h Handler) getAnalytics(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "shortURL")
	if err := validation.IsValidShortURL(shortURL); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	analytics, err := h.service.GetAnalytics(r.Context(), shortURL)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrNotFound):
			utils.WriteError(w, http.StatusNotFound, "Short URL not found")
		default:
			utils.WriteError(w, http.StatusInternalServerError, "Failed to get analytics")
		}
		return
	}

	utils.SendJSON(w, http.StatusOK, analytics)
}
