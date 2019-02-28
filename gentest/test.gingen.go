// Please don't edit this file!
package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func httpCodeWith(err error) int {
	if herr, ok := err.(interface {
		HTTPCode() int
	}); ok {
		return herr.HTTPCode()
	}
	return http.StatusInternalServerError
}

func InitStringSvc(mux gin.IRouter, svc StringSvc) {
	mux.GET("/echo", func(c *gin.Context) {
		var a = ctx.Query("a")

		result := svc.Echo(a)
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/echo", func(c *gin.Context) {

		result, err := svc.EchoBody(ctx.Request.Body)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/concat", func(c *gin.Context) {
		var a = ctx.Query("a")
		var b = ctx.Query("b")

		result, err := svc.Concat(a, b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/concat1", func(c *gin.Context) {
		var a *string
		if s := ctx.Query("a"); s != "" {
			a = &s
		}
		var b *string
		if s := ctx.Query("b"); s != "" {
			b = &s
		}

		result, err := svc.Concat1(a, b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/concat2/:a/:b", func(c *gin.Context) {
		var a = ctx.Param("a")
		var b = ctx.Param("b")

		result, err := svc.Concat2(a, b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/concat3/:a/:b", func(c *gin.Context) {
		var a = ctx.Param("a")
		var b = ctx.Param("b")

		result, err := svc.Concat3(&a, &b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/sub", func(c *gin.Context) {
		var a = ctx.Query("a")
		var start int64
		if s := ctx.Query("start"); s != "" {
			v64, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				ctx.String(fmt.Errorf("argument %q is invalid - %q", start, s, err).Error())
				return
			}
			start = v64
		}

		result, err := svc.Sub(a, start)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.POST("/save/:a", func(c *gin.Context) {
		var a = ctx.Param("a")
		var b string
		if err := ctx.Bind(&b); err != nil {
			ctx.String(fmt.Errorf("argument %q is invalid - %q", b, "<no value>", err).Error())
			return
		}

		result, err := svc.Save(a, b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusCreated, result)
		return
	})
	mux.POST("/save2/:a", func(c *gin.Context) {
		var a = ctx.Param("a")
		var b string
		if err := ctx.Bind(&b); err != nil {
			ctx.String(fmt.Errorf("argument %q is invalid - %q", b, "<no value>", err).Error())
			return
		}

		result, err := svc.Save2(&a, &b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusCreated, result)
		return
	})
	mux.GET("/add/:a/:b", func(c *gin.Context) {
		var a int
		if v64, err := strconv.ParseInt(ctx.Param("a"), 10, 64); err != nil {
			s := ctx.Param("a")
			ctx.String(fmt.Errorf("argument %q is invalid - %q", a, s, err).Error())
			return
		} else {
			a = int(v64)
		}
		var b int
		if v64, err := strconv.ParseInt(ctx.Param("b"), 10, 64); err != nil {
			s := ctx.Param("b")
			ctx.String(fmt.Errorf("argument %q is invalid - %q", b, s, err).Error())
			return
		} else {
			b = int(v64)
		}

		result, err := svc.Add(a, b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/add2/:a/:b", func(c *gin.Context) {
		var a *int
		if v64, err := strconv.ParseInt(ctx.Param("a"), 10, 64); err != nil {
			s := ctx.Param("a")
			ctx.String(fmt.Errorf("argument %q is invalid - %q", a, s, err).Error())
			return
		} else {
			a = new(int)
			*a = int(v64)
		}
		var b *int
		if v64, err := strconv.ParseInt(ctx.Param("b"), 10, 64); err != nil {
			s := ctx.Param("b")
			ctx.String(fmt.Errorf("argument %q is invalid - %q", b, s, err).Error())
			return
		} else {
			b = new(int)
			*b = int(v64)
		}

		result, err := svc.Add2(a, b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/add3", func(c *gin.Context) {
		var a *int
		if s := ctx.Query("a"); s != "" {
			v64, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				ctx.String(fmt.Errorf("argument %q is invalid - %q", a, s, err).Error())
				return
			}
			a = new(int)
			*a = int(v64)
		}
		var b *int
		if s := ctx.Query("b"); s != "" {
			v64, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				ctx.String(fmt.Errorf("argument %q is invalid - %q", b, s, err).Error())
				return
			}
			b = new(int)
			*b = int(v64)
		}

		result, err := svc.Add3(a, b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	// Misc: annotation is missing
}

func InitStringSvcImpl(mux gin.IRouter, svc *StringSvcImpl) {
	mux.GET("/echo", func(c *gin.Context) {
		var a = ctx.Query("a")

		result := svc.Echo(a)
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/echo_body", func(c *gin.Context) {

		result, err := svc.EchoBody(ctx.Request.Body)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/concat", func(c *gin.Context) {
		var a = ctx.Query("a")
		var b = ctx.Query("b")

		result, err := svc.Concat(a, b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/concat1", func(c *gin.Context) {
		var a *string
		if s := ctx.Query("a"); s != "" {
			a = &s
		}
		var b *string
		if s := ctx.Query("b"); s != "" {
			b = &s
		}

		result, err := svc.Concat1(a, b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/concat2/:a/:b", func(c *gin.Context) {
		var a = ctx.Param("a")
		var b = ctx.Param("b")

		result, err := svc.Concat2(a, b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/concat3/:a/:b", func(c *gin.Context) {
		var a = ctx.Param("a")
		var b = ctx.Param("b")

		result, err := svc.Concat3(&a, &b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/sub", func(c *gin.Context) {
		var a = ctx.Query("a")
		var start int64
		if s := ctx.Query("start"); s != "" {
			v64, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				ctx.String(fmt.Errorf("argument %q is invalid - %q", start, s, err).Error())
				return
			}
			start = v64
		}

		result, err := svc.Sub(a, start)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.POST("/save/:a", func(c *gin.Context) {
		var a = ctx.Param("a")
		var b string
		if err := ctx.Bind(&b); err != nil {
			ctx.String(fmt.Errorf("argument %q is invalid - %q", b, "<no value>", err).Error())
			return
		}

		result, err := svc.Save(a, b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusCreated, result)
		return
	})
	mux.POST("/save2/:a", func(c *gin.Context) {
		var a = ctx.Param("a")
		var b string
		if err := ctx.Bind(&b); err != nil {
			ctx.String(fmt.Errorf("argument %q is invalid - %q", b, "<no value>", err).Error())
			return
		}

		result, err := svc.Save2(&a, &b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusCreated, result)
		return
	})
	mux.GET("/add/:a/:b", func(c *gin.Context) {
		var a int
		if v64, err := strconv.ParseInt(ctx.Param("a"), 10, 64); err != nil {
			s := ctx.Param("a")
			ctx.String(fmt.Errorf("argument %q is invalid - %q", a, s, err).Error())
			return
		} else {
			a = int(v64)
		}
		var b int
		if v64, err := strconv.ParseInt(ctx.Param("b"), 10, 64); err != nil {
			s := ctx.Param("b")
			ctx.String(fmt.Errorf("argument %q is invalid - %q", b, s, err).Error())
			return
		} else {
			b = int(v64)
		}

		result, err := svc.Add(a, b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/add2/:a/:b", func(c *gin.Context) {
		var a *int
		if v64, err := strconv.ParseInt(ctx.Param("a"), 10, 64); err != nil {
			s := ctx.Param("a")
			ctx.String(fmt.Errorf("argument %q is invalid - %q", a, s, err).Error())
			return
		} else {
			a = new(int)
			*a = int(v64)
		}
		var b *int
		if v64, err := strconv.ParseInt(ctx.Param("b"), 10, 64); err != nil {
			s := ctx.Param("b")
			ctx.String(fmt.Errorf("argument %q is invalid - %q", b, s, err).Error())
			return
		} else {
			b = new(int)
			*b = int(v64)
		}

		result, err := svc.Add2(a, b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/add3", func(c *gin.Context) {
		var a *int
		if s := ctx.Query("a"); s != "" {
			v64, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				ctx.String(fmt.Errorf("argument %q is invalid - %q", a, s, err).Error())
				return
			}
			a = new(int)
			*a = int(v64)
		}
		var b *int
		if s := ctx.Query("b"); s != "" {
			v64, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				ctx.String(fmt.Errorf("argument %q is invalid - %q", b, s, err).Error())
				return
			}
			b = new(int)
			*b = int(v64)
		}

		result, err := svc.Add3(a, b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	// Misc: annotation is missing
}

func InitStringSvcWithContext(mux gin.IRouter, svc *StringSvcWithContext) {
	mux.GET("/echo", func(c *gin.Context) {
		var a = ctx.Query("a")

		result := svc.Echo(ctx.Request.Context(), a)
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/echo", func(c *gin.Context) {

		result, err := svc.EchoBody(ctx.Request.Context(), ctx.Request.Body)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/concat", func(c *gin.Context) {
		var a = ctx.Query("a")
		var b = ctx.Query("b")

		result, err := svc.Concat(ctx.Request.Context(), a, b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/concat1", func(c *gin.Context) {
		var a *string
		if s := ctx.Query("a"); s != "" {
			a = &s
		}
		var b *string
		if s := ctx.Query("b"); s != "" {
			b = &s
		}

		result, err := svc.Concat1(ctx.Request.Context(), a, b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/concat2/:a/:b", func(c *gin.Context) {
		var a = ctx.Param("a")
		var b = ctx.Param("b")

		result, err := svc.Concat2(ctx.Request.Context(), a, b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/concat3/:a/:b", func(c *gin.Context) {
		var a = ctx.Param("a")
		var b = ctx.Param("b")

		result, err := svc.Concat3(ctx.Request.Context(), &a, &b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/sub", func(c *gin.Context) {
		var a = ctx.Query("a")
		var start int64
		if s := ctx.Query("start"); s != "" {
			v64, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				ctx.String(fmt.Errorf("argument %q is invalid - %q", start, s, err).Error())
				return
			}
			start = v64
		}

		result, err := svc.Sub(ctx.Request.Context(), a, start)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.POST("/save/:a", func(c *gin.Context) {
		var a = ctx.Param("a")
		var b string
		if err := ctx.Bind(&b); err != nil {
			ctx.String(fmt.Errorf("argument %q is invalid - %q", b, "<no value>", err).Error())
			return
		}

		result, err := svc.Save(ctx.Request.Context(), a, b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusCreated, result)
		return
	})
	mux.POST("/save2/:a", func(c *gin.Context) {
		var a = ctx.Param("a")
		var b string
		if err := ctx.Bind(&b); err != nil {
			ctx.String(fmt.Errorf("argument %q is invalid - %q", b, "<no value>", err).Error())
			return
		}

		result, err := svc.Save2(ctx.Request.Context(), &a, &b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusCreated, result)
		return
	})
	mux.GET("/add/:a/:b", func(c *gin.Context) {
		var a int
		if v64, err := strconv.ParseInt(ctx.Param("a"), 10, 64); err != nil {
			s := ctx.Param("a")
			ctx.String(fmt.Errorf("argument %q is invalid - %q", a, s, err).Error())
			return
		} else {
			a = int(v64)
		}
		var b int
		if v64, err := strconv.ParseInt(ctx.Param("b"), 10, 64); err != nil {
			s := ctx.Param("b")
			ctx.String(fmt.Errorf("argument %q is invalid - %q", b, s, err).Error())
			return
		} else {
			b = int(v64)
		}

		result, err := svc.Add(ctx.Request.Context(), a, b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/add2/:a/:b", func(c *gin.Context) {
		var a *int
		if v64, err := strconv.ParseInt(ctx.Param("a"), 10, 64); err != nil {
			s := ctx.Param("a")
			ctx.String(fmt.Errorf("argument %q is invalid - %q", a, s, err).Error())
			return
		} else {
			a = new(int)
			*a = int(v64)
		}
		var b *int
		if v64, err := strconv.ParseInt(ctx.Param("b"), 10, 64); err != nil {
			s := ctx.Param("b")
			ctx.String(fmt.Errorf("argument %q is invalid - %q", b, s, err).Error())
			return
		} else {
			b = new(int)
			*b = int(v64)
		}

		result, err := svc.Add2(ctx.Request.Context(), a, b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/add3", func(c *gin.Context) {
		var a *int
		if s := ctx.Query("a"); s != "" {
			v64, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				ctx.String(fmt.Errorf("argument %q is invalid - %q", a, s, err).Error())
				return
			}
			a = new(int)
			*a = int(v64)
		}
		var b *int
		if s := ctx.Query("b"); s != "" {
			v64, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				ctx.String(fmt.Errorf("argument %q is invalid - %q", b, s, err).Error())
				return
			}
			b = new(int)
			*b = int(v64)
		}

		result, err := svc.Add3(ctx.Request.Context(), a, b)
		if err != nil {
			ctx.String(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	// Misc: annotation is missing
}