# Bug reproduction

## Bug
海缆维修清单第一次看到 east-span 和 repeater-17，第二个班组补充备用跨段后，前一个页面里的内容也跟着变了。历史快照不能被后续操作改写，帮我修好。

## Trigger

```bash
go test ./manifestownership -race -run '^TestManifestSnapshotHasIndependentMemory$' -count=1
```

## Error

```text
--- FAIL: TestManifestSnapshotHasIndependentMemory (0.00s)
    record001_regression_test.go:29: domain snapshot changed with caller input: {PlanID:rig-17 Spans:[mutated-input south] Contacts:[control-room] RequiredChecks:map[north:[mutated-check corrosion]]}
FAIL
FAIL	github.com/kekelele996/subsea-cable-repair-control-service/manifestownership	0.872s
FAIL
```
