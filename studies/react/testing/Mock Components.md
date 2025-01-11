# Mock Components

```tsx
vi.mock("../kg/simple-kg-icon", () => ({
  SimpleKgIcon: vi.fn(() => <div>Icon</div>),
}));

...

expect(screen.queryAllByText('Icon')).toBeTruthy();
```
