## Rule Sets

* "A depends on B", or "for A to be selected, B needs to be selected"
```
ruleSet.AddDep(A, B) =>
if isSelected(A) then isSelected(B)
```

* "A and B are exclusive",  or "B and A are exclusive",  or "for A to be
selected, B needs to be unselected; and for B to be selected, A needs to be
unselected"

```
ruleSet.AddConflict(A, B) <=> ruleSet.AddConflict(B, A) =>
if isSelected(A) then !isSelected(B) && if isSelected(B) then !isSelected(A)
```

We say that a set of relations are _coherent_ if the laws above are valid for
that set. For example, this set of relations is coherent:

```
AddDep(A, B) // "A depends on B"
AddDep(B, C) // "B depends on C"
AddConflict(C, D) // "C and D are exclusive"
```

And these sets are _not_ coherent:

```
AddDep(A, B)
AddConflict(A, B)
```

A depends on B, so it's a contradiction that they are exclusive. If A is selected, then B would need to be selected, but that's impossible because, by the exclusion rule, both can't be selected at the same time.

```
AddDep(A, B)
AddDep(B, C)
AddConflict(A, C)
```

The dependency relation is transitive; it's easy to see, from the rules above,
that if A depends on B, and B depends on C, then A also depends on C. So this
is a contradiction for the same reason as the previous case.

## RulseSet

* `NewRuleSet()`: Returns an empty rule set.
* `RuleSet.AddDep(A, B)`: a method for rule sets that adds a new dependency A
  between and B.
* `RuleSet.AddConflict(A, B)`: a method for rule sets that add a new conflict
  between A and B.

* `RuleSet.IsCoherent()`: a method for rule sets that returns true if it is a
  coherent rule set, false otherwise.

## Options

* `New(rs)`: returns a new (empty) collection of selected options (`Opts`) for
  the rule set rs.
* `Opts.Toggle(o)`: a method for a collection of selected options, to set or
  unset option o.
* `Opts.StringSlice()`: returns a slice of string with the current list of
  selected options.


## Test

```
go test
```
