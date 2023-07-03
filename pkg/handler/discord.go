package handler

import (
	"context"
	"time"

	"github.com/IDEA/SERVER/pkg/dto"
	"github.com/IDEA/SERVER/pkg/util"
	"github.com/labstack/echo"
)

func (h *Handler) HandleApplication(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var application dto.Application
	var limiter = util.NewRateLimiter(5, 10)
	if err := limiter.Limit(ctx, func() error {
		if err := c.Bind(&application); err != nil {
			return c.JSON(403, err.Error())
		}
		if err := h.ns.NotifyApplication(&application); err != nil {
			return c.JSON(500, err.Error())
		}
		return nil
	}); err != nil {
		return c.JSON(401, err.Error())
	}
	return c.JSON(200, "application success")
}
