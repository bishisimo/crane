default:
	make etcd_snapshot
build:
	go build
etcd_snapshot:
	make build
	./crane etcd snapshot