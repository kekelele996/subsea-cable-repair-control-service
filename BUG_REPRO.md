# Bug reproduction

## Bug
安全规则适配器刚切换时底层对象还没准备好，但维修计划居然被允许出航。不要修改代码，帮我定位这个空实现为什么穿过了所有校验。

## Trigger

```bash
go test ./internal/service -run '^TestTypedNilSafetyRejectedAtEveryBoundary$' -count=1
```

## Error

```text
--- FAIL: TestTypedNilSafetyRejectedAtEveryBoundary (0.00s)
    record004_regression_test.go:24: domain handle reported a typed-nil provider as available
FAIL
FAIL	github.com/kekelele996/subsea-cable-repair-control-service/internal/service	0.374s
FAIL
```
