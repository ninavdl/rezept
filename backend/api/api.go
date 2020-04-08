package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/sour-dough/rezept-backend/db"
)

type request struct {
	code   int
	writer http.ResponseWriter
	req    *http.Request
	params httprouter.Params
	api    *API
	user   *User
	token  string
}

func (api *API) makeHandler(handler func(request) error, requireAuth bool) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		req := request{
			code:   200,
			writer: w,
			req:    r,
			params: p,
			api:    api,
		}
		if !req.auth() {
			return
		}
		if req.user == nil && requireAuth {
			w.WriteHeader(401)
			enc := json.NewEncoder(w)
			enc.Encode(map[string]string{"Error": "Authorization required"})
			return
		}
		err := handler(req)
		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			enc := json.NewEncoder(w)
			enc.Encode(map[string]string{"Error": "Internal Server Error"})
		}
	}
}

func (r *request) writeJson(data interface{}) error {
	r.writer.WriteHeader(r.code)
	enc := json.NewEncoder(r.writer)
	err := enc.Encode(data)
	return err
}

func (r *request) auth() bool {
	authHeader := r.req.Header.Get("Authorization")
	if authHeader == "" {
		return true
	}

	if len(authHeader) < len("Bearer ") || authHeader[:len("Bearer ")] != "Bearer " {
		r.writeError("Use Bearer Authorization", 401)
		return false
	}

	r.token = authHeader[len("Bearer "):]
	dbUser := r.api.db.GetUserBySession(r.token)
	if dbUser == nil {
		r.writeError("Unknown Session ID", 401)
		return false
	}
	u := newUser(dbUser)
	r.user = &u
	return true
}

func (r *request) writeError(message string, code int) error {
	r.writer.WriteHeader(code)
	enc := json.NewEncoder(r.writer)
	err := enc.Encode(map[string]string{"Error": message})
	return err
}

func (r *request) validateStruct(data interface{}) bool {
	err := r.api.validate.Struct(data)
	if err != nil {
		valErrors := err.(validator.ValidationErrors)
		textErrors := valErrors.Translate(r.api.translator)
		messages := make([]string, len(textErrors))
		i := 0
		for _, e := range textErrors {
			messages[i] = e
			i++
		}
		r.writeError(strings.Join(messages, ", "), 400)
		return false
	}
	return true
}

type APIConfig struct {
	Prefix           string
	MaxUploadSize    uint
	UploadPath       string
	MaxThumbnailSize uint
	ImageURL         string
	SignupAllowed    bool
}

type API struct {
	db         *db.DB
	router     *httprouter.Router
	translator ut.Translator
	validate   *validator.Validate
	config     APIConfig
}

func Init(db *db.DB, config APIConfig) *API {
	api := API{
		router: httprouter.New(),
		db:     db,
		config: config,
	}

	var prefix string
	if len(config.Prefix) != 0 && config.Prefix[0] != '/' {
		prefix = "/" + prefix
	}

	if len(config.Prefix) > 1 && config.Prefix[len(config.Prefix)-1] == '/' {
		prefix = config.Prefix[:len(config.Prefix)-2]
	}

	api.router.GET(prefix+"/recipes", api.makeHandler(api.listRecipes, false))
	api.router.GET(prefix+"/recipes/:id", api.makeHandler(api.getRecipe, false))
	api.router.PUT(prefix+"/recipes", api.makeHandler(api.putRecipe, true))
	api.router.POST(prefix+"/recipes/:id", api.makeHandler(api.updateRecipe, true))
	api.router.DELETE(prefix+"/recipes/:id", api.makeHandler(api.deleteRecipe, true))
	api.router.PUT(prefix+"/users", api.makeHandler(api.registerUser, false))
	api.router.PUT(prefix+"/login", api.makeHandler(api.login, false))
	api.router.GET(prefix+"/login", api.makeHandler(api.getLoggedInUser, false))
	api.router.DELETE(prefix+"/login", api.makeHandler(api.logout, true))
	api.router.PUT(prefix+"/image", api.makeHandler((api.uploadImage), true))
	api.router.GET(prefix+"/data", api.makeHandler(api.getPageData, false))

	api.router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	en := en.New()
	translator := ut.New(en, en)
	api.translator, _ = translator.GetTranslator("en")
	api.validate = validator.New()

	return &api
}

func (api *API) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	api.router.ServeHTTP(w, req)
}

func (api *API) CreateImageDir() error {
	return os.MkdirAll(api.config.UploadPath, 0755)
}

func (api *API) GetImagePath(id uint) string {
	return fmt.Sprintf("%s/%d.webp", api.config.UploadPath, id)
}

func (api *API) GetThumbnailPath(id uint) string {
	return fmt.Sprintf("%s/%d_thumbnail.webp", api.config.UploadPath, id)
}

func (api *API) GetImageURL(id uint) string {
	return fmt.Sprintf("%s/%d.webp", api.config.ImageURL, id)
}

func (api *API) GetThumbnailURL(id uint) string {
	return fmt.Sprintf("%s/%d_thumbnail.webp", api.config.ImageURL, id)
}
