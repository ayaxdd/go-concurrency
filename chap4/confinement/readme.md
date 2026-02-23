Safe work with concurrent code:

- Sync for sharing memory (sync.Mutex)
- Sync via communicating (channels)

Also:

- Immutable data (concurrent safe, smaller critical sections)
- Data protected by confinement

# Confinement

Idea of enshuring information is only ever avaliable from _one_ concurrent process. When this achieved, no sync is needed

## Two kinds of confinement

1. ad hoc - confinement through convention

2. lexical - confinement enforced by compiler
