.PHONY: bin/logseq-pages
bin/logseq-pages:
	go build -o bin/logseq-pages ./cmd/logseq-pages

.PHONY: clean
clean:
	rm -rf bin
