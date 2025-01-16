remove_ds_store:
	find . -name .DS_Store -print0 | xargs -0 git rm -f --ignore-unmatch

gen-domain:
	gen --sqltype=postgres --connstr='host=localhost port=5432 user=dream dbname=dream password=password sslmode=disable connect_timeout=10000' --database=public --model=domain --out=./gen/code --json --json-fmt=lower_camel --guregu --gorm --overwrite --templateDir=./gen/templates

install-tools:
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/go-delve/delve/cmd/dlv@latest
	go install github.com/air-verse/air@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/smallnest/gen@latest