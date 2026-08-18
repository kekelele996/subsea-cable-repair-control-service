# Bug reproduction

## Bug
维修协调器收到 shore、mid-sea、deep-sea 三个跨段，第一条 worker 回传后整批任务偶尔就结束，另外两条结果直接丢失。下面是现场打印的 Go 调用栈，请修复并保留并发检查。

## Trigger

```bash
go test -race ./zonelifecycle -run '^TestCoordinatorKeepsAllThreeRepairSpans$' -count=1
```

## Error

```text
--- FAIL: TestCoordinatorKeepsAllThreeRepairSpans (0.00s)
    record002_regression_test.go:46: domain completion gate opened before all spans finished
FAIL
FAIL	github.com/kekelele996/subsea-cable-repair-control-service/zonelifecycle	0.369s
FAIL
```
