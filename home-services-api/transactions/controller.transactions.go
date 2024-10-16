package transactions

import (
	"errors"
	"net/http"

	"github.com/dot-slash-ann/home-services-api/lib"
	"github.com/dot-slash-ann/home-services-api/lib/httpErrors"
	"github.com/gin-gonic/gin"
)

type TransactionsController struct {
	transactionsService TransactionsService
}

func NewTransactionsController(transactionsService TransactionsService) *TransactionsController {
	return &TransactionsController{
		transactionsService: transactionsService,
	}
}

func (controller *TransactionsController) Create(c *gin.Context) {
	var createTransactionDto CreateTransactionDto

	if !lib.HandleDecodeTime(c, &createTransactionDto) {
		return
	}

	transaction, err := controller.transactionsService.Create(createTransactionDto)

	if err != nil {
		httpErr := httpErrors.BadRequestError(err, nil)

		c.Error((httpErr))

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": TransactionToJson(transaction),
	})
}

func (controller *TransactionsController) FindAll(c *gin.Context) {
	filters := make(map[string]string)

	if tags := c.Query("tags"); tags != "" {
		filters["tags"] = tags
	}

	if categoryName := c.Query("category_name"); categoryName != "" {
		filters["categoryName"] = categoryName
	}

	if categoryID := c.Query("category_id"); categoryID != "" {
		filters["categoryID"] = categoryID
	}

	if minAmount := c.Query("min"); minAmount != "" {
		filters["min"] = minAmount
	}

	if maxAmount := c.Query("max"); maxAmount != "" {
		filters["max"] = maxAmount
	}

	if vendorName := c.Query("vendor_name"); vendorName != "" {
		filters["vendorName"] = vendorName
	}

	if vendorID := c.Query("vendor_id"); vendorID != "" {
		filters["vendorID"] = vendorID
	}

	if transactionOnFrom := c.Query("transaction_on_from"); transactionOnFrom != "" {
		filters["transactionOnFrom"] = transactionOnFrom
	}

	if transactionOnTo := c.Query("transaction_on_to"); transactionOnTo != "" {
		filters["transactionOnTo"] = transactionOnTo
	}

	if postedOnFrom := c.Query("posted_on_from"); postedOnFrom != "" {
		filters["postedOnFrom"] = postedOnFrom
	}

	if postedOnTo := c.Query("posted_on_to"); postedOnTo != "" {
		filters["postedOnTo"] = postedOnTo
	}

	transactionsList, err := controller.transactionsService.FindAll(filters)

	if err != nil && err.Error() != "record not found" {
		httpErr := httpErrors.InternalServerError(err, nil)

		c.Error(httpErr)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ManyTransactionsToJson(transactionsList),
	})
}

func (controller *TransactionsController) FindOne(c *gin.Context) {
	id, found := c.Params.Get("id")

	if !found {
		httpErr := httpErrors.BadRequestError(errors.New("id arg not found"), nil)

		c.Error(httpErr)

		return
	}

	if !lib.IsNumeric(id) {
		httpErr := httpErrors.BadRequestError(errors.New("id must be an integer"), nil)

		c.Error(httpErr)

		return
	}
	transaction, err := controller.transactionsService.FindOne(id)

	if err != nil {
		httpErr := httpErrors.NotFoundError(err, nil)

		c.Error(httpErr)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": TransactionToJson(transaction),
	})
}

func (controller *TransactionsController) Update(c *gin.Context) {
	id, found := c.Params.Get("id")

	if !found {
		httpErr := httpErrors.BadRequestError(errors.New("id arg not found"), nil)

		c.Error(httpErr)

		return
	}

	if !lib.IsNumeric(id) {
		httpErr := httpErrors.BadRequestError(errors.New("id must be an integer"), nil)

		c.Error(httpErr)

		return
	}

	var updateTransactionDto UpdateTransactionDto

	if !lib.HandleDecodeTime(c, &updateTransactionDto) {
		return
	}

	transaction, err := controller.transactionsService.Update(id, updateTransactionDto)

	if err != nil {
		lib.HandleError(c, http.StatusNotFound, err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": TransactionToJson(transaction),
	})
}

func (controller *TransactionsController) Delete(c *gin.Context) {
	id, found := c.Params.Get("id")

	if !found {
		httpErr := httpErrors.BadRequestError(errors.New("id arg not found"), nil)

		c.Error(httpErr)

		return
	}

	if !lib.IsNumeric(id) {
		httpErr := httpErrors.BadRequestError(errors.New("id must be an integer"), nil)

		c.Error(httpErr)

		return
	}

	transaction, err := controller.transactionsService.Delete(id)

	if err != nil {
		lib.HandleError(c, http.StatusNotFound, err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": TransactionToJson(transaction),
	})
}

func (controller *TransactionsController) TagTransaction(c *gin.Context) {
	id, found := c.Params.Get("id")

	if !found {
		httpErr := httpErrors.BadRequestError(errors.New("id arg not found"), nil)

		c.Error(httpErr)

		return
	}

	if !lib.IsNumeric(id) {
		httpErr := httpErrors.BadRequestError(errors.New("id must be an integer"), nil)

		c.Error(httpErr)

		return
	}

	var tagTransactionDto TagTransactionDto

	if err := c.ShouldBind(&tagTransactionDto); err != nil {
		httpErr := httpErrors.BadRequestError(err, nil)

		c.Error(httpErr)

		return
	}

	transaction, err := controller.transactionsService.TagTransaction(tagTransactionDto, id)

	if err != nil {
		httpErr := httpErrors.InternalServerError(err, nil)

		c.Error(httpErr)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": TransactionToJson(transaction),
	})
}
