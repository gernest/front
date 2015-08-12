# front [![Build Status](https://travis-ci.org/gernest/front.svg)](https://travis-ci.org/gernest/front)

extracts frontmatter from text files with ease.

## Installation

	go get github.com/gernest/front

## How To use

create a new `Matter` instance.

	m:=front.NewMatter()

register a frontmatter handler, this is any function which implements `FrontFunc` interface. For the moment front has `JSONHandler`` which handllers json frontmatter. If we register "+++" it means anything between +++ and the next +++  from the beginning of the file is interpreted as frontmatter.

	m.Handle("+++",front.JSONHandler)
	
now you can have your front matter saparated from its body by passing an `io.Reader` to the Parse method.
	
	// The variable front isof type map[string]interface{}
	// it contains all the key value pairs defined in the frontmatter
	//
	// file is an io.Reader (the text imput you want to process)
	front,body,err:=m.Parse(file)

Please see the tests formore details

## Licence

This project is under the MIT Licence. See the [LICENCE](LICENCE) file for the full license text.

