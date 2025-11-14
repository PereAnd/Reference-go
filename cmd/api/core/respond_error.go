// Package core proporciona utilidades y ayudantes compartidos para manejadores HTTP.
// Este paquete contiene funcionalidad común usada en diferentes paquetes de handlers,
// como el formateo de respuestas de error y el mapeo de códigos de estado HTTP.
package core

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-l/pkg/domain"
)

// InternalServerErrorMessage es el mensaje de error genérico retornado a los clientes
// cuando ocurre un error interno inesperado del servidor. Este mensaje se usa para evitar
// exponer detalles sensibles de implementación a los consumidores de la API.
const InternalServerErrorMessage = "Ooops! Something went wrong. Please help us by reporting this issue."

// RespondError formatea y envía una respuesta de error al cliente HTTP.
// Verifica si el error es un domain.AppError y lo mapea al código de estado HTTP
// apropiado usando ErrCodeMapping. Si el error no es un AppError o no tiene un
// código de estado mapeado, retorna un 500 Internal Server Error con un mensaje
// genérico para evitar exponer detalles de implementación.
//
// Parámetros:
//   - c: El contexto de Gin para la petición HTTP
//   - err: El error a formatear y enviar como respuesta
func RespondError(c *gin.Context, err error) {
	c.Header("Content-Type", "application/json")
	var appErr domain.AppError
	if errors.As(err, &appErr) {
		if status, ok := ErrCodeMapping[appErr.Code]; ok {
			c.JSON(status, appErr)
			return
		}
	}
	c.JSON(http.StatusInternalServerError, domain.AppError{Code: domain.ErrCodeInternalServerError, Msg: InternalServerErrorMessage})
}

// ErrCodeMapping mapea códigos de error del dominio a sus códigos de estado HTTP correspondientes.
// Este mapeo es usado por RespondError para determinar el código de estado HTTP
// apropiado para errores específicos del dominio.
var ErrCodeMapping map[string]int = map[string]int{
	domain.ErrCodeDuplicateKey:  http.StatusConflict,    // 409
	domain.ErrCodeNotFound:       http.StatusNotFound,    // 404
	domain.ErrCodeInvalidParams:  http.StatusBadRequest,  // 400
}
