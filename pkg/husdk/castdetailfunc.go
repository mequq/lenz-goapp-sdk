package husdk

import (
	"context"
)

func (c CastDetail) GetCastID(ctx context.Context) string {
	return c.CastID
}

func (c CastDetail) GetPicture(ctx context.Context) string {
	if c.Picture.Icon == nil {
		return ""
	}
	return *c.Picture.Icon
}
