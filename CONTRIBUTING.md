# Contribution Guidelines

Please ensure your pull request adheres to the following guidelines:

- Make an individual pull request for each suggestion.
- Choose the corresponding patterns section for your suggestion.
- List, after your addition, should be in lexicographical order.

## Commit Messages Guidelines

- The message should be in imperative form and uncapitalized.
- If possible, please include an explanation in the commit message body
- Use the form `<pattern-section>/<pattern-name>: <message>` (e.g. `creational/singleton: refactor singleton constructor`)

## Pattern Template

Each pattern should have a single markdown file containing the important part of the implementation, the usage and the explanations for it. This is to ensure that the reader doesn't have to read bunch of boilerplate to understand what's going on and the code is as simple as possible and not simpler.

Please use the following template for adding new patterns:

```markdown
# <Pattern-Name>
<Pattern description>

## Implementation

## Usage

// Optional
## Rules of Thumb 
```
