package ports

import (
	"fmt"
	"net/http"

	"github.com/gosom/kit/lib"
	"github.com/gosom/kit/logging"
	"github.com/gosom/kit/web"

	"github.com/gosom/address-parser-go-rest/addressparser"
)

// AddressParserHandler is a handler for parsing addresses
type AddressParserHandler struct {
	log    logging.Logger
	parser addressparser.AddressParser
}

// NewAddressParserHandler creates a new AddressParserHandler
func NewAddressParserHandler(log logging.Logger, parser addressparser.AddressParser) AddressParserHandler {
	return AddressParserHandler{
		log:    log,
		parser: parser,
	}
}

// RegisterRoutes registers the routes for the AddressParserHandler
func (o *AddressParserHandler) RegisterRouters(r web.Router) {
	r.Post("/parse", o.Parse)
}

// Parse is a handler for parsing addresses
//
// @Summary Parse an address into its components
// @Description Parses an address into its components
// @Tags AddressParser
// @Accept  json
// @Produce  json
// @Param input body addressparser.AddressParserInput true "AddressParserInput"
// @Success 200 {object} addressparser.Address
// @Failure 400 {object} web.ErrResponse
// @Failure 422 {object} web.ErrResponse
// @Failure 500 {object} web.ErrResponse
// @Router /parse [post]
func (o *AddressParserHandler) Parse(w http.ResponseWriter, r *http.Request) {
	var payload addressparser.AddressParserInput
	if err := web.DecodeBody(r, &payload, true); err != nil {
		web.JSONError(w, r, fmt.Errorf("%w %s", lib.ErrBadRequest, err))
		return
	}
	result, err := o.parser.Parse(payload)
	if err != nil {
		ae := fmt.Errorf("%w %s", lib.ErrUnprocessable, err.Error())
		web.JSONError(w, r, ae)
		return
	}
	web.JSON(w, r, http.StatusOK, result)
}
