# Bug reproduction

## Bug
同一段海缆已经被另一艘维修船占用时，控制台没有显示可以稍后重试，直接把任务标成永久失败。先不要改文件，帮我查清错误分类为什么走偏。

## Trigger

```bash
go test ./internal/service -run '^TestRetryRecognizesLeaseConflict$' -count=1
```

## Error

```text
--- FAIL: TestRetryRecognizesLeaseConflict (0.00s)
    system_test.go:55: got fail
FAIL
FAIL	github.com/kekelele996/subsea-cable-repair-control-service/internal/service	0.360s
FAIL
```
