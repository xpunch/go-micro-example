module github.com/xpunch/go-micro-example/v4

go 1.19

require (
	github.com/gin-gonic/gin v1.8.1
	github.com/go-micro/plugins/v4/broker/kafka v1.1.0
	github.com/go-micro/plugins/v4/broker/mqtt v1.1.0
	github.com/go-micro/plugins/v4/client/grpc v1.1.0
	github.com/go-micro/plugins/v4/registry/etcd v1.1.0
	github.com/go-micro/plugins/v4/server/grpc v1.1.1
	github.com/go-micro/plugins/v4/server/http v1.1.0
	github.com/go-micro/plugins/v4/wrapper/breaker/hystrix v1.1.0
	github.com/go-micro/plugins/v4/wrapper/trace/opentelemetry v0.0.0
	github.com/go-micro/plugins/v4/wrapper/trace/opentracing v1.1.0
	github.com/google/uuid v1.3.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	go-micro.dev/v4 v4.8.1
	go.opentelemetry.io/otel v1.9.0
	go.opentelemetry.io/otel/exporters/jaeger v1.9.0
	go.opentelemetry.io/otel/sdk v1.9.0
	google.golang.org/protobuf v1.28.0
)

replace github.com/go-micro/plugins/v4/wrapper/trace/opentelemetry => ../../../go/go-micro/plugins/v4/wrapper/trace/opentelemetry

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/Microsoft/go-winio v0.5.0 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20210428141323-04723f9f07d7 // indirect
	github.com/Shopify/sarama v1.30.1 // indirect
	github.com/acomagu/bufpipe v1.0.3 // indirect
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/eapache/go-resiliency v1.2.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/eclipse/paho.mqtt.golang v1.3.5 // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/evanphx/json-patch/v5 v5.5.0 // indirect
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-acme/lego/v4 v4.4.0 // indirect
	github.com/go-git/gcfg v1.5.0 // indirect
	github.com/go-git/go-billy/v5 v5.3.1 // indirect
	github.com/go-git/go-git/v5 v5.4.2 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.0 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.0.4 // indirect
	github.com/goccy/go-json v0.9.7 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.0.0 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.2 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kevinburke/ssh_config v0.0.0-20201106050909-4977a11b4351 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/miekg/dns v1.1.43 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/hashstructure v1.1.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	github.com/urfave/cli/v2 v2.3.0 // indirect
	github.com/xanzy/ssh-agent v0.3.0 // indirect
	go.etcd.io/etcd/api/v3 v3.5.2 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.2 // indirect
	go.etcd.io/etcd/client/v3 v3.5.2 // indirect
	go.opentelemetry.io/otel/trace v1.9.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.17.0 // indirect
	golang.org/x/crypto v0.0.0-20210920023735-84f357641f63 // indirect
	golang.org/x/net v0.0.0-20211020060615-d418f374d309 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20210806184541-e5e7981a1069 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211020151524-b7c3a969101a // indirect
	google.golang.org/grpc v1.42.0 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
