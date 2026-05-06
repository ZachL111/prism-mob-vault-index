# Review Journal

I treated `prism-mob-vault-index` as a project where the smallest useful behavior should still be inspectable.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its mobile workflows focus without claiming live deployment or external usage.

## Cases

- `baseline`: `form pressure`, score 124, lane `watch`
- `stress`: `sync drift`, score 182, lane `ship`
- `edge`: `local state`, score 204, lane `ship`
- `recovery`: `conflict cost`, score 163, lane `ship`
- `stale`: `form pressure`, score 243, lane `ship`

## Note

The useful failure mode here is a wrong decision on a named case, not a vague style disagreement.
