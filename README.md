# refire
Simple cli tool for reddit to read filtrated feed.  
2022-05-29 filtration is primitive and relies on at least 1 keyword being mentioned.

## how to use
1. configure `~/.refire.json` in the following format:
```
[
    {
        "name": "some-subreddit",
        "filter_keywords": ["keyword 1", "keyword 2"]
    }
]
```

2. call `refire` in terminal to get a response:
```
>refire
some title
https://link

another post - another title
https://another link
```

## how to build
`go install "github.com/romanthekat/refire@latest"` 