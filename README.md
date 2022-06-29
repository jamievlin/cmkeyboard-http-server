# Coolermaster Masterkeys Pro lighting HTTP server

This is a REST API server that controls cooler master keyboard.

## Instructions on how to build:

1. Make sure go 1.18.2+ for windows is installed.
2. Clone this repository to $GOPATH/src/jamievlin.github.io directory
3. Download the Coolermaster SDK [here](https://templates.coolermaster.com/assets/sdk/coolermaster-sdk.zip).
   If the link does not work, check the website at [here](https://templates.coolermaster.com/).
4. Use the x64 version of the sdk. There should be a `*.lib` file, a header file and a dll file. Copy the *.lib
   file to `third_party/cmsdk/lib` directory, the header to `third_party/cmsdk/include` directory and the DLL
   to where your output directory is (can be `build` directory).
5. Patch the header file to be C-compatible with `sdkdll_patch.patch`. I use msys2 mingw64 shell to achieve this.
6. Make sure you have a gcc-compatible C compiler. I use msys mingw64-shell. Run `go build -o <output_file>`
   with CC environment variable set to your C compiler. I use msys2 mingw64 compiler, though clang for 64-bit Windows
   should also work.
7. (Optional) If you are working with another language, generate an API package with swagger-codegen or
   openapi-generator from `api/api.yaml`. Note that your generator support OAS 3.0+ and `oneOf` semantics.
8. Enjoy!

## License

See LICENSE.txt
