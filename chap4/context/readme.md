### Preventing goroutine leaks cancellation aspects:

- a goroutine's parent want to cancel it
- a goroutine may want to cancel its children
- any blocking operations within a goroutine need to be preemtable so that it may be canceled

### Context

- can't mutate the state of underlying structure
- accepting function can't cancel recieved context

### Rules for context.WithValue()
