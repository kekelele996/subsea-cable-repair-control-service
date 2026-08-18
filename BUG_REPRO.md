# Bug reproduction

## Bug
请求已经取消，安全评估、计划落库、队列投递和船端 worker 仍然继续，任务最后进入 scheduled。只做诊断，查清取消信号在哪些边界断掉。

## Trigger

```bash
go test ./internal/service -run '^TestCancelledContextStopsEveryMutationBoundary$' -count=1
```

## Error

```text
--- FAIL: TestCancelledContextStopsEveryMutationBoundary (0.00s)
    record005_regression_test.go:17: assessment accepted cancelled context: <nil>
FAIL
FAIL	github.com/kekelele996/subsea-cable-repair-control-service/internal/service	0.363s
FAIL
```
