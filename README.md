## Example project

Sample project to create a website and perform tests of xUnit type and e2e in AppEngine.
We will use Beegae, Ginkgo, Gomega, Jasmine, AngularJS and Agouti 

- Beegae    https://github.com/astaxie/beegae
- Ginkgo    https://github.com/onsi/ginkgo
- Gomega    https://github.com/onsi/gomega
- Agouti    https://github.com/sclevine/agouti
- Jasmine   http://jasmine.github.io
- AngularJS https://angularjs.org

### AppEngine

To run the project, in addition to the libraries listed above, you need the [AppEngine SDK](https://cloud.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go)

### Running the project

In the folder `main`

```
$ goapp serve
```

### Running the tests

In the folder `main`

```
$ goapp test
```

### Running Jasmine

In the folder `static/js/jasmine`
Your need download the [Jasmine Standalone](https://github.com/jasmine/jasmine/releases)
and unzip in this folder, maintaining the original file SpecRunner.html 

You should open the file `SpecRunner.html` in your browser
