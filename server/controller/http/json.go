package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/labstack/echo/v4"
)

type JSONSerializer struct{}

func (JSONSerializer) Serialize(ctx echo.Context, v any, indent string) error {
	enc := json.NewEncoder(ctx.Response())
	if indent != "" {
		enc.SetIndent("", indent)
	}
	return enc.Encode(v)
}

func (JSONSerializer) Deserialize(ctx echo.Context, v any) error {
	err := json.NewDecoder(ctx.Request().Body).Decode(v)
	if ute := (*json.UnmarshalTypeError)(nil); errors.As(err, &ute) {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			fmt.Sprintf(
				"unmarshal type error: expected=%v, got=%v, field=%v, offset=%v",
				ute.Type, ute.Value, ute.Field, ute.Offset,
			),
		).SetInternal(err)
	} else if se := (*json.SyntaxError)(nil); errors.As(err, &se) {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			fmt.Sprintf("syntax error: offset=%v, error=%v",
				se.Offset, se.Error(),
			),
		).SetInternal(err)
	}
	return err
}
