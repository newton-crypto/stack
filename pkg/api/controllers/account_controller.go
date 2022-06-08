package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/numary/ledger/pkg/core"
	"github.com/numary/ledger/pkg/ledger"
	"github.com/numary/ledger/pkg/ledger/query"
)

type AccountController struct{}

func NewAccountController() AccountController {
	return AccountController{}
}

func (ctl *AccountController) CountAccounts(c *gin.Context) {
	l, _ := c.Get("ledger")

	count, err := l.(*ledger.Ledger).CountAccounts(
		c.Request.Context(),
		query.Address(c.Query("address")),
		query.Metadata(c.QueryMap("metadata")),
	)
	if err != nil {
		ResponseError(c, err)
		return
	}

	c.Header("Count", fmt.Sprint(count))
}

func (ctl *AccountController) GetAccounts(c *gin.Context) {
	l, _ := c.Get("ledger")

	cursor, err := l.(*ledger.Ledger).GetAccounts(
		c.Request.Context(),
		query.After(c.Query("after")),
		query.Address(c.Query("address")),
		query.Metadata(c.QueryMap("metadata")),
	)
	if err != nil {
		ResponseError(c, err)
		return
	}

	respondWithCursor[core.Account](c, http.StatusOK, cursor)
}

func (ctl *AccountController) GetAccount(c *gin.Context) {
	l, _ := c.Get("ledger")

	acc, err := l.(*ledger.Ledger).GetAccount(
		c.Request.Context(),
		c.Param("address"))
	if err != nil {
		ResponseError(c, err)
		return
	}

	respondWithData[core.Account](c, http.StatusOK, acc)
}

func (ctl *AccountController) PostAccountMetadata(c *gin.Context) {
	l, _ := c.Get("ledger")

	var m core.Metadata
	if err := c.ShouldBindJSON(&m); err != nil {
		ResponseError(c, err)
		return
	}

	addr := c.Param("address")
	if !core.ValidateAddress(addr) {
		ResponseError(c, errors.New("invalid address"))
		return
	}

	if err := l.(*ledger.Ledger).SaveMeta(c.Request.Context(), core.MetaTargetTypeAccount, addr, m); err != nil {
		ResponseError(c, err)
		return
	}

	respondWithNoContent(c)
}
