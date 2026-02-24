### Pipeline

Pipeline is a series of operations (stages). Stage take data _in_, perform an operation on it and pass data back _out_

### Properties of pipeline stage

- a stage consumes and returns same type
- a stage should be able to be passed around (slices, funcs, structs, ...)

### Stages

#### Batch processing

When stages operate on chunks of data all at once insted of one discrete value at a time

Each stage doubles the size of initial pipeline data

#### Stream processing

When stage recieves and emits one element at a time

Each stage receiving and emitting a discrete value so the size of initial pipeline never expands

##### Downsides

- pipeline is pulling down into the body of _for_ loop and let the _range_ fo the heavy lifting of feeding pipeline (multiple function calls for each eleent in pipeline)
- less scalability
