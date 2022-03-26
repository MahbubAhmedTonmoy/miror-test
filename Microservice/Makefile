
install_swagger:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger
	go get github.com/go-playground/validator/v10
	go get -u github.com/gorilla/mux
go111module=off swagger generate spec -o ./swagger.yaml –scan-models
swagger: install_swagger