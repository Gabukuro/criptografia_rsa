run-decrypt:
	go run *.go decrypt private.txt source_encoded.txt target_decoded.txt

run-encrypt:
	go run *.go encrypt public.txt source_decoded.txt target_encoded.txt

test:
	cd ./utils && go test