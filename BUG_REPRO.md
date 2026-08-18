# Bug reproduction

## Bug
已经 aborted 的维修计划仍能穿过仓储、协调器和事件发布层，最后产生 completed 事件。下面是现场调用栈；先不要改代码，只定位整条非法状态路径。

## Trigger

```bash
go test ./internal/service -run '^TestRejectedCompleteIsFencedAtEveryBoundary$' -count=1
```

## Error

```text
--- FAIL: TestRejectedCompleteIsFencedAtEveryBoundary (0.00s)
    record008_regression_test.go:15: aggregate completed terminal state: {State:completed Revision:8} err=<nil>
FAIL
FAIL	github.com/kekelele996/subsea-cable-repair-control-service/internal/service	0.338s
FAIL
```
