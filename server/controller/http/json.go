package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

type JSONSerializer struct{}

func (JSONSerializer) Serialize(ctx echo.Context, v any, indent string) error {
	enc := jsoniter.ConfigCompatibleWithStandardLibrary.NewEncoder(ctx.Response())
	if indent != "" {
		enc.SetIndent("", indent)
	}
	return enc.Encode(v)
}

func (JSONSerializer) Deserialize(ctx echo.Context, v any) error {
	err := jsoniter.ConfigCompatibleWithStandardLibrary.NewDecoder(ctx.Request().Body).Decode(v)
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
