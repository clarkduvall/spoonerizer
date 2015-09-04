# Spoonerizer
Spoonerize everything.

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

This project allows you to easily deploy a spoonerized website to Heroku. For
the uninformed, read about
[Spoonerisms](https://en.wikipedia.org/wiki/Spoonerism).

## Example
Check out [Nacker Hews](http://www.nackerhews.com) a spoonerized version of
[Hacker News](https://news.ycombinator.com/news). Some gems from there:

![Sirefox in my wocket](img/sirefox.png)
![Mina's Charket Mash](img/mina.png)

Read comments on any post for a crazy time.

## Usage

### Running Locally
Set `BASE_URL` to any site you want:
```bash
PORT=8000 BASE_URL=https://news.ycombinator.com go run web.go
```

### Deploying to Heroku
1.  Click the button above.
1.  Set the `BASE_URL` Heroku config variable to the URL of the site you want
    to spoonerize.
1.  Repeat with another site.

## License
MIT
