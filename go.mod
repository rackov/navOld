<<<<<<< HEAD
module github.com/rackov/NavControlSystem

go 1.24.6

replace github.com/rackov/NavControlSystem/pkg/logger => ./pkg/logger

replace github.com/rackov/NavControlSystem/pkg/tnats => ./pkg/tnats

replace github.com/rackov/NavControlSystem/pkg/models => ./pkg/models

replace github.com/rackov/NavControlSystem/pkg/monitoring => ./pkg/monitoring

require (
	github.com/BurntSushi/toml v1.5.0
	github.com/golang/protobuf v1.5.4
	github.com/google/uuid v1.6.0
	github.com/nats-io/nats.go v1.45.0
	github.com/rackov/NavControlSystem/pkg/logger v0.0.0-00010101000000-000000000000
	github.com/rackov/NavControlSystem/pkg/monitoring v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.11.0
	google.golang.org/grpc v1.75.0
	google.golang.org/protobuf v1.36.8
=======
module project

go 1.23.0

toolchain go1.24.4

require (
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible
	github.com/gorilla/mux v1.8.1
	github.com/lib/pq v1.10.9
	github.com/naoina/toml v0.1.1
	github.com/prometheus/client_golang v1.23.0
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.10.0
	golang.org/x/net v0.40.0
	google.golang.org/grpc v1.74.2
	google.golang.org/protobuf v1.36.6
>>>>>>> 28998fa (Проба)
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
<<<<<<< HEAD
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/nats-io/nkeys v0.4.11 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.23.0 // indirect
	github.com/prometheus/client_model v0.6.2 // indirect
	github.com/prometheus/common v0.65.0 // indirect
	github.com/prometheus/procfs v0.16.1 // indirect
	golang.org/x/crypto v0.39.0 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250707201910-8d1bb00bc6a7 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
=======
	github.com/kr/text v0.2.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/naoina/go-stringutil v0.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.6.2 // indirect
	github.com/prometheus/common v0.65.0 // indirect
	github.com/prometheus/procfs v0.16.1 // indirect
	github.com/technoweenie/multipartstreamer v1.0.1 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250528174236-200df99c418a // indirect
>>>>>>> 28998fa (Проба)
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
