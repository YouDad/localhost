cc: cc_back cc_front

cc_back:
	go build

cc_front:
	yarn run build

b:
	@-bee run 2>&1 |\
		ack --flush --passthru --color --color-match "underline bold red" "(\[ERROR\]|NotImplement|.*panic).*" |\
		ack --flush --passthru --color --color-match "bold cyan" "\[(INFO)\].*" |\
		ack --flush --passthru --color --color-match "bold black" "\[(DEBUG)\].*" |\
		ack --flush --passthru --color --color-match "bold yellow" "\[(WARN)\].*" |\
		ack --flush --passthru --color --color-match "underline bold red on_green" "\[(TRACE)\].*" |\
		ack --flush --passthru --color --color-match "underline bold green" "$(pwd)" |\
		tee -a ".log"

f:
	yarn run build
