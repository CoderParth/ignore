package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

var help string = `
Type 'ignore <project-type>' to create a .gitignore file for your project. 

Example: 'ignore rust'
This will create a .gitignore file for a rust project. 

Incase, you already have a .gitignore file, recommended patterns will be appended, 
and duplicates will not be added to your existing .gitignore file. 

See 'ignore --all' to see the list of supported projects.
`

var supported string = `
List of supported projects:

-> actionscript
-> ada
-> agda
-> al
-> android
-> appceleratortitanium
-> appengine
-> archlinuxpackages
-> autotools
-> ballerina
-> c
-> c++
-> cakephp
-> cfwheels
-> chefcookbook
-> cmake
-> codeigniter
-> commonlisp
-> composer
-> concrete5
-> coq
-> craftcms
-> cuda
-> dart
-> delphi
-> dm
-> drupal
-> eagle
-> elisp
-> elixir
-> elm
-> episerver
-> expressionengine
-> extjs
-> fancy
-> flaxeengine
-> forcedotcom
-> gcov
-> gitbook
-> githubpages
-> go
-> gradle
-> grails
-> gwt
-> haskell
-> igorpro
-> jboss
-> jekyll
-> jenkins_home
-> joomla
-> java
-> labview
-> leiningen
-> lemonstand
-> lithium
-> lua
-> magento
-> maven
-> mercury
-> metaprogrammingsystem
-> nim
-> node
-> objective-c
-> opa
-> packer
-> perl
-> playframework
-> plone
-> prestashop
-> processing
-> python
-> qooxdoo
-> rails
-> racket
-> rescript
-> ruby
-> rust
-> scons
-> scala
-> scheme
-> scrivener
-> swift
-> stella
-> symfony
-> symphonycms
-> tex
-> terraform
-> textpattern
-> twincat3
-> unity
-> visualstudio
-> waf
-> wordpress
-> xojp
-> yeoman
-> zig
-> zendframework
-> zephir
-> zigo
`

func create(projectType string) {
	template := openTemplate(projectType)
	defer template.Close()
	isANewFile := createIfNotExist("./.gitignore")
	ignoreFile := openInAppendMode("./.gitignore")
	defer ignoreFile.Close()
	copy(template, ignoreFile)
	if !isANewFile {
		fs := fileScanner(".gitignore")
		arr := createArr(fs) // array of lines of the file
		unqArr := uniq(arr)
		str := convToStr(unqArr)
		f := openInOverWriteMode(".gitignore")
		defer f.Close()
		writeToFile(f, str)
	}
	fmt.Printf(".gitignore has been created for %s \n", projectType)
}

func openTemplate(file string) *os.File {
	f := fmt.Sprintf("./templates/%s", file)
	template, err := os.Open(f)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("Template not available for %s. Please check 'ignore --all' to see supported projects\n", file)
		}
		log.Fatal(err)
	}
	return template
}

func createIfNotExist(f string) bool {
	isNew := false
	_, err := os.Stat(f)
	if os.IsNotExist(err) {
		isNew = true
		_, err := os.Create(f)
		if err != nil {
			log.Fatal(err)
		}
	}
	return isNew
}

func openInAppendMode(file string) *os.File {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func copy(temp, ignoreFile *os.File) {
	_, err := io.Copy(ignoreFile, temp)
	if err != nil {
		log.Fatal(err)
	}
}

func openInOverWriteMode(file string) *os.File {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func writeToFile(f *os.File, s string) {
	_, err := f.WriteString(s)
	if err != nil {
		log.Fatalf("Error writing to file: %s", err)
	}
}

func createArr(fs *bufio.Scanner) []string {
	arr := []string{}
	for fs.Scan() {
		arr = append(arr, fs.Text())
	}
	return arr
}

func uniq(a []string) []string {
	uniq := []string{}
	n := len(a)
	for i := 0; i < n; i++ {
		if a[i] == " " {
			uniq = append(uniq, a[i])
			continue
		}
		dup := false
		for j := 0; j < len(uniq); j++ {
			if uniq[j] == a[i] {
				dup = true
				break
			}
		}
		if !dup {
			uniq = append(uniq, a[i])
		}
	}
	return uniq
}

func convToStr(arr []string) string {
	joinedStr := ""
	for _, l := range arr {
		joinedStr += l + "\n"
	}
	return joinedStr
}

func fileScanner(file string) *bufio.Scanner {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)
	return fs
}

func main() {
	if len(os.Args) > 1 {
		command := os.Args[1]
		if command == "--all" {
			fmt.Printf("%s \n", supported)
			return
		}
		create(os.Args[1])
		return
	}
	fmt.Printf("%s \n", help)
}
