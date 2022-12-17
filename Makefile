default:
	make build
build:
	go build
etcd_snapshot:
	make build
	./crane etcd snapshot