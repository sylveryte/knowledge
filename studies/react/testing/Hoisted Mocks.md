# Hoisted Mocks
#testing
#react

To properly mock say `useQuery` and have diff return values
on each test

1. Mock at top

```tsx
const mocks = vi.hoisted(() => {
  return {
    useQuery: vi.fn(),
    useNavigate: vi.fn(),
    useSearch: vi.fn(),
    Link: vi.fn(),
  };
});

vi.mock("@tanstack/react-router", () => ({
  useNavigate: mocks.useNavigate,
  useSearch: mocks.useSearch,
  Link: mocks.Link,
}));

vi.mock("@tanstack/react-query", () => ({
  useQuery: mocks.useQuery,
}));
```

2. Then you manipulate mocks in it

```tsx
it("go back works", async () => {
  mocks.useNavigate.mockReset();
  mocks.useSearch.mockReset();
  mocks.useSearch.mockReturnValue(query);
  const nav = vi.fn();
  mocks.useNavigate.mockReturnValue(nav);
  mocks.useQuery.mockReturnValue({
    data: {
      m: "success",
      d: {
        gantt_chart_change_tickets: [
          {
            id: "ticket-1",
            priority: "Critical",
            planned_implementation_start: "2024-12-01",
            planned_implementation_end: "2024-12-02",
          },
        ],
      },
    },
    isLoading: false,
  });
  render(<Gantt />);
  const gb = screen.getByTestId("sylgoback");
  await userEvent.click(gb);
  expect(nav).toHaveBeenCalled();

  // can also do
  mocks.useSearch.mockImplementationOnce(() => ({ gses: 1 }));
});
```
