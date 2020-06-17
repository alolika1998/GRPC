package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kartheekvadde/grpc/proto"
	"google.golang.org/grpc"
)

//AdditionHandler performs Addition Operation
func AdditionHandler(ctx *gin.Context) {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewCalculatorServiceClient(conn)

	a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"})
		return
	}

	b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B"})
		return
	}

	req := &proto.Request{A: int64(a), B: int64(b)}
	if response, err := client.Add(ctx, req); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(response.Result),
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

//SubtractionHandler performs Subtraction Operation
func SubtractionHandler(ctx *gin.Context) {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewCalculatorServiceClient(conn)

	a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"})
		return
	}

	b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B"})
		return
	}

	req := &proto.Request{A: int64(a), B: int64(b)}
	if response, err := client.Subtract(ctx, req); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(response.Result),
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

//MultiplicationHandler performs Multiplication Operation
func MultiplicationHandler(ctx *gin.Context) {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewCalculatorServiceClient(conn)

	a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"})
		return
	}

	b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B"})
		return
	}

	req := &proto.Request{A: int64(a), B: int64(b)}
	if response, err := client.Multiply(ctx, req); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(response.Result),
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

//DivisionHandler performs Division Operation
func DivisionHandler(ctx *gin.Context) {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewCalculatorServiceClient(conn)

	a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"})
		return
	}

	b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B"})
		return
	}

	req := &proto.Request{A: int64(a), B: int64(b)}
	if response, err := client.Divide(ctx, req); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(response.Result),
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

//RemainderHandler performs Modulo Operation
func RemainderHandler(ctx *gin.Context) {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewCalculatorServiceClient(conn)

	a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"})
		return
	}

	b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B"})
		return
	}

	req := &proto.Request{A: int64(a), B: int64(b)}
	if response, err := client.Remainder(ctx, req); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(response.Result),
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
