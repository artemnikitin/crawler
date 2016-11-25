package crawler

const (
	ValidLinks = 3

	ValidHTML = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN"
    "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" lang="en" xml:lang="en">
  <head>
    <meta http-equiv="content-type" content="text/html; charset=utf-8"/>
    <title>title</title>
    <link rel="stylesheet" type="text/css" href="style.css"/>
    <script type="text/javascript" src="script.js"></script>
  </head>
  <body>
  	<div>
   	<a href="http://example.com">Example</a>
   	<a href="http://example.com/2">Example 2</a>
   	<a href="http://example.com/3">Example 3</a>
   	</div>
  </body>
</html>`

	ValidHTMLNoURL = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN"
    "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" lang="en" xml:lang="en">
  <head>
    <meta http-equiv="content-type" content="text/html; charset=utf-8"/>
    <title>title</title>
    <link rel="stylesheet" type="text/css" href="style.css"/>
    <script type="text/javascript" src="script.js"></script>
  </head>
  <body>
  	<div>
	<p>Text</p> <br/>
   	</div>
  </body>
</html>`

	InvalidHTML = `<html>
  <head>
    <title>title</title>
  </head>
  <body>
  	<a/> href="http://example.com">Example</a>
   	<div>
	<p>Text</p> <br/>
   	</div>
</html>`
)
