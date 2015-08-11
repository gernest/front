# front [![Build Status](https://travis-ci.org/gernest/front.svg)](https://travis-ci.org/gernest/front)

extracts frontmatter from text files with ease.

## How To use

create a new `Matter` instance.

	m:=NewMatter()

register a frontmatter handler, this is any function which implements `FrontFunc` interface. For the moment front has `JSONHandler`` which handllers json frontmatter. If we register "+++" it means anything between +++ and the next +++  from the beginning of the file is interpreted as frontmatter.

	m.Handle("+++",front.JSONHandler)
	
now you can have your front matter saparated from its body by passing an `io.Reader` to the Parse method.
	
	// The variable front isof type map[string]interface{}
	// it contains all the key value pairs defined in the frontmatter
	//
	// file is an io.Reader (the text imput you want to process)
	front,body,err:=m.Parse(file)

Please see the tests formore details

## License

This project is under the MIT License. See the [LICENSE](LICENSE) file for the full license text.

