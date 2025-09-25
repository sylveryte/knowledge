# Python Collections

- any [[Python Data Types|data types]] in list

## List

```python
x = [4, True, 'hi']
```

- Lists are **Mutable**
- length `len(x)`, works with [[Python Strings]]
- **extend** is append from go
- **pop** removes last one and returns
- Accessby `x[0]`
- Assign `x[1]="Cool"`

```python
x.extend([4, 5, 6, 2, False])
a = x.pop()
print(x[0]) # access
x[2]="cool" # assign
```

- List assignment to another variable **Lists are refs**

```python
x = [4, True, 'hi']
y = x
x[1] = "Katu"
print(y) # [4, 'Katu', 'hi']
```

- **To copy** make new list from exising
  - Tip use `x[:]` makes a slice of same
- Can have list and tuples inside list

```python
x = [[1,4],('nice',43)]
```

## Tuple

- Tuples are **Immutable**

```python
t = (2, 5, "hello", True)
print(t[2]) # hello
```

---

[[Python as fast as Possible]]
