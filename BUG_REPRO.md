# Bug reproduction

## Bug
岸站通知网关提交失败时，调用方还是收到 nil，事务也没回滚，worker 随后把消息确认掉了。帮我把这条失败链修完整。

## Trigger

```bash
go test ./outboxdelivery -run '^TestNotificationDeliveryReturnsCommitFailure$' -count=1
```

## Error

```text
--- FAIL: TestNotificationDeliveryReturnsCommitFailure (0.00s)
    record009_regression_test.go:16: notification failure lost cause: commit: outbox commit failed
FAIL
FAIL	github.com/kekelele996/subsea-cable-repair-control-service/outboxdelivery	0.367s
FAIL
```
