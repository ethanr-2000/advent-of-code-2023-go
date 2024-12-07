SHELL := /bin/bash # Use bash syntax

skeleton: ## make skeleton main(_test).go files, optional: $DAY and $YEAR
	@ if [[ -n $$DAY && -n $$YEAR ]]; then \
		go run scripts/cmd/skeleton/main.go -day $(DAY) -year $(YEAR) ; \
	elif [[ -n $$DAY ]]; then \
		go run scripts/cmd/skeleton/main.go -day $(DAY); \
	else \
		go run scripts/cmd/skeleton/main.go; \
	fi