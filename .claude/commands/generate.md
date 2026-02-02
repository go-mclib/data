# Regenerate Code

Regenerate Go code from Minecraft server reports.

## Instructions

Run the code generation:

```bash
cd pkg/data && go generate ./... && go fmt ./...
```

Verify the generated code compiles:

```bash
go build ./...
```

If generation fails, check that the JSON report files are present in `vanilla_server_reports/generated/reports/`.
