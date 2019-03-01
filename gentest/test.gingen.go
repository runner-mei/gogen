// +build gin

// Please don't edit this file!
package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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
	mux.GET("/echo", func(ctx *gin.Context) {
		var a = ctx.Query("a")

		result := svc.Echo(a)
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/echo", func(ctx *gin.Context) {
		result, err := svc.EchoBody(ctx.Request.Body)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/concat", func(ctx *gin.Context) {
		var a = ctx.Query("a")
		var b = ctx.Query("b")

		result, err := svc.Concat(a, b)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/concat1", func(ctx *gin.Context) {
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
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/concat2/:a/:b", func(ctx *gin.Context) {
		var a = ctx.Param("a")
		var b = ctx.Param("b")

		result, err := svc.Concat2(a, b)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/concat3/:a/:b", func(ctx *gin.Context) {
		var a = ctx.Param("a")
		var b = ctx.Param("b")

		result, err := svc.Concat3(&a, &b)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/sub", func(ctx *gin.Context) {
		var a = ctx.Query("a")
		var start int64
		if s := ctx.Query("start"); s != "" {
			startValue, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", start, s, err)), fmt.Errorf("argument %q is invalid - %q", start, s, err).Error())
				return
			}
			start = startValue
		}

		result, err := svc.Sub(a, start)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.POST("/save/:a", func(ctx *gin.Context) {
		var a = ctx.Param("a")
		var b string
		if err := ctx.Bind(&b); err != nil {
			ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", b, "<no value>", err)), fmt.Errorf("argument %q is invalid - %q", b, "<no value>", err).Error())
			return
		}

		result, err := svc.Save(a, b)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusCreated, result)
		return
	})
	mux.POST("/save2/:a", func(ctx *gin.Context) {
		var a = ctx.Param("a")
		var b string
		if err := ctx.Bind(&b); err != nil {
			ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", b, "<no value>", err)), fmt.Errorf("argument %q is invalid - %q", b, "<no value>", err).Error())
			return
		}

		result, err := svc.Save2(&a, &b)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusCreated, result)
		return
	})
	mux.GET("/add/:a/:b", func(ctx *gin.Context) {
		var a int
		if aValue, err := strconv.ParseInt(ctx.Param("a"), 10, 64); err != nil {
			s := ctx.Param("a")
			ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", a, s, err)), fmt.Errorf("argument %q is invalid - %q", a, s, err).Error())
			return
		} else {
			a = int(aValue)
		}
		var b int
		if bValue, err := strconv.ParseInt(ctx.Param("b"), 10, 64); err != nil {
			s := ctx.Param("b")
			ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", b, s, err)), fmt.Errorf("argument %q is invalid - %q", b, s, err).Error())
			return
		} else {
			b = int(bValue)
		}

		result, err := svc.Add(a, b)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/add2/:a/:b", func(ctx *gin.Context) {
		var a *int
		if aValue, err := strconv.ParseInt(ctx.Param("a"), 10, 64); err != nil {
			s := ctx.Param("a")
			ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", a, s, err)), fmt.Errorf("argument %q is invalid - %q", a, s, err).Error())
			return
		} else {
			a = new(int)
			*a = int(aValue)
		}
		var b *int
		if bValue, err := strconv.ParseInt(ctx.Param("b"), 10, 64); err != nil {
			s := ctx.Param("b")
			ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", b, s, err)), fmt.Errorf("argument %q is invalid - %q", b, s, err).Error())
			return
		} else {
			b = new(int)
			*b = int(bValue)
		}

		result, err := svc.Add2(a, b)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/add3", func(ctx *gin.Context) {
		var a *int
		if s := ctx.Query("a"); s != "" {
			aValue, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", a, s, err)), fmt.Errorf("argument %q is invalid - %q", a, s, err).Error())
				return
			}
			a = new(int)
			*a = int(aValue)
		}
		var b *int
		if s := ctx.Query("b"); s != "" {
			bValue, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", b, s, err)), fmt.Errorf("argument %q is invalid - %q", b, s, err).Error())
				return
			}
			b = new(int)
			*b = int(bValue)
		}

		result, err := svc.Add3(a, b)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/query1", func(ctx *gin.Context) {
		var a = ctx.Query("a")
		var beginAt time.Time
		if s := ctx.Query("beginAt"); s != "" {
			beginAtValue, err := toDatetime(s)
			if err != nil {
				ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", beginAt, s, err)), fmt.Errorf("argument %q is invalid - %q", beginAt, s, err).Error())
				return
			}
			beginAt = beginAtValue
		}
		var endAt time.Time
		if s := ctx.Query("endAt"); s != "" {
			endAtValue, err := toDatetime(s)
			if err != nil {
				ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", endAt, s, err)), fmt.Errorf("argument %q is invalid - %q", endAt, s, err).Error())
				return
			}
			endAt = endAtValue
		}
		var isRaw bool
		if s := ctx.Query("isRaw"); s != "" {
			isRaw = toBool(s)
		}

		result := svc.Query1(a, beginAt, endAt, isRaw)
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/query2/:isRaw", func(ctx *gin.Context) {
		var a = ctx.Query("a")
		var beginAt time.Time
		if s := ctx.Query("beginAt"); s != "" {
			beginAtValue, err := toDatetime(s)
			if err != nil {
				ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", beginAt, s, err)), fmt.Errorf("argument %q is invalid - %q", beginAt, s, err).Error())
				return
			}
			beginAt = beginAtValue
		}
		var endAt time.Time
		if s := ctx.Query("endAt"); s != "" {
			endAtValue, err := toDatetime(s)
			if err != nil {
				ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", endAt, s, err)), fmt.Errorf("argument %q is invalid - %q", endAt, s, err).Error())
				return
			}
			endAt = endAtValue
		}
		var isRaw = toBool(ctx.Param("isRaw"))

		result := svc.Query2(a, beginAt, endAt, isRaw)
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/query3/:isRaw", func(ctx *gin.Context) {
		var a = ctx.Query("a")
		var beginAt time.Time
		if s := ctx.Query("beginAt"); s != "" {
			beginAtValue, err := toDatetime(s)
			if err != nil {
				ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", beginAt, s, err)), fmt.Errorf("argument %q is invalid - %q", beginAt, s, err).Error())
				return
			}
			beginAt = beginAtValue
		}
		var endAt time.Time
		if s := ctx.Query("endAt"); s != "" {
			endAtValue, err := toDatetime(s)
			if err != nil {
				ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", endAt, s, err)), fmt.Errorf("argument %q is invalid - %q", endAt, s, err).Error())
				return
			}
			endAt = endAtValue
		}
		var isRaw = toBool(ctx.Param("isRaw"))

		result := svc.Query3(a, beginAt, endAt, &isRaw)
		ctx.JSON(http.StatusOK, result)
		return
	})
	// Misc: annotation is missing
}

func InitStringSvcImpl(mux gin.IRouter, svc *StringSvcImpl) {
	mux.GET("/echo", func(ctx *gin.Context) {
		var a = ctx.Query("a")

		result := svc.Echo(a)
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/echo_body", func(ctx *gin.Context) {
		result, err := svc.EchoBody(ctx.Request.Body)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/concat", func(ctx *gin.Context) {
		var a = ctx.Query("a")
		var b = ctx.Query("b")

		result, err := svc.Concat(a, b)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/concat1", func(ctx *gin.Context) {
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
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/concat2/:a/:b", func(ctx *gin.Context) {
		var a = ctx.Param("a")
		var b = ctx.Param("b")

		result, err := svc.Concat2(a, b)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/concat3/:a/:b", func(ctx *gin.Context) {
		var a = ctx.Param("a")
		var b = ctx.Param("b")

		result, err := svc.Concat3(&a, &b)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/sub", func(ctx *gin.Context) {
		var a = ctx.Query("a")
		var start int64
		if s := ctx.Query("start"); s != "" {
			startValue, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", start, s, err)), fmt.Errorf("argument %q is invalid - %q", start, s, err).Error())
				return
			}
			start = startValue
		}

		result, err := svc.Sub(a, start)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.POST("/save/:a", func(ctx *gin.Context) {
		var a = ctx.Param("a")
		var b string
		if err := ctx.Bind(&b); err != nil {
			ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", b, "<no value>", err)), fmt.Errorf("argument %q is invalid - %q", b, "<no value>", err).Error())
			return
		}

		result, err := svc.Save(a, b)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusCreated, result)
		return
	})
	mux.POST("/save2/:a", func(ctx *gin.Context) {
		var a = ctx.Param("a")
		var b string
		if err := ctx.Bind(&b); err != nil {
			ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", b, "<no value>", err)), fmt.Errorf("argument %q is invalid - %q", b, "<no value>", err).Error())
			return
		}

		result, err := svc.Save2(&a, &b)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusCreated, result)
		return
	})
	mux.GET("/add/:a/:b", func(ctx *gin.Context) {
		var a int
		if aValue, err := strconv.ParseInt(ctx.Param("a"), 10, 64); err != nil {
			s := ctx.Param("a")
			ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", a, s, err)), fmt.Errorf("argument %q is invalid - %q", a, s, err).Error())
			return
		} else {
			a = int(aValue)
		}
		var b int
		if bValue, err := strconv.ParseInt(ctx.Param("b"), 10, 64); err != nil {
			s := ctx.Param("b")
			ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", b, s, err)), fmt.Errorf("argument %q is invalid - %q", b, s, err).Error())
			return
		} else {
			b = int(bValue)
		}

		result, err := svc.Add(a, b)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/add2/:a/:b", func(ctx *gin.Context) {
		var a *int
		if aValue, err := strconv.ParseInt(ctx.Param("a"), 10, 64); err != nil {
			s := ctx.Param("a")
			ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", a, s, err)), fmt.Errorf("argument %q is invalid - %q", a, s, err).Error())
			return
		} else {
			a = new(int)
			*a = int(aValue)
		}
		var b *int
		if bValue, err := strconv.ParseInt(ctx.Param("b"), 10, 64); err != nil {
			s := ctx.Param("b")
			ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", b, s, err)), fmt.Errorf("argument %q is invalid - %q", b, s, err).Error())
			return
		} else {
			b = new(int)
			*b = int(bValue)
		}

		result, err := svc.Add2(a, b)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/add3", func(ctx *gin.Context) {
		var a *int
		if s := ctx.Query("a"); s != "" {
			aValue, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", a, s, err)), fmt.Errorf("argument %q is invalid - %q", a, s, err).Error())
				return
			}
			a = new(int)
			*a = int(aValue)
		}
		var b *int
		if s := ctx.Query("b"); s != "" {
			bValue, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", b, s, err)), fmt.Errorf("argument %q is invalid - %q", b, s, err).Error())
				return
			}
			b = new(int)
			*b = int(bValue)
		}

		result, err := svc.Add3(a, b)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/query1", func(ctx *gin.Context) {
		var a = ctx.Query("a")
		var beginAt time.Time
		if s := ctx.Query("beginAt"); s != "" {
			beginAtValue, err := toDatetime(s)
			if err != nil {
				ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", beginAt, s, err)), fmt.Errorf("argument %q is invalid - %q", beginAt, s, err).Error())
				return
			}
			beginAt = beginAtValue
		}
		var endAt time.Time
		if s := ctx.Query("endAt"); s != "" {
			endAtValue, err := toDatetime(s)
			if err != nil {
				ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", endAt, s, err)), fmt.Errorf("argument %q is invalid - %q", endAt, s, err).Error())
				return
			}
			endAt = endAtValue
		}
		var isRaw bool
		if s := ctx.Query("isRaw"); s != "" {
			isRaw = toBool(s)
		}

		result := svc.Query1(a, beginAt, endAt, isRaw)
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/query2/:isRaw", func(ctx *gin.Context) {
		var a = ctx.Query("a")
		var beginAt time.Time
		if s := ctx.Query("beginAt"); s != "" {
			beginAtValue, err := toDatetime(s)
			if err != nil {
				ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", beginAt, s, err)), fmt.Errorf("argument %q is invalid - %q", beginAt, s, err).Error())
				return
			}
			beginAt = beginAtValue
		}
		var endAt time.Time
		if s := ctx.Query("endAt"); s != "" {
			endAtValue, err := toDatetime(s)
			if err != nil {
				ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", endAt, s, err)), fmt.Errorf("argument %q is invalid - %q", endAt, s, err).Error())
				return
			}
			endAt = endAtValue
		}
		var isRaw = toBool(ctx.Param("isRaw"))

		result := svc.Query2(a, beginAt, endAt, isRaw)
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/query3/:isRaw", func(ctx *gin.Context) {
		var a = ctx.Query("a")
		var beginAt time.Time
		if s := ctx.Query("beginAt"); s != "" {
			beginAtValue, err := toDatetime(s)
			if err != nil {
				ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", beginAt, s, err)), fmt.Errorf("argument %q is invalid - %q", beginAt, s, err).Error())
				return
			}
			beginAt = beginAtValue
		}
		var endAt time.Time
		if s := ctx.Query("endAt"); s != "" {
			endAtValue, err := toDatetime(s)
			if err != nil {
				ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", endAt, s, err)), fmt.Errorf("argument %q is invalid - %q", endAt, s, err).Error())
				return
			}
			endAt = endAtValue
		}
		var isRaw = toBool(ctx.Param("isRaw"))

		result := svc.Query3(a, beginAt, endAt, &isRaw)
		ctx.JSON(http.StatusOK, result)
		return
	})
	// Misc: annotation is missing
}

func InitStringSvcWithContext(mux gin.IRouter, svc *StringSvcWithContext) {
	mux.GET("/echo", func(ctx *gin.Context) {
		var a = ctx.Query("a")

		result := svc.Echo(ctx.Request.Context(), a)
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/echo", func(ctx *gin.Context) {
		result, err := svc.EchoBody(ctx.Request.Context(), ctx.Request.Body)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/concat", func(ctx *gin.Context) {
		var a = ctx.Query("a")
		var b = ctx.Query("b")

		result, err := svc.Concat(ctx.Request.Context(), a, b)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/concat1", func(ctx *gin.Context) {
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
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/concat2/:a/:b", func(ctx *gin.Context) {
		var a = ctx.Param("a")
		var b = ctx.Param("b")

		result, err := svc.Concat2(ctx.Request.Context(), a, b)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/concat3/:a/:b", func(ctx *gin.Context) {
		var a = ctx.Param("a")
		var b = ctx.Param("b")

		result, err := svc.Concat3(ctx.Request.Context(), &a, &b)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/sub", func(ctx *gin.Context) {
		var a = ctx.Query("a")
		var start int64
		if s := ctx.Query("start"); s != "" {
			startValue, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", start, s, err)), fmt.Errorf("argument %q is invalid - %q", start, s, err).Error())
				return
			}
			start = startValue
		}

		result, err := svc.Sub(ctx.Request.Context(), a, start)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.POST("/save/:a", func(ctx *gin.Context) {
		var a = ctx.Param("a")
		var b string
		if err := ctx.Bind(&b); err != nil {
			ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", b, "<no value>", err)), fmt.Errorf("argument %q is invalid - %q", b, "<no value>", err).Error())
			return
		}

		result, err := svc.Save(ctx.Request.Context(), a, b)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusCreated, result)
		return
	})
	mux.POST("/save2/:a", func(ctx *gin.Context) {
		var a = ctx.Param("a")
		var b string
		if err := ctx.Bind(&b); err != nil {
			ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", b, "<no value>", err)), fmt.Errorf("argument %q is invalid - %q", b, "<no value>", err).Error())
			return
		}

		result, err := svc.Save2(ctx.Request.Context(), &a, &b)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusCreated, result)
		return
	})
	mux.GET("/add/:a/:b", func(ctx *gin.Context) {
		var a int
		if aValue, err := strconv.ParseInt(ctx.Param("a"), 10, 64); err != nil {
			s := ctx.Param("a")
			ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", a, s, err)), fmt.Errorf("argument %q is invalid - %q", a, s, err).Error())
			return
		} else {
			a = int(aValue)
		}
		var b int
		if bValue, err := strconv.ParseInt(ctx.Param("b"), 10, 64); err != nil {
			s := ctx.Param("b")
			ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", b, s, err)), fmt.Errorf("argument %q is invalid - %q", b, s, err).Error())
			return
		} else {
			b = int(bValue)
		}

		result, err := svc.Add(ctx.Request.Context(), a, b)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/add2/:a/:b", func(ctx *gin.Context) {
		var a *int
		if aValue, err := strconv.ParseInt(ctx.Param("a"), 10, 64); err != nil {
			s := ctx.Param("a")
			ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", a, s, err)), fmt.Errorf("argument %q is invalid - %q", a, s, err).Error())
			return
		} else {
			a = new(int)
			*a = int(aValue)
		}
		var b *int
		if bValue, err := strconv.ParseInt(ctx.Param("b"), 10, 64); err != nil {
			s := ctx.Param("b")
			ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", b, s, err)), fmt.Errorf("argument %q is invalid - %q", b, s, err).Error())
			return
		} else {
			b = new(int)
			*b = int(bValue)
		}

		result, err := svc.Add2(ctx.Request.Context(), a, b)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	mux.GET("/add3", func(ctx *gin.Context) {
		var a *int
		if s := ctx.Query("a"); s != "" {
			aValue, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", a, s, err)), fmt.Errorf("argument %q is invalid - %q", a, s, err).Error())
				return
			}
			a = new(int)
			*a = int(aValue)
		}
		var b *int
		if s := ctx.Query("b"); s != "" {
			bValue, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				ctx.String(httpCodeWith(fmt.Errorf("argument %q is invalid - %q", b, s, err)), fmt.Errorf("argument %q is invalid - %q", b, s, err).Error())
				return
			}
			b = new(int)
			*b = int(bValue)
		}

		result, err := svc.Add3(ctx.Request.Context(), a, b)
		if err != nil {
			ctx.String(httpCodeWith(err), err.Error())
			return
		}
		ctx.JSON(http.StatusOK, result)
		return
	})
	// Misc: annotation is missing
}
