# Bug reproduction

## Bug
交接名单临时加入 dive-team-b，后面读取 standby-team 时，前一个页面的 dive-team-b 会自己变成 standby-team。不同班组的快照不该共用内存，麻烦修复。

## Trigger

```bash
go test ./handoverroster -run '^TestHandoverContingencySnapshotRemainsDetached$' -count=1
```

## Error

```text
--- FAIL: TestHandoverContingencySnapshotRemainsDetached (0.00s)
    record006_regression_test.go:17: domain compaction changed retained roster: [marine subsea medic ]
FAIL
FAIL	github.com/kekelele996/subsea-cable-repair-control-service/handoverroster	0.350s
FAIL
```
