package err

import "errors"

var (
	//ErrRequiredPorfolio error message
	ErrRequiredPorfolio = errors.New("Portfolio es requerido")
	//ErrInvalidInitDate error message
	ErrInvalidInitDate = errors.New("Fecha de inicio incorrecta")
	//ErrRequiredInitDate error message
	ErrRequiredInitDate = errors.New("Fecha de inicio es requerida")
	//ErrInvalidFinishDate error message
	ErrInvalidFinishDate = errors.New("Fecha de fin es incorrecta")
	//ErrRequiredFinishDate error message
	ErrRequiredFinishDate = errors.New("Fecha de fin es requerida")
	//ErrStockNotFound error message
	ErrStockNotFound = errors.New("Stock no encontrada")
	//ErrDateRange error message
	ErrDateRange = errors.New("Error obteniendo rango de fechas")
	//ErrGetStockPrice error message
	ErrGetStockPrice = errors.New("Error obteniendo el precio del stock")
)
