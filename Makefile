# Makefile

commit:
	@read -p "Введите сообщение для коммита: " message; \
	git add .; \
	git commit -m "$$message"; \
	git push -u origin main
	