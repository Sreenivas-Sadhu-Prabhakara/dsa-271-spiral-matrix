# Spiral Matrix

> NeetCode 150 · Math & Geometry · Blind 75 — four shrinking boundary pointers and two guards; the guards are where every buggy spiral dies. Go, stdlib only, tested.

## Problem

Read an m×n matrix in clockwise spiral order.

## Approach

Keep the unvisited region as a rectangle [top, bottom] × [left, right]; each lap walks its four edges — top row rightward, right column downward, bottom row leftward, left column upward — shrinking the corresponding boundary after each edge. The two mid-lap guards (`top <= bottom` before the bottom edge, `left <= right` before the left edge) handle the degenerate finishes: when a lap starts on a single remaining row or column, the first two edges consume it entirely, and walking back without the guard emits those cells twice. Single-row, single-column and 1×1 tests pin exactly those paths. No visited array, no direction-vector state machine — the rectangle is the state.

**Complexity:** O(m·n) time, O(1) space beyond the output.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
