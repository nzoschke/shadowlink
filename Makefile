deploy:
	flyctl auth docker

	go run github.com/tailscale/mkctr@latest \
		--base="alpine:latest" \
		--gopaths="github.com/nzoschke/shadowlink/cmd/bot:/usr/local/bin/bot" \
		--tags="latest" \
		--repos="registry.fly.io/shadowlink" \
		--target=flyio \
		--push \
		/usr/local/bin/bot

	flyctl deploy -a shadowlink -i registry.fly.io/shadowlink:latest

.PHONY: deploy
