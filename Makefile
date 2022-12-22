documentation: ## Generate documentation
	if [ -f ~/gojelastic.env ]; then \
    	mv ~/gojelastic.env ~/gojelastic.env.bak; \
	fi

	mkdir -p ./docs/GoJelastic

	go run ./main.go documentation --token your_token --url your_url

	mv ./docs/GoJelastic/GoJelastic.md ./docs/index.md

	grep -Rl "GoJelastic.md" ./docs/GoJelastic | xargs sed -i 's/GoJelastic.md/..\/index.md/g'
	
	grep -Rl "GoJelastic_" ./docs/index.md | xargs sed -i 's/GoJelastic_/GoJelastic\/GoJelastic_/g'

	if [ -f ~/gojelastic.env.bak ]; then \
		mv ~/gojelastic.env.bak ~/gojelastic.env; \
	fi
.PHONY: documentation