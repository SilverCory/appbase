package apperr

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/rs/zerolog"
)

func HandleError(ctx gin.Context, err error) {
	if err == nil {
		return
	}

	var appErr AppError
	if !errors.As(err, &appErr) {
		appErr = New(http.StatusInternalServerError, err)
	} else {
		appErr = appErr.WithData("_originalError", appErr)
	}

	render(ctx, appErr)
}

func render(ctx gin.Context, appErr AppError) {

	// Status
	var statusCode = appErr.StatusCode
	if statusCode <= 0 {
		statusCode = http.StatusInternalServerError
	}
	if statusCode >= 500 {
		LogErr(ctx, appErr, "")
	}

	ctx.JSON(statusCode, cloneAppErr(appErr, false))
}

func LogErr(ctx gin.Context, err error, msg string) {
	if err == nil {
		return
	}

	erEv := zerolog.Ctx(ctx.Request.Context()).Error()

	// Make AppError's marshal json.
	switch err.(type) {
	case AppError:
		erEv = erEv.Interface("error", err)
	default:
		erEv = erEv.Err(err)
	}

	// TODO add context data here

	erEv.Msg(msg)
}

func cloneAppErr(appErr AppError, showSecrets bool) AppError {
	var newAppErr = New(appErr.StatusCode, appErr)
	for k, v := range appErr.Data {
		if strings.HasPrefix(k, "_") && !showSecrets {
			continue
		}

		newAppErr = newAppErr.WithData(k, v)
	}

	return newAppErr
}
