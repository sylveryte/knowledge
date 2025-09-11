# content projection

#angular

## Problem

We are creating a _barrier between the external template that knows the custom properties_ that need to be applied to the input and the plain HTML input itself.

```html
<fa-input icon="envelope" (value)="onNewValue($event)"></fa-input>
```

Suppose you want to make _custom input_ with icons.
You'll have issues with all input _props send, all events to bubble, formControlName things_

## Solution

Using content projection

```html
<fa-input icon="envelope">
  <input type="email" placeholder="Email" />
</fa-input>
```

## How?

We can actually query anything in the content part of the component HTML tag and use it in the internal template as a configuration API, using the **@ContentChild** and **@ContentChildren** decorators.
But we can do more than that, we can also if necessary take anything that is inside the content section, and use it _directly_ inside the component.

```html
<ng-content></ng-content>
```

### Styling projected content

```css
:host ::ng-deep input {
  border: none;
  outline: none;
}
```

_:host_ selector, meaning that the styles will be applied only inside this component
_::ng-deep_ modifier, which means that the style will no longer be scoped only to HTML elements of this particular component, but it will also affect any descendant elements

### Interacting with host or child content

**@ContentChild** gives child ref that is in ng-content
**@HostBinding** can apply style on host item.

```ts
  @ContentChild(InputRefDirective)
  input: InputRefDirective;

  @HostBinding("class.focus")
  get focus() {
    return this.input ? this.input.focus : false;
  }
```

### Multi slot content projection

Here **select** can be any _querySelector_ selector, could be class _.class-name, id, element tag_ etc

```html
<ng-content select="i"></ng-content>
<ng-content select="input"></ng-content>
<ng-content select="input.test-class"></ng-content>
```
