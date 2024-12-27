sass --no-source-map --style=compressed web/source/scss:web/compiled/css

deno compile -A --include ./scripts/parser/netlify-dark.json --no-check -o ./scripts/parser/compiled/parser ./scripts/parser/parser.ts

deno compile -A --include ./scripts/parser/netlify-dark.json --target x86_64-unknown-linux-gnu --no-check -o ./scripts/parser/compiled/parser ./scripts/parser/parser.ts