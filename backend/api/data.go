package api

type PageData struct {
	Users int
	SignupAllowed bool
}

func (api *API) getPageData(r request) error {
	userCount := api.db.GetUserCount()

	return r.writeJson(PageData{
		Users: userCount,
		SignupAllowed: api.config.SignupAllowed,
	})
}