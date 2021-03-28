# Struct and pointers

# OPERATORS

Turn address into value with \*address / dereferencing
Turn value into address with &value

# Value Types Vs. Reference Types in GO

VALUE:

1. int
2. float
3. string
4. bool
5. stucts ?! WHY THE FUCK

REF:

1. slices
2. maps
3. channels
4. pointers
5. functions

## slices

when slices are created, go creates a slice type with: length, capacity, and pointer to head/data array. and data array. When we refer this slice, yes we are referring to a copy to the slice type, but both slice type have pointer to the head/data array, meaning if we change the array data, it is mutating the original array.
