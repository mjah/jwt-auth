module github.com/mjah/jwt-auth

go 1.13

require (
	github.com/asaskevich/govalidator v0.0.0-20200108200545-475eaeb16496
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-gonic/gin v1.5.0
	github.com/google/uuid v1.1.1
	github.com/jinzhu/gorm v1.9.12
	github.com/mitchellh/go-homedir v1.1.0
	github.com/sirupsen/logrus v1.2.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.6.1
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd
)

replace github.com/asaskevich/govalidator => github.com/asaskevich/govalidator v0.0.0-20200108184127-ac8f5a34c3f7
