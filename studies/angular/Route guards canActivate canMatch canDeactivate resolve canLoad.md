# Route guards canActivate canMatch canDeactivate resolve canLoad

#angular
[Preventing unauthorized access](https://angular.io/guide/router#preventing-unauthorized-access)
NOTE: _Use functional guards instead of class based_

## Route guards

guard returns

- true : navigate
- false : stay put
- UrlTree : redirect
- return asynchronous Observable<boolean>

## canActivate

can navigate to this next url?

## canDeactivate

can navigate from this current url?

## canActivateChild

for children

## canMatch

simply true or false to match or not match
can use in case of {\*\* canLoad} for modules, better use this
in route config

```typescript
canMatch: [() => true];
```

## resolve

load data before component is mounted

## canLoad

for module load
