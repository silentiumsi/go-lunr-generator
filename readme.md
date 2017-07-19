# Go Lunr generator

Parse markdown and html file to extract yaml front-matter and generate Lunr documents.

## Example

### Input

```yaml
---
title: My awesome post
slug: my-awesome-post
url: https://mysite.com/my-awesome-post
tags: just, some, tags
---

# I'm some content written in markdown
```

### Output
```json
[
  {
    "id": "my-awesome-post",
    "title": "My awesome post",
    "url": "https://mysite.com/my-awesome-post",
    "tags": "just, some, tags",
    "content": "\n# I\u0026#39;m some content written in markdown"
  }
]
```

## Note

I am not a Go developer and this was built for a specific purpose and may not fit your needs. Feel free to fork, make more generic, submit merge requests etc.