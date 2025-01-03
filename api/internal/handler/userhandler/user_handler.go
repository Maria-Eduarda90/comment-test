package userhandler

import (
	"api/internal/common/utils"
	"api/internal/dto"
	"api/internal/handler/httperr"
	"api/internal/handler/validation"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserDto

    if r.Body == http.NoBody {
        slog.Error("Body is empty", slog.String("package", "userhandler"))
        w.WriteHeader(http.StatusBadRequest)
        msg := httperr.NewBadRequestError("body is required")
        json.NewEncoder(w).Encode(msg)
        return
    }

    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        slog.Error("error to decode body", "err", err, slog.String("package", "userhandler"))
        w.WriteHeader(http.StatusBadRequest)
        msg := httperr.NewBadRequestError("error to decode body")
        json.NewEncoder(w).Encode(msg)
        
        return
    }

    httpErr := validation.ValidateHttpData(req)
    if httpErr != nil {
        slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "userhandler"))
        w.WriteHeader(httpErr.Code)
        json.NewEncoder(w).Encode(httpErr)
        return
    }

    err = h.service.CreateUser(r.Context(), req)
    if err != nil {
        slog.Error(fmt.Sprintf("error to create user: %v", err), slog.String("package", "userhandler"))
        w.WriteHeader(http.StatusInternalServerError)
        msg := httperr.NewBadRequestError("error to create user")
        json.NewEncoder(w).Encode(msg)
        return
    }
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
    var req dto.UpdateUserDto;

    user, err := utils.DecodeJwt(r)
    if err != nil {
      slog.Error("error to decode jwt", slog.String("package", "userhandler"))
      w.WriteHeader(http.StatusBadRequest)
      msg := httperr.NewBadRequestError("error to decode jwt")
      json.NewEncoder(w).Encode(msg)
      return
    }

    if r.Body == http.NoBody {
        slog.Error("body is empty", slog.String("package", "userhandler"))
        w.WriteHeader(http.StatusBadRequest)
        msg := httperr.NewBadRequestError("body is required")
        json.NewEncoder(w).Encode(msg)
        return
    }

    err = json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        slog.Error("error to decode body", "err", err, slog.String("package", "handler_user"))
        w.WriteHeader(http.StatusBadRequest)
        msg := httperr.NewBadRequestError("error to decode body")
        json.NewEncoder(w).Encode(msg)
        return
    }
    
    httpErr := validation.ValidateHttpData(req)
    if httpErr != nil {
        slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "handler_user"))
        w.WriteHeader(httpErr.Code)
        json.NewEncoder(w).Encode(httpErr)
        return
    }

    err = h.service.UpdateUser(r.Context(), req, user.ID)
    if err != nil {
        slog.Error(fmt.Sprintf("error to update user: %v", err), slog.String("package", "handler_user"))
        if err.Error() == "user not found" {
          w.WriteHeader(http.StatusNotFound)
          msg := httperr.NewNotFoundError("user not found")
          json.NewEncoder(w).Encode(msg)
          return
        }
        w.WriteHeader(http.StatusInternalServerError)
        msg := httperr.NewBadRequestError("error to update user")
        json.NewEncoder(w).Encode(msg)
        return
    }
}

func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
    user, err := utils.DecodeJwt(r)
    if err != nil {
      slog.Error("error to decode jwt", slog.String("package", "userhandler"))
      w.WriteHeader(http.StatusBadRequest)
      msg := httperr.NewBadRequestError("error to decode jwt")
      json.NewEncoder(w).Encode(msg)
      return
    }
    res, err := h.service.GetUserByID(r.Context(), user.ID)
    if err != nil {
      slog.Error(fmt.Sprintf("error to get user: %v", err), slog.String("package", "userhandler"))
      if err.Error() == "user not found" {
        w.WriteHeader(http.StatusNotFound)
        msg := httperr.NewNotFoundError("user not found")
        json.NewEncoder(w).Encode(msg)
        return
      }
      w.WriteHeader(http.StatusInternalServerError)
      msg := httperr.NewBadRequestError("error to get user")
      json.NewEncoder(w).Encode(msg)
      return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(res)
}

func (h* handler) DeleteUser(w http.ResponseWriter, r *http.Request){
    
    user, err := utils.DecodeJwt(r)
    if err != nil {
      slog.Error("error to decode jwt", slog.String("package", "userhandler"))
      w.WriteHeader(http.StatusBadRequest)
      msg := httperr.NewBadRequestError("error to decode jwt")
      json.NewEncoder(w).Encode(msg)
      return
    }

    err = h.service.DeleteUser(r.Context(), user.ID)
    if err != nil {
      slog.Error(fmt.Sprintf("error to delete user: %v", err), slog.String("package", "handler_user"))
      if err.Error() == "user not found" {
        w.WriteHeader(http.StatusNotFound)
        msg := httperr.NewNotFoundError("user not found")
        json.NewEncoder(w).Encode(msg)
        return
      }
      w.WriteHeader(http.StatusInternalServerError)
      msg := httperr.NewBadRequestError("error to delete user")
      json.NewEncoder(w).Encode(msg)
      return
    }
    
    w.WriteHeader(http.StatusNoContent)
}

func (h *handler) FindManyUsers(w http.ResponseWriter, r *http.Request) {
    res, err := h.service.FindManyUsers(r.Context())
    if err != nil {
      slog.Error(fmt.Sprintf("error to find many users: %v", err), slog.String("package", "handler_user"))
      w.WriteHeader(http.StatusInternalServerError)
      msg := httperr.NewBadRequestError("error to find many users")
      json.NewEncoder(w).Encode(msg)
      return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(res)
}

func (h *handler) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
    var req dto.UpdateUserPasswordDto

    user, err := utils.DecodeJwt(r)
    if err != nil {
      slog.Error("error to decode jwt", slog.String("package", "userhandler"))
      w.WriteHeader(http.StatusBadRequest)
      msg := httperr.NewBadRequestError("error to decode jwt")
      json.NewEncoder(w).Encode(msg)
      return
    }
    if r.Body == http.NoBody {
      slog.Error("body is empty", slog.String("package", "userhandler"))
      w.WriteHeader(http.StatusBadRequest)
      msg := httperr.NewBadRequestError("body is required")
      json.NewEncoder(w).Encode(msg)
      return
    }
    err = json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
      slog.Error("error to decode body", "err", err, slog.String("package", "handler_user"))
      w.WriteHeader(http.StatusBadRequest)
      msg := httperr.NewBadRequestError("error to decode body")
      json.NewEncoder(w).Encode(msg)
      return
    }
    httpErr := validation.ValidateHttpData(req)
    if httpErr != nil {
      slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "handler_user"))
      w.WriteHeader(httpErr.Code)
      json.NewEncoder(w).Encode(httpErr)
      return
    }
    err = h.service.UpdateUserPassword(r.Context(), &req, user.ID)
    if err != nil {
      slog.Error(fmt.Sprintf("error to update user password: %v", err), slog.String("package", "handler_user"))
      if err.Error() == "user not found" {
        w.WriteHeader(http.StatusNotFound)
        msg := httperr.NewNotFoundError("user not found")
        json.NewEncoder(w).Encode(msg)
        return
      }
      w.WriteHeader(http.StatusInternalServerError)
      msg := httperr.NewBadRequestError("error to update user password")
      json.NewEncoder(w).Encode(msg)
      return
    }
}