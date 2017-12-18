package authtokens

type CreateAuthToken struct{}

func (c CreateAuthToken) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
