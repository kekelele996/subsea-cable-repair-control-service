# Bug reproduction

## Bug
风险缓存先返回 armor-damage，另一个消费者追加 vessel-note 后，前一个结果会被反向改写。缓存读者之间必须隔离，请修复这组跨层切片污染。

## Trigger

```bash
go test ./riskbuffers -run '^TestRiskCacheReadersReceiveIndependentSnapshots$' -count=1
```

## Error

```text
--- FAIL: TestRiskCacheReadersReceiveIndependentSnapshots (0.00s)
    record010_regression_test.go:16: domain window reused retained view: [{Code:weather Severity:high Message:}]
FAIL
FAIL	github.com/kekelele996/subsea-cable-repair-control-service/riskbuffers	0.364s
FAIL
```
