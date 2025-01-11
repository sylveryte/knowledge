# User Click

```tsx
import userEvent from '@testing-library/user-event';

...
    const gb = screen.getByText('CHG1330').parentNode?.querySelector('img') as HTMLImageElement;
    await userEvent.click(gb);
    expect(mocks.useNavigate).toHaveBeenCalled();
```
