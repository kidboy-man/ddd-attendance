package restapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	datatype "github.com/kidboy-man/ddd-attendance/internal/data-types"
	domain "github.com/kidboy-man/ddd-attendance/internal/domain/account"
)

type AccountController struct {
	Controller
	accountDomain domain.AccountDomain
}

func NewAccountController() *AccountController {
	return &AccountController{}
}

func (c *AccountController) Prepare() {
	c.accountDomain = domain.NewAccountDomain()
}

// @Summary Register Employee
// @Description Register a new employee
// @Accept json
// @Produce json
// @Param request body schemas.RegisterRequest true "registration data"
// @Success 200 {object} schemas.RegisterResponse
// @Failure 400 {object} schemas.RegisterResponse
// @Router /account/register [post]
func (ac *AccountController) RegisterEmployee(c *gin.Context) {
	registration := datatype.Registration{}
	err := c.BindJSON(&registration)
	if err != nil {
		ac.ReturnNotOK(c, err)
		return
	}
	res, err := ac.accountDomain.Register(c.Request.Context(), registration)
	if err != nil {
		ac.ReturnNotOK(c, err)
		return
	}

	ac.ReturnOK(c, http.StatusCreated, "Successfully Registered", res)
}
