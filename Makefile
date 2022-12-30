documentation: ## Generate documentation
	if [ -f ~/gojelastic.env ]; then \
    	mv ~/gojelastic.env ~/gojelastic.env.bak; \
	fi

	mkdir -p ./docs/documentation

	go run ./main.go documentation --token your_token --url your_url

	cp ./README.md ./docs/index.md

	grep -Rl "https://yoanbernabeu.github.io/GoJelastic/" ./docs/index.md | xargs sed -i 's/https:\/\/yoanbernabeu.github.io\/GoJelastic\//documentation\/GoJelastic/g'
	
	grep -Rl --exclude=GoJelastic.md "## GoJelastic" ./docs/documentation | xargs sed -i 's/## GoJelastic/# /g'

	grep -Rl --exclude=GoJelastic.md "### " ./docs/documentation | xargs sed -i 's/### /## /g'


	grep -Rl "## GoJelastic" ./docs/documentation/GoJelastic.md | xargs sed -i 's/## GoJelastic/# GoJelastic/g'

	if [ -f ~/gojelastic.env.bak ]; then \
		mv ~/gojelastic.env.bak ~/gojelastic.env; \
	fi
.PHONY: documentation