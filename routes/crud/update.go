package crud

import (
	"fmt"
	"git.kolaente.de/konrad/list/models"
	"github.com/labstack/echo"
	"net/http"
)

// UpdateWeb is the webhandler to update an object
func (c *WebHandler) UpdateWeb(ctx echo.Context) error {
	// Get the object
	if err := ctx.Bind(&c.CObject); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "No or invalid model provided.")
	}

	// Get the ID
	id, err := models.GetIntURLParam("id", ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID.")
	}

	// Check if the user has the right to do that
	currentUser, err := models.GetCurrentUser(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not determine the current user.")
	}

	// Do the update
	err = c.CObject.Update(id, &currentUser)
	if err != nil {
		fmt.Println(err)
		if models.IsErrNeedToBeListAdmin(err) {
			return echo.NewHTTPError(http.StatusForbidden, "You need to be list admin to do that.")
		}

		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, c.CObject)
}
