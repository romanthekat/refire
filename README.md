# refire
Simple cli tool for reddit to read filtrated feed.  
[2022-05-29] filtration is primitive and relies on at least one keyword being mentioned in title.

## usage
1. configure `~/.refire.json` in the following format:
```
[
    {
        "name": "some-subreddit",
        "filter_keywords": ["keyword 1", "keyword 2"]
    }
]
```

2. `refire` in terminal:
```
>refire
- subreddit name
some title
https://link

another post - another title
https://another link
```

## install
`go install "github.com/romanthekat/refire@latest"` 