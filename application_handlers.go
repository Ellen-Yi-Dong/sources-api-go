package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/RedHatInsights/sources-api-go/dao"
	m "github.com/RedHatInsights/sources-api-go/model"
	"github.com/RedHatInsights/sources-api-go/util"
	"github.com/labstack/echo/v4"
)

// function that defines how we get the dao - default implementation below.
var getApplicationDao func(c echo.Context) (dao.ApplicationDao, error)

func getApplicationDaoWithTenant(c echo.Context) (dao.ApplicationDao, error) {
	var tenantID int64
	var ok bool

	tenantVal := c.Get("tenantID")
	if tenantID, ok = tenantVal.(int64); !ok {
		return nil, fmt.Errorf("failed to pull tenant from request")
	}

	return &dao.ApplicationDaoImpl{TenantID: &tenantID}, nil
}

func ApplicationList(c echo.Context) error {
	applicationDB, err := getApplicationDao(c)
	if err != nil {
		return err
	}

	filters, err := getFilters(c)
	if err != nil {
		return err
	}

	limit, offset, err := getLimitAndOffset(c)
	if err != nil {
		return err
	}

	var (
		applications []m.Application
		count        *int64
	)

	requestURL, err := util.NewRequestURL(c.Request().RequestURI, c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorDoc("Bad Request", "400"))
	}

	if requestURL.IsSubCollection() {
		applications, count, err = applicationDB.SubCollectionList(requestURL.PrimaryResource(), limit, offset, filters)
	} else {
		applications, count, err = applicationDB.List(limit, offset, filters)
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorDoc("Bad Request", "400"))
	}
	c.Logger().Infof("tenant: %v", *applicationDB.Tenant())

	out := make([]interface{}, len(applications))
	for i, a := range applications {
		out[i] = *a.ToResponse()
	}

	return c.JSON(http.StatusOK, util.CollectionResponse(out, c.Path(), int(count), limit, offset))
}

func ApplicationGet(c echo.Context) error {
	applicationDB, err := getApplicationDao(c)
	if err != nil {
		return err
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	c.Logger().Infof("Getting Application ID %v", id)

	app, err := applicationDB.GetById(&id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, app.ToResponse())
}
