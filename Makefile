USERNAME := sippma
PASSWORD := sippma
DATABASE := sippma
IMAGE := registry.redhat.io/rhel9/postgresql-16:1-14
PORT := 5432

ID := $(shell podman ps -a -q -f "name=postgresql_database")

.PHONY: database-ephemeral
database-ephemeral:
	podman run -d --name postgresql_database \
	  -e POSTGRESQL_USER=$(USERNAME) \
	  -e POSTGRESQL_PASSWORD=$(PASSWORD) \
	  -e POSTGRESQL_DATABASE=$(DATABASE) \
	  -p $(PORT):5432 \
	  $(IMAGE)

.PHONY: clean
clean:
ifneq ($(ID),)
	podman rm -f postgresql_database
endif

.PHONY: start
start:
	podman start postgresql_database
