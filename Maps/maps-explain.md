# Maps

## when to use maps or struct?

Use Maps when the property values are closely related, use struct when it's not. Use Maps when you don't know the KEYS ahead of time.

Struct are used more often

Map (VALUE TYPE):

1. all KEYS must be the same type
2. DON'T need to know all the fields at compile
3. KEYS are indexed so we can iterate over it
4. all VALUES must be the same defined type

Struct (REFERENCE TYPE):

1. KEYS can be different types
2. NEED to know all the different fields at compile
3. No indexed KEYS
4. VALUES can be different type
