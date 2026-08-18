# Bug reproduction

## Bug
海缆批次收尾写库失败后接口却返回成功，维修船占用槽位没有释放，失败任务还被确认了。请修复错误返回和清理顺序，不能只遮住其中一个症状。

## Trigger

```bash
go test ./batchcleanup -run '^TestBatchFinalizationPreservesFailureAcrossCleanup$' -count=1
```

## Error

```text
--- FAIL: TestBatchFinalizationPreservesFailureAcrossCleanup (0.00s)
    record007_regression_test.go:15: domain outcome lost primary failure: <nil>
FAIL
FAIL	github.com/kekelele996/subsea-cable-repair-control-service/batchcleanup	0.376s
FAIL
```
