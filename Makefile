documentation: ## Generate documentation
	if [ -f ~/gojelastic.env ]; then \
    	mv ~/gojelastic.env ~/gojelastic.env.bak; \
	fi

	mkdir -p ./docs/documentation

	go run ./main.go documentation --token your_token --url your_url

	cp ./README.md ./docs/index.md

	grep -Rl "https://yoanbernabeu.github.io/GoJelastic/" ./docs/index.md | xargs sed -i 's/https:\/\/yoanbernabeu.github.io\/GoJelastic\//\/documentation\/GoJelastic/g'
	
	if [ -f ~/gojelastic.env.bak ]; then \
		mv ~/gojelastic.env.bak ~/gojelastic.env; \
	fi
.PHONY: documentation