# Reviews Page Generator

This is a python script that pulls the comments listed under the [Guest Book issue](https://github.com/nisrulz/app-privacy-policy-generator/issues/65) and transforms the returned JSON responses to feed into a Static web page.

Some highlights, the python script:

1. will fetch and download the comments jsons by page only once a week. If attempted again within a week, then it will not make a network call but use the downloaded json files.
1. will download all images contained in the comments into downloaded_images directory.
1. will download all profile pictures for each comment and resize them into 48x48 size.
1. will process markdown contained in the comments into html.
1. will then transform the `template.mustache` into `reviews.html` file.

There are some helper shell scripts available.

1. `setup.sh` - will setup and activate the virtual env for the python tool. It will also download dependencies.
1. `generate.sh` - will generate the `reviews.html` page and copy the required files into `public` directory.
