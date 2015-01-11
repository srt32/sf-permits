Search built on top of San Francisco Building Permits.

Data sourced from
[http://sfdbi.org/building-permits-filed-and-issued](http://sfdbi.org/building-permits-filed-and-issued)

To regenerate the list of links on the page open a browser console and run:

```
links = []

jQuery("a").each(function(tag) { links.push(this.href) })

cleanlinks = ""

for (i = 0; i < links.length - 1; i++) { cleanlinks = cleanlinks + ", " + links[i] }

cleanlinks
```
