package translations

/*
Explanation:

-srclang: The base source lang we are using in our application.
- update: the gotext function we want to execute.
- out: the path you want the message catalog to be output to, relative to the command go generate is called.
- lang: comma-separated list of BCP-47 tags you want to create translations for.
*/

//go:generate gotext -srclang=en-GB update -out=catalog.go -lang=en-GB,de-DE,fr-CH github.com/alextanhongpin/go-text
